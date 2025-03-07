package routine

import (
	"database/sql"
	"errors"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"routinely/mysqldb"
)

const (
	DEFAULT_ROUTINE_MODE = "Daily"
	DEFAULT_ROUTINE_TIME = "00:00:00"
	JWT_SIGNING_SECRET   = "STAUNAN@ROUTINELY"
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

func CreateRoutine(routine Routine) (int64, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		panic("Unable to connect to db")
	}
	// Preparing SQL statement --
	query := "INSERT INTO `routines` (user_id, slug, routine_title, routine_description, routine_mode, daily_basis_days, weekly_basis_weekday, monthly_basis_date, yearly_basis_month_date, routine_time, is_trash) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	insert, err := db.Prepare(query)
	if err != nil {
		panic("Unable to prepare query")
	}

	// Finalizing DB Column values --
	// User ID --
	var user_id int64 = routine.UserId
	if user_id == 0 {
		panic("UserID should be present")
	}
	// Title --
	var title string = routine.Title
	if title == "" {
		panic("Routine title should be present")
	} else if len(title) > 200 {
		panic("Routine title is too big")
	}
	// Slug --
	var slug string = createSlug(title)
	// Check if slug already exists --
	var slug_exists bool = checkIfSlugAlreadyExists(user_id, slug)
	if slug_exists {
		panic("Duplicate routine slug")
	}
	// Description --
	var description string = routine.Description
	if len(description) > 5000 {
		panic("Routine description is too big")
	}
	// Mode --
	var mode string = routine.Mode
	Modes := []string{"Daily", "Weekly", "Monthly", "Yearly"}
	if mode == "" {
		mode = DEFAULT_ROUTINE_MODE
	} else if !slices.Contains(Modes, mode) {
		panic("Invalid routine mode provided")
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
			for _, d := range days {
				if !slices.Contains(Days, d) {
					panic("Invalid days value")
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
				panic("Invalid weekday value")
			}
			weekly_basis_weekday = routine.DailyBasisDays
		}
	} else if mode == "Monthly" {
		// Weekly Basis --
		if routine.MonthlyBasisDate < 0 || routine.MonthlyBasisDate > 33 {
			panic("Invalid date index")
		}
		monthly_basis_date = routine.MonthlyBasisDate
	} else if mode == "Yearly" {
		arr := strings.Split(routine.YearlyBasisMonthDate, "-")
		for _, value := range arr {
			i, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				panic("Invalid date format for yearly month date")
			}
			if int8(i) < 0 || int8(i) > 31 {
				panic("Month should be between 0 to 12 in yearly month date")
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
	var is_trash int8 = routine.IsTrash

	// Execute DB Query --
	dbResponse, err := insert.Exec(user_id, slug, title, description, mode, daily_basis_days, weekly_basis_weekday, monthly_basis_date, yearly_basis_month_date, time, is_trash)
	if err != nil {
		panic("Unable to execute query")
	}
	insert.Close()

	// Return last inserted ID --
	return dbResponse.LastInsertId()
}

func GetRoutines(user_id int64) []Routine {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		panic("Unable to connect to db")
	}

	// Prepare statement for reading data
	var user_id_str string = strconv.Itoa(int(user_id))

	rows, err := db.Query("SELECT id, user_id, slug, routine_title, routine_description, routine_mode, daily_basis_days, weekly_basis_weekday, monthly_basis_date, yearly_basis_month_date, routine_time, is_trash, created_at FROM routines where user_id = ?", user_id_str)
	if err != nil {
		panic("Unable to retrieve routine list from Database")
	}
	defer rows.Close()

	// Map to Routine List --
	return mapDBDataToRoutineList(rows)
}

func GetProgress(user_id int64) ([]RoutineEntry, error) {
	var routine_entries []RoutineEntry
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return routine_entries, errors.New("unable to connect to mysqldb")
	}

	// Prepare statement for reading data
	var user_id_str string = strconv.Itoa(int(user_id))
	rows, err := db.Query("SELECT id, user_id, routine_id, checked_on_date, created_at FROM routine_entries where user_id = ?", user_id_str)
	if err != nil {
		return routine_entries, errors.New("unable to retrieve routine list from Database")
	}
	defer rows.Close()

	// Map to Routine List --
	routine_entries = mapDBDataToRoutineEntryList(rows)

	return routine_entries, nil
}

func GetRoutineDetailsById(routine_id int64) Routine {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		panic("Unable to connect to db")
	}

	// Get Routine Details from DB --
	var routine_id_str string = strconv.Itoa(int(routine_id))
	row := db.QueryRow("SELECT id, user_id, slug, routine_title, routine_description, routine_mode, daily_basis_days, weekly_basis_weekday, monthly_basis_date, yearly_basis_month_date, routine_time, is_trash, created_at FROM routines where id = ?", routine_id_str)
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
	row := db.QueryRow("SELECT id, user_id, slug, routine_title, routine_description, routine_mode, daily_basis_days, weekly_basis_weekday, monthly_basis_date, yearly_basis_month_date, routine_time, is_trash, created_at FROM routines where user_id = ? and slug = ?", user_id_str, routine_slug)
	return mapDBDataToRoutineDetails(row)
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
	var is_trash int8
	var created_at string

	// Scan fields --
	err := row.Scan(&routine_id, &user_id, &slug, &routine_title, &routine_description, &routine_mode, &daily_basis_days, &weekly_basis_weekday, &monthly_basis_date, &yearly_basis_month_date, &routine_time, &is_trash, &created_at)
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
	routine.IsTrash = int8(is_trash)
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
	var is_trash int8
	var created_at string

	for rows.Next() {
		var routine Routine
		if err := rows.Scan(&routine_id, &user_id, &slug, &title, &description, &routine_mode, &daily_basis_days, &weekly_basis_weekday, &monthly_basis_date, &yearly_basis_month_date, &routine_time, &is_trash, &created_at); err != nil {
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

func MarkRoutineAsDone(routine_entry RoutineEntry) (RoutineEntry, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return routine_entry, errors.New("unable to connect to db")
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
		return routine_entry, nil
	case err != nil:
		return routine_entry, err
	default:
		return routine_entry, errors.New("row exists")
	}
}

func MarkRoutineAsNotDone(routine_entry RoutineEntry) (bool, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return false, errors.New("unable to connect to db")
	}

	var routine_id_str string = strconv.Itoa(int(routine_entry.RoutineID))
	var checked_on_date_str string = routine_entry.CheckedOnDate
	var exsisted_row_id int64
	err = db.QueryRow("SELECT id FROM routine_entries where routine_id = ? and checked_on_date = ?", routine_id_str, checked_on_date_str).Scan(&exsisted_row_id)
	switch {
	case err == sql.ErrNoRows:
		return false, errors.New("entry not found")
	case err != nil:
		return false, err
	default:
		_, err := db.Exec(`DELETE FROM routine_entries WHERE id = ?`, exsisted_row_id)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}
