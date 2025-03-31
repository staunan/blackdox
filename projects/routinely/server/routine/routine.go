package routine

import (
	"database/sql"
	"errors"
	_ "fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"routinely/mysqldb"

	"github.com/jmoiron/sqlx"
)

const (
	DEFAULT_ROUTINE_MODE = "Daily"
	DEFAULT_ROUTINE_TIME = "00:00:00"
	JWT_SIGNING_SECRET   = "STAUNAN@ROUTINELY"

	ROUTINE_STATUS_ACTIVE  = "active"
	ROUTINE_STATUS_DELETED = "deleted"
)

type Routine struct {
	ID                   int64
	UserId               int64
	Slug                 string
	Title                string
	Description          string
	Mode                 string
	DailyBasisDays       string
	WeeklyBasisWeekDays  string
	MonthlyBasisDate     int8
	YearlyBasisMonthDate string
	Time                 string
	Status               string
	IsTrash              int8
	CreatedAt            string
}
type RoutineEntry struct {
	ID            int64
	UserID        int64
	RoutineID     int64
	CheckedOnDate string
	CreatedAt     string
}
type RoutineHistory struct {
	ID             int64
	UserID         int64
	RoutineID      int64
	HistoryType    string
	HistoryContent string
	CreatedAt      string
}
type ProgressItem struct {
	RoutineID    int64
	RoutineTitle string
	IsCompleted  bool
	EntryData    RoutineEntry
}

func CreateRoutine(routine Routine) (int64, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return 0, err
	}
	// Preparing SQL statement --
	query := "INSERT INTO `routines` (user_id, slug, routine_title, routine_description, routine_mode, daily_basis_days, weekly_basis_weekday, monthly_basis_date, yearly_basis_month_date, routine_time, routine_status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	insert, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}

	// Finalizing DB Column values --
	// User ID --
	var user_id int64 = routine.UserId
	if user_id == 0 {
		return 0, errors.New("user id should be present")
	}
	// Title --
	var title string = routine.Title
	if title == "" {
		return 0, errors.New("routine title should be present")
	} else if len(title) > 200 {
		return 0, errors.New("routine title is too big")
	}
	// Slug --
	var slug string = createSlug(title)
	// Check if slug already exists --
	var slug_exists bool = checkIfSlugAlreadyExists(user_id, slug)
	if slug_exists {
		return 0, errors.New("duplicate routine slug")
	}
	// Description --
	var description string = routine.Description
	if len(description) > 5000 {
		return 0, errors.New("routine description is too big")
	}
	// Mode --
	var mode string = routine.Mode
	Modes := []string{"Daily", "Weekly", "Monthly", "Yearly"}
	if mode == "" {
		mode = DEFAULT_ROUTINE_MODE
	} else if !slices.Contains(Modes, mode) {
		return 0, errors.New("invalid routine mode provided")
	}
	var daily_basis_days string = ""
	var weekly_basis_weekday string = ""
	var monthly_basis_date int8 = 0
	var yearly_basis_month_date string = ""
	Days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	if mode == "Daily" {
		// Daily Basis --
		if routine.DailyBasisDays == "" {
			daily_basis_days = ""
		} else {
			days := strings.Split(routine.DailyBasisDays, ",")
			if len(days) == 0 {
				return 0, errors.New("select at least one day")
			}
			for _, d := range days {
				if !slices.Contains(Days, d) {
					return 0, errors.New("invalid days value")
				}
			}
			daily_basis_days = routine.DailyBasisDays
		}
	} else if mode == "Weekly" {
		// Weekly Basis --
		if routine.WeeklyBasisWeekDays == "" {
			weekly_basis_weekday = ""
		} else {
			if !slices.Contains(Days, routine.WeeklyBasisWeekDays) {
				return 0, errors.New("invalid weekday value")
			}
			weekly_basis_weekday = routine.DailyBasisDays
		}
	} else if mode == "Monthly" {
		// Weekly Basis --
		if routine.MonthlyBasisDate < 0 || routine.MonthlyBasisDate > 33 {
			return 0, errors.New("invalid date index")
		}
		monthly_basis_date = routine.MonthlyBasisDate
	} else if mode == "Yearly" {
		arr := strings.Split(routine.YearlyBasisMonthDate, "-")
		for _, value := range arr {
			i, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				return 0, errors.New("invalid date format for yearly month date")
			}
			if int8(i) < 0 || int8(i) > 31 {
				return 0, errors.New("month should be between 0 to 12 in yearly month date")
			}
		}
		yearly_basis_month_date = routine.YearlyBasisMonthDate
	}

	// Time --
	var time string = routine.Time
	if time == "" {
		time = DEFAULT_ROUTINE_TIME
	}
	// Is Trash --
	var routine_status string = ROUTINE_STATUS_ACTIVE

	// Execute DB Query --
	dbResponse, err := insert.Exec(user_id, slug, title, description, mode, daily_basis_days, weekly_basis_weekday, monthly_basis_date, yearly_basis_month_date, time, routine_status)
	if err != nil {
		return 0, errors.New("unable to execute query")
	}
	insert.Close()

	// Create History --
	lastInsertedId, err := dbResponse.LastInsertId()
	if err != nil {
		return 0, err
	}
	success, err := createRoutineHistory(user_id, lastInsertedId, "Routine Created", "You have <b>Created</b> this routine")
	if err != nil {
		return 0, err
	}
	if !success {
		return 0, errors.New("unable to create history")
	}
	// Return last inserted ID --
	return lastInsertedId, nil
}

func UpdateRoutine(routine Routine) (Routine, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return routine, err
	}

	// Get routine details --
	routine_details := GetRoutineDetailsById(routine.ID)

	// User ID --
	var user_id int64 = routine.UserId
	if user_id == 0 {
		return routine, errors.New("user id should be present")
	}
	if routine_details.UserId != routine.UserId {
		return routine, errors.New("access denied")
	}
	// Title --
	var title string = routine.Title
	if title == "" {
		return routine, errors.New("routine title should be present")
	} else if len(title) > 200 {
		return routine, errors.New("routine title is too big")
	}
	// Slug --
	var slug string = createSlug(title)
	if slug != routine_details.Slug {
		// Check if slug already exists --
		var slug_exists bool = checkIfSlugAlreadyExists(user_id, slug)
		if slug_exists {
			return routine, errors.New("duplicate routine slug")
		}
	}
	// Description --
	var description string = routine.Description
	if len(description) > 5000 {
		return routine, errors.New("routine description is too big")
	}
	// Update Title and slug --
	if routine_details.Title != routine.Title {
		// Update description --
		// Preparing SQL statement --
		query := "UPDATE `routines` set slug = ?, routine_title = ?  where id = ?"
		updateQuery, err := db.Prepare(query)
		if err != nil {
			return routine, err
		}
		// Execute DB Query --
		_, err = updateQuery.Exec(slug, routine.Title, routine.ID)
		if err != nil {
			return routine, errors.New("unable to execute query")
		}
		updateQuery.Close()
	}
	// Update Description --
	if routine_details.Description != routine.Description {
		// Update description --
		// Preparing SQL statement --
		query := "UPDATE `routines` set routine_description = ? where id = ?"
		updateQuery, err := db.Prepare(query)
		if err != nil {
			return routine, err
		}
		// Execute DB Query --
		_, err = updateQuery.Exec(routine.Description, routine.ID)
		if err != nil {
			return routine, errors.New("unable to execute query")
		}
		updateQuery.Close()
	}
	if routine_details.Mode == "Daily" {
		// Update Daily Basis Days --
		if routine_details.DailyBasisDays != routine.DailyBasisDays {
			Days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
			days := strings.Split(routine.DailyBasisDays, ",")
			for _, d := range days {
				if !slices.Contains(Days, d) {
					return routine, errors.New("invalid days value")
				}
			}
			// Update DailyBasisDays --
			// Preparing SQL statement --
			query := "UPDATE `routines` set daily_basis_days = ? where id = ?"
			updateQuery, err := db.Prepare(query)
			if err != nil {
				return routine, err
			}
			// Execute DB Query --
			_, err = updateQuery.Exec(routine.DailyBasisDays, routine.ID)
			if err != nil {
				return routine, errors.New("unable to execute query")
			}
			updateQuery.Close()
		}
	} else if routine_details.Mode == "Weekly" {
		// Update Weekly Basis Days --
		if routine_details.WeeklyBasisWeekDays != routine.WeeklyBasisWeekDays {
			Days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
			if !slices.Contains(Days, routine.WeeklyBasisWeekDays) {
				return routine, errors.New("invalid days value")
			}
			// Update WeeklyBasisWeekDays --
			// Preparing SQL statement --
			query := "UPDATE `routines` set weekly_basis_weekday = ? where id = ?"
			updateQuery, err := db.Prepare(query)
			if err != nil {
				return routine, err
			}
			// Execute DB Query --
			_, err = updateQuery.Exec(routine.WeeklyBasisWeekDays, routine.ID)
			if err != nil {
				return routine, errors.New("unable to execute query")
			}
			updateQuery.Close()
		}
	} else if routine_details.Mode == "Monthly" {
		// Update Monthly Basis Days --
		if routine_details.MonthlyBasisDate != routine.MonthlyBasisDate {
			// Update MonthlyBasisDate --
			if routine.MonthlyBasisDate < 0 || routine.MonthlyBasisDate > 33 {
				return routine, errors.New("invalid date index")
			}
			// Preparing SQL statement --
			query := "UPDATE `routines` set monthly_basis_date = ? where id = ?"
			updateQuery, err := db.Prepare(query)
			if err != nil {
				return routine, err
			}
			// Execute DB Query --
			_, err = updateQuery.Exec(routine.MonthlyBasisDate, routine.ID)
			if err != nil {
				return routine, errors.New("unable to execute query")
			}
			updateQuery.Close()
		}
	} else if routine_details.Mode == "Yearly" {
		// Update Yearly Basis Days --
		if routine_details.YearlyBasisMonthDate != routine.YearlyBasisMonthDate {
			// Update YearlyBasisMonthDate --
			arr := strings.Split(routine.YearlyBasisMonthDate, "-")
			for _, value := range arr {
				i, err := strconv.ParseInt(value, 10, 32)
				if err != nil {
					return routine, errors.New("invalid date format for yearly month date")
				}
				if int8(i) < 0 || int8(i) > 31 {
					return routine, errors.New("month should be between 1 to 31")
				}
			}
			// Preparing SQL statement --
			query := "UPDATE `routines` set yearly_basis_month_date = ? where id = ?"
			updateQuery, err := db.Prepare(query)
			if err != nil {
				return routine, err
			}
			// Execute DB Query --
			_, err = updateQuery.Exec(routine.YearlyBasisMonthDate, routine.ID)
			if err != nil {
				return routine, errors.New("unable to execute query")
			}
			updateQuery.Close()
		}
	} else {
		return routine, errors.New("mode not recognized")
	}

	// Update Time --
	if routine_details.Time != routine.Time {
		// Update Time --
		// Preparing SQL statement --
		query := "UPDATE `routines` set routine_time = ? where id = ?"
		updateQuery, err := db.Prepare(query)
		if err != nil {
			return routine, err
		}
		// Execute DB Query --
		_, err = updateQuery.Exec(routine.Time, routine.ID)
		if err != nil {
			return routine, errors.New("unable to execute query")
		}
		updateQuery.Close()
	}
	fresh_routine_details := GetRoutineDetailsById(routine.ID)
	return fresh_routine_details, nil
}

func UpdateRoutineStatus(routine Routine) (bool, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return false, err
	}

	// Get routine details --
	routine_details := GetRoutineDetailsById(routine.ID)
	// User ID --
	var user_id int64 = routine.UserId
	if user_id == 0 {
		return false, errors.New("user id should be present")
	}
	if routine_details.UserId != routine.UserId {
		return false, errors.New("access denied")
	}

	// Update Title and slug --
	if routine_details.Status != routine.Status {
		// Update description --
		// Preparing SQL statement --
		query := "UPDATE `routines` set routine_status = ? where id = ?"
		updateQuery, err := db.Prepare(query)
		if err != nil {
			return false, err
		}
		// Execute DB Query --
		_, err = updateQuery.Exec(routine.Status, routine.ID)
		if err != nil {
			return false, errors.New("unable to execute query")
		}
		updateQuery.Close()
	}

	return true, nil
}

func MoveToTrash(routine Routine) (bool, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return false, err
	}

	// Get routine details --
	if routine.ID == 0 {
		return false, errors.New("routine id not present")
	}
	routine_details := GetRoutineDetailsById(routine.ID)
	// User ID --
	var user_id int64 = routine.UserId
	if user_id == 0 {
		return false, errors.New("user id should be present")
	}
	if routine_details.UserId != routine.UserId {
		return false, errors.New("access denied")
	}

	// Update is_trash to 1 --
	if routine_details.IsTrash == 0 {
		// Preparing SQL statement --
		query := "UPDATE `routines` set is_trash = ? where id = ?"
		updateQuery, err := db.Prepare(query)
		if err != nil {
			return false, err
		}
		// Execute DB Query --
		_, err = updateQuery.Exec(1, routine.ID)
		if err != nil {
			return false, errors.New("unable to execute query")
		}
		updateQuery.Close()

		return true, nil
	} else {
		return false, errors.New("already in trash")
	}
}

func RestoreFromTrash(routine Routine) (bool, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return false, err
	}

	// Get routine details --
	if routine.ID == 0 {
		return false, errors.New("routine id not present")
	}
	routine_details := GetRoutineDetailsById(routine.ID)
	// User ID --
	var user_id int64 = routine.UserId
	if user_id == 0 {
		return false, errors.New("user id should be present")
	}
	if routine_details.UserId != routine.UserId {
		return false, errors.New("access denied")
	}

	// Update is_trash to 0 --
	if routine_details.IsTrash == 1 {
		// Preparing SQL statement --
		query := "UPDATE `routines` set is_trash = ? where id = ?"
		updateQuery, err := db.Prepare(query)
		if err != nil {
			return false, err
		}
		// Execute DB Query --
		_, err = updateQuery.Exec(0, routine.ID)
		if err != nil {
			return false, errors.New("unable to execute query")
		}
		updateQuery.Close()

		return true, nil
	} else {
		return false, errors.New("item not in trash")
	}
}

func DeleteRoutineForever(routine Routine) (bool, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return false, err
	}

	// Get routine details --
	if routine.ID == 0 {
		return false, errors.New("routine id not present")
	}
	routine_details := GetRoutineDetailsById(routine.ID)
	// User ID --
	var user_id int64 = routine.UserId
	if user_id == 0 {
		return false, errors.New("user id should be present")
	}
	if routine_details.UserId != routine.UserId {
		return false, errors.New("access denied")
	}

	if routine_details.IsTrash == 1 {
		// Delete all routine history --
		// Preparing SQL statement --
		routine_history_query := "DELETE FROM `routine_history` where routine_id = ?"
		routine_history_updateQuery, err := db.Prepare(routine_history_query)
		if err != nil {
			return false, err
		}
		// Execute DB Query --
		_, err = routine_history_updateQuery.Exec(routine.ID)
		if err != nil {
			return false, errors.New("unable to execute query")
		}

		// Delete all entries --
		// Preparing SQL statement --
		routine_entry_query := "DELETE FROM `routine_entries` where routine_id = ?"
		routine_entry_updateQuery, err := db.Prepare(routine_entry_query)
		if err != nil {
			return false, err
		}
		// Execute DB Query --
		_, err = routine_entry_updateQuery.Exec(routine.ID)
		if err != nil {
			return false, errors.New("unable to execute query")
		}

		// Delete Routine --
		// Preparing SQL statement --
		query := "DELETE FROM `routines` where id = ?"
		updateQuery, err := db.Prepare(query)
		if err != nil {
			return false, err
		}
		// Execute DB Query --
		_, err = updateQuery.Exec(routine.ID)
		if err != nil {
			return false, errors.New("unable to execute query")
		}
		updateQuery.Close()

		return true, nil
	} else {
		return false, errors.New("item not in trash")
	}
}

func VerifyRoutineTitle(title string, user_id int64) (bool, error) {
	// Slug --
	var slug string = createSlug(title)
	// Check if slug already exists --
	var slug_exists bool = checkIfSlugAlreadyExists(user_id, slug)
	if slug_exists {
		return true, nil
	} else {
		return false, nil
	}
}

func GetRoutines(user_id int64) ([]Routine, error) {
	var routines []Routine
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return routines, err
	}

	// Prepare statement for reading data
	var user_id_str string = strconv.Itoa(int(user_id))
	rows, err := db.Query("SELECT id, user_id, slug, routine_title, routine_description, routine_mode, daily_basis_days, weekly_basis_weekday, monthly_basis_date, yearly_basis_month_date, routine_time, routine_status, is_trash, created_at FROM routines where user_id = ?", user_id_str)
	if err != nil {
		panic("Unable to retrieve routine list from Database")
	}
	defer rows.Close()

	// Map to Routine List --
	routines = mapDBDataToRoutineList(rows)
	return routines, nil
}

func GetProgress(user_id int64) ([]RoutineEntry, error) {
	var routine_entries []RoutineEntry
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return routine_entries, errors.New("unable to connect to mysqldb")
	}

	// Get all routines --
	routines, err := GetRoutines(user_id)
	if err != nil {
		return routine_entries, err
	}
	var daily_routine_ids []int64 = []int64{}
	for _, v := range routines {
		if v.Mode == "Daily" {
			daily_routine_ids = append(daily_routine_ids, v.ID)
		}
	}

	// Get Daily Progress -
	var user_id_str string = strconv.Itoa(int(user_id))
	currentTime := time.Now()
	var todayDate string = currentTime.Format("2006-01-02")
	daily_progress_query, args, err := sqlx.In("SELECT id, user_id, routine_id, checked_on_date, created_at FROM routine_entries where routine_id IN(?) and user_id = ? and checked_on_date = ?;", daily_routine_ids, user_id_str, todayDate) // creates the query string and arguments
	if err != nil {
		return routine_entries, errors.New("unable to build query")
	}
	rows, err := db.Query(daily_progress_query, args...)
	if err != nil {
		return routine_entries, errors.New("unable to retrieve daily progress")
	}
	defer rows.Close()

	// Map to Routine List --
	routine_entries = mapDBDataToRoutineEntryList(rows)

	return routine_entries, nil
}

func GetDayProgress(user_id int64, date string) (any, error) {
	var progress_items []ProgressItem
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return progress_items, errors.New("unable to connect to mysqldb")
	}

	// Get Routine Entries --
	var user_id_str string = strconv.Itoa(int(user_id))
	rows, err := db.Query("select id, user_id, routine_id, checked_on_date, created_at FROM routine_entries where user_id = ? and checked_on_date = ?", user_id_str, date)
	if err != nil {
		return progress_items, errors.New("unable to build query")
	}
	defer rows.Close()
	routine_entries := mapDBDataToRoutineEntryList(rows)

	// Get Routines --
	var routine_ids []int64 = []int64{}
	for _, entry := range routine_entries {
		routine_ids = append(routine_ids, entry.RoutineID)
	}
	routines_query, args, err := sqlx.In("SELECT id, routine_title FROM routines where id IN(?) or created_at <= ?;", routine_ids, date)
	if err != nil {
		return progress_items, errors.New("unable to build query")
	}
	routine_rows, err := db.Query(routines_query, args...)
	if err != nil {
		return progress_items, errors.New("unable to retrieve execute query")
	}

	var routine_id int64
	var routine_title string

	for routine_rows.Next() {
		var item ProgressItem
		if err := routine_rows.Scan(&routine_id, &routine_title); err != nil {
			panic("Error while scaning data")
		}
		item.RoutineID = routine_id
		item.RoutineTitle = routine_title
		if slices.Contains(routine_ids, routine_id) {
			item.IsCompleted = true
			item.EntryData = filter(routine_entries, func(s RoutineEntry) bool { return s.RoutineID == routine_id })[0]
		} else {
			item.IsCompleted = false
		}
		progress_items = append(progress_items, item)
	}

	return progress_items, nil
}

func GetRoutineDetailsById(routine_id int64) Routine {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		panic("Unable to connect to db")
	}

	// Get Routine Details from DB --
	var routine_id_str string = strconv.Itoa(int(routine_id))
	row := db.QueryRow("SELECT id, user_id, slug, routine_title, routine_description, routine_mode, daily_basis_days, weekly_basis_weekday, monthly_basis_date, yearly_basis_month_date, routine_time, routine_status, is_trash, created_at FROM routines where id = ?", routine_id_str)
	return mapDBDataToRoutineDetails(row)
}

func GetRoutineDetailsBySlug(user_id int64, routine_slug string) Routine {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		panic("Unable to connect to db")
	}

	var user_id_str string = strconv.Itoa(int(user_id))
	// Get Routine Details from DB --
	row := db.QueryRow("SELECT id, user_id, slug, routine_title, routine_description, routine_mode, daily_basis_days, weekly_basis_weekday, monthly_basis_date, yearly_basis_month_date, routine_time, routine_status, is_trash, created_at FROM routines where user_id = ? and slug = ?", user_id_str, routine_slug)
	return mapDBDataToRoutineDetails(row)
}

func GetRoutineHistories(user_id int64, routine_id int64) ([]RoutineHistory, error) {
	var histories []RoutineHistory
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return histories, errors.New("unable to connect to mysqldb")
	}

	// Prepare statement for reading data
	var user_id_str string = strconv.Itoa(int(user_id))
	var routine_id_str string = strconv.Itoa(int(routine_id))
	rows, err := db.Query("SELECT id, user_id, routine_id, history_type, history_content, created_at FROM routine_history where user_id = ? and routine_id = ? order by created_at desc", user_id_str, routine_id_str)
	if err != nil {
		return histories, errors.New("unable to retrieve routine list from Database")
	}
	defer rows.Close()

	// Map to Routine List --
	histories = mapDBDataToRoutineHistoryList(rows)

	return histories, nil
}

func MarkRoutineAsDone(routine_entry RoutineEntry) (RoutineEntry, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return routine_entry, errors.New("unable to connect to db")
	}

	// Get routine details --
	routine_details := GetRoutineDetailsById(routine_entry.RoutineID)

	// User ID --
	var user_id int64 = routine_entry.UserID
	if user_id == 0 {
		return routine_entry, errors.New("user id should be present")
	}
	if routine_details.UserId != user_id {
		return routine_entry, errors.New("access denied")
	}

	if routine_details.Mode == "Daily" {
		// Check if this routine is allwoed for today to make an entry --
		days := strings.Split(routine_details.DailyBasisDays, ",")
		weekDay := getTodayDayName()
		if !slices.Contains(days, weekDay) {
			return routine_entry, errors.New("this routine is inactive for weekday : " + weekDay)
		}
		// Check if data is already present in database --
		var routine_id_str string = strconv.Itoa(int(routine_entry.RoutineID))
		var checked_on_date_str string = routine_entry.CheckedOnDate
		var user_id_str string = strconv.Itoa(int(routine_entry.UserID))
		var exsisted_row_id int64
		err = db.QueryRow("SELECT id FROM routine_entries where routine_id = ? and checked_on_date = ?", routine_id_str, checked_on_date_str).Scan(&exsisted_row_id)
		switch {
		case err == sql.ErrNoRows:
			// Insert an entry --
			query := "INSERT INTO `routine_entries` (user_id, routine_id, checked_on_date) VALUES (?, ?, ?);"
			insert, err := db.Prepare(query)
			if err != nil {
				return routine_entry, errors.New("unable to prepare query")
			}

			dbResponse, err := insert.Exec(user_id_str, routine_id_str, checked_on_date_str)
			if err != nil {
				return routine_entry, err
			}
			// Return last inserted ID --
			lastInsertId, err := dbResponse.LastInsertId()
			if err != nil {
				return routine_entry, errors.New("unable to retrieve last inserted id")
			}
			routine_entry.ID = lastInsertId
			routine_entry.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

			// Create history --
			success, err := createRoutineHistory(user_id, routine_entry.RoutineID, "Routine Checked", "You have marked this routine as <b>Completed</b>")
			if err != nil {
				return routine_entry, err
			}
			if success {
				return routine_entry, nil
			} else {
				return routine_entry, errors.New("unable to create history")
			}

		case err != nil:
			return routine_entry, err
		default:
			return routine_entry, errors.New("row exists")
		}
	} else {
		return routine_entry, nil
	}
}

func MarkRoutineAsNotDone(routine_entry RoutineEntry) (RoutineEntry, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return routine_entry, errors.New("unable to connect to db")
	}

	// Get routine details --
	routine_details := GetRoutineDetailsById(routine_entry.RoutineID)

	// User ID --
	var user_id int64 = routine_entry.UserID
	if user_id == 0 {
		return routine_entry, errors.New("user id should be present")
	}
	if routine_details.UserId != user_id {
		return routine_entry, errors.New("access denied")
	}

	var routine_id_str string = strconv.Itoa(int(routine_entry.RoutineID))
	var checked_on_date_str string = routine_entry.CheckedOnDate
	var exsisted_row_id int64
	err = db.QueryRow("SELECT id FROM routine_entries where routine_id = ? and checked_on_date = ?", routine_id_str, checked_on_date_str).Scan(&exsisted_row_id)
	switch {
	case err == sql.ErrNoRows:
		return routine_entry, errors.New("entry not found")
	case err != nil:
		return routine_entry, err
	default:
		_, err := db.Exec(`DELETE FROM routine_entries WHERE id = ?`, exsisted_row_id)
		if err != nil {
			return routine_entry, err
		}

		// Create history --
		success, err := createRoutineHistory(user_id, routine_entry.RoutineID, "Routine Unchecked", "You have marked this routine as <b>Not Completed</b>")
		if err != nil {
			return routine_entry, err
		}
		if success {
			return routine_entry, nil
		} else {
			return routine_entry, errors.New("unable to create history")
		}
	}
}

func mapDBDataToRoutineDetails(row *sql.Row) Routine {
	// Declaring variable --
	var routine_id int64
	var user_id int64
	var slug string
	var routine_title string
	var routine_description string
	var routine_mode string
	var daily_basis_days string
	var weekly_basis_weekday string
	var monthly_basis_date int8
	var yearly_basis_month_date string
	var routine_time string
	var routine_status string
	var is_trash int8
	var created_at string

	// Scan fields --
	err := row.Scan(&routine_id, &user_id, &slug, &routine_title, &routine_description, &routine_mode, &daily_basis_days, &weekly_basis_weekday, &monthly_basis_date, &yearly_basis_month_date, &routine_time, &routine_status, &is_trash, &created_at)
	if err != nil {
		panic(err)
	}

	// Create a routine object --
	var routine Routine
	routine.ID = routine_id
	routine.UserId = user_id
	routine.Slug = slug
	routine.Title = routine_title
	routine.Description = routine_description
	routine.Mode = routine_mode
	routine.DailyBasisDays = daily_basis_days
	routine.WeeklyBasisWeekDays = weekly_basis_weekday
	routine.MonthlyBasisDate = monthly_basis_date
	routine.YearlyBasisMonthDate = yearly_basis_month_date
	routine.Time = routine_time
	routine.Status = routine_status
	routine.IsTrash = is_trash
	routine.CreatedAt = created_at
	return routine
}

func mapDBDataToRoutineList(rows *sql.Rows) []Routine {
	var routines []Routine
	var routine_id int64
	var user_id int64
	var slug string
	var title string
	var description string
	var routine_mode string
	var daily_basis_days string
	var weekly_basis_weekday string
	var monthly_basis_date int8
	var yearly_basis_month_date string
	var routine_time string
	var routine_status string
	var is_trash int8
	var created_at string

	for rows.Next() {
		var routine Routine
		if err := rows.Scan(&routine_id, &user_id, &slug, &title, &description, &routine_mode, &daily_basis_days, &weekly_basis_weekday, &monthly_basis_date, &yearly_basis_month_date, &routine_time, &routine_status, &is_trash, &created_at); err != nil {
			panic("Error while scaning routines")
		}
		routine.ID = routine_id
		routine.UserId = user_id
		routine.Slug = slug
		routine.Title = title
		routine.Description = description
		routine.Mode = routine_mode
		routine.DailyBasisDays = daily_basis_days
		routine.WeeklyBasisWeekDays = weekly_basis_weekday
		routine.MonthlyBasisDate = monthly_basis_date
		routine.YearlyBasisMonthDate = yearly_basis_month_date
		routine.Time = routine_time
		routine.Status = routine_status
		routine.IsTrash = is_trash
		routine.CreatedAt = created_at
		routines = append(routines, routine)
	}

	return routines
}

func mapDBDataToRoutineEntryList(rows *sql.Rows) []RoutineEntry {
	var routine_entries []RoutineEntry
	var id int64
	var user_id int64
	var routine_id int64
	var checked_on_date string
	var created_at string

	for rows.Next() {
		var routine_entry RoutineEntry
		if err := rows.Scan(&id, &user_id, &routine_id, &checked_on_date, &created_at); err != nil {
			panic("Error while scaning routines")
		}
		routine_entry.ID = id
		routine_entry.UserID = user_id
		routine_entry.RoutineID = routine_id
		routine_entry.CheckedOnDate = checked_on_date
		routine_entry.CreatedAt = created_at
		routine_entries = append(routine_entries, routine_entry)
	}

	return routine_entries
}

func mapDBDataToRoutineHistoryList(rows *sql.Rows) []RoutineHistory {
	var histories []RoutineHistory

	var id int64
	var user_id int64
	var routine_id int64
	var history_type string
	var history_content string
	var created_at string

	for rows.Next() {
		if err := rows.Scan(&id, &user_id, &routine_id, &history_type, &history_content, &created_at); err != nil {
			panic("Error while scaning routine history")
		}
		var history RoutineHistory
		history.ID = id
		history.UserID = user_id
		history.RoutineID = routine_id
		history.HistoryType = history_type
		history.HistoryContent = history_content
		history.CreatedAt = created_at
		histories = append(histories, history)
	}

	return histories
}

func createSlug(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		panic(err)
	}
	processedString := reg.ReplaceAllString(str, " ")
	processedString = strings.TrimSpace(processedString)
	slug := strings.ReplaceAll(processedString, " ", "-")
	slug = strings.ToLower(slug)
	return slug
}

func checkIfSlugAlreadyExists(user_id int64, slug string) bool {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		panic("Unable to connect to db")
	}

	// Check if data is already present in database --
	var user_id_str string = strconv.Itoa(int(user_id))
	var exsisted_row_id int64
	var row_exist bool = false
	query_err := db.QueryRow("SELECT id FROM routines where user_id = ? and slug = ?", user_id_str, slug).Scan(&exsisted_row_id)
	switch {
	case query_err == sql.ErrNoRows:
		row_exist = false
	case query_err != nil:
		panic(query_err)
	default:
		row_exist = true
	}
	return row_exist
}

func createRoutineHistory(user_id int64, routine_id int64, history_type string, history_content string) (bool, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return false, err
	}

	// Validation --
	if user_id == 0 {
		return false, errors.New("user id should be present")
	}
	if routine_id == 0 {
		return false, errors.New("routine id should be present")
	}
	if history_type == "" {
		return false, errors.New("history type should be present")
	}
	if history_content == "" {
		return false, errors.New("history content should be present")
	}
	// Preparing SQL statement --
	query := "INSERT INTO `routine_history` (user_id, routine_id, history_type, history_content) VALUES (?, ?, ?, ?);"
	insert, err := db.Prepare(query)
	if err != nil {
		return false, err
	}

	dbResponse, err := insert.Exec(user_id, routine_id, history_type, history_content)
	if err != nil {
		return false, errors.New("unable to execute query")
	}
	insert.Close()

	lastInsertedId, err := dbResponse.LastInsertId()
	if err != nil {
		return false, err
	}
	if lastInsertedId > 0 {
		return true, nil
	} else {
		return false, errors.New("unable to retrieve last inserted id")
	}
}

func getTodayDayName() string {
	t := time.Now()
	return t.Weekday().String()[0:3]
}

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
