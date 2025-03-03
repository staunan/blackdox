package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	MySQLUserName        = "root"
	MySQLPassword        = "0000"
	MySQLHostName        = "127.0.0.1:3306"
	MySQLDatabaseName    = "routinely"
	DEFAULT_ROUTINE_MODE = "Daily"
	DEFAULT_ROUTINE_TIME = "00:00:00"
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
type Response struct {
	HasError bool
	Message  string
	Data     any
}

var db *sql.DB

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Create a DB Connection --
	cfg := mysql.Config{
		User:   MySQLUserName,
		Passwd: MySQLPassword,
		Net:    "tcp",
		Addr:   MySQLHostName,
		DBName: MySQLDatabaseName,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	} // Connected

	e.GET("/", func(c echo.Context) error {
		var user_id int64 = 1
		var routines []Routine = getRoutines(user_id)

		// Return Response --
		var response Response
		response.HasError = false
		response.Message = "Successfully Retrieved List!"
		response.Data = routines
		return c.JSON(http.StatusOK, response)
	})

	e.POST("/all_routines", func(c echo.Context) error {
		var user_id int64 = 1
		var routines []Routine = getRoutines(user_id)

		// Return Response --
		var response Response
		response.HasError = false
		response.Message = "Successfully retrieved list"
		response.Data = routines
		return c.JSON(http.StatusOK, response)
	})

	e.POST("/create_routine", func(c echo.Context) error {
		var routine_details Routine

		// Get Request Data --
		var reqData map[string]any = getRequestData(c)

		// Create Routine Object
		var routine Routine
		routine.UserId = 1
		if reqData["title"] == nil {
			routine.Title = ""
		} else {
			routine.Title = reqData["title"].(string)
		}
		if reqData["description"] == nil {
			routine.Description = ""
		} else {
			routine.Description = reqData["description"].(string)
		}
		if reqData["mode"] == nil {
			routine.Mode = ""
		} else {
			routine.Mode = reqData["mode"].(string)
		}
		if reqData["days"] == nil {
			routine.DailyBasisDays = ""
		} else {
			routine.DailyBasisDays = reqData["days"].(string)
		}
		if reqData["weekday"] == nil {
			routine.WeeklyBasisWeekDays = ""
		} else {
			routine.WeeklyBasisWeekDays = reqData["weekday"].(string)
		}
		if reqData["monthday"] == nil {
			routine.MonthlyBasisDate = 0
		} else {
			routine.MonthlyBasisDate = reqData["monthday"].(int8)
		}
		if reqData["yearlymonthdate"] == nil {
			routine.YearlyBasisMonthDate = ""
		} else {
			routine.YearlyBasisMonthDate = reqData["yearlymonthdate"].(string)
		}
		if reqData["time"] == nil {
			routine.Time = ""
		} else {
			routine.Time = reqData["time"].(string)
		}
		routine.IsTrash = 0
		var last_inserted_id int64 = createRoutine(routine)
		if last_inserted_id == 0 {
			panic("Unable to retrieve last return id")
		} else {
			routine_details = getRoutineDetailsById(last_inserted_id)
		}

		// Return Response --
		var response Response
		response.HasError = false
		response.Message = "Routine has been created sucessfully"
		response.Data = routine_details
		return c.JSON(http.StatusOK, response)
	})

	e.POST("/routine_details", func(c echo.Context) error {
		var user_id int64 = 1
		var routine_details Routine

		// Get Request Data --
		var reqData map[string]any = getRequestData(c)

		if reqData["routine_slug"] == nil {
			panic("Routine Slug is required!")
		} else {
			routine_details.Slug = reqData["routine_slug"].(string)
		}

		routine_details = getRoutineDetailsBySlug(user_id, routine_details.Slug)

		// Return Response --
		var response Response
		response.HasError = false
		response.Message = "Successfully retrieved routine data!"
		response.Data = routine_details
		return c.JSON(http.StatusOK, response)
	})

	e.POST("/mark_routine_as_done", func(c echo.Context) error {
		var routine_entry RoutineEntry

		// Get Request Data --
		var reqData map[string]any = getRequestData(c)
		routine_entry.UserID = 1
		if reqData["routine_id"] == nil {
			panic("Routine ID is required!")
		} else {
			routine_entry.RoutineID = int64(reqData["routine_id"].(float64))
		}
		if reqData["checked_on_date"] == nil {
			panic("Checked On Date is required!")
		} else {
			routine_entry.CheckedOnDate = reqData["checked_on_date"].(string)
		}

		// Check if data is already present in database --
		var routine_id_str string = strconv.Itoa(int(routine_entry.RoutineID))
		var checked_on_date_str string = routine_entry.CheckedOnDate
		var user_id_str string = strconv.Itoa(int(routine_entry.UserID))
		var exsisted_row_id int64
		var row_exist bool = false
		err = db.QueryRow("SELECT id FROM routine_entries where routine_id = ? and checked_on_date = ?", routine_id_str, checked_on_date_str).Scan(&exsisted_row_id)
		switch {
		case err == sql.ErrNoRows:
			row_exist = false
		case err != nil:
			panic(err)
		default:
			row_exist = true
		}
		if row_exist && exsisted_row_id > 0 {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = "Already marked as done"
			response.Data = nil
			return c.JSON(http.StatusOK, response)
		} else {
			// Insert an entry --
			query := "INSERT INTO `routine_entries` (user_id, routine_id, checked_on_date) VALUES (?, ?, ?);"
			insert, err := db.Prepare(query)
			if err != nil {
				panic("Unable to prepare query")
			}

			dbResponse, err := insert.Exec(user_id_str, routine_id_str, checked_on_date_str)
			if err != nil {
				panic(err)
			}
			// Return last inserted ID --
			lastInsertId, err := dbResponse.LastInsertId()
			if err != nil {
				panic("Unable to retrieve last inserted id")
			}
			routine_entry.ID = lastInsertId
			routine_entry.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

			// Return Response --
			var response Response
			response.HasError = false
			response.Message = "Today's entry has been created for this routine"
			response.Data = routine_entry
			return c.JSON(http.StatusOK, response)
		}
	})

	e.POST("/mark_routine_as_not_done", func(c echo.Context) error {
		var routine_entry RoutineEntry

		// Get Request Data --
		var reqData map[string]any = getRequestData(c)
		routine_entry.UserID = 1
		if reqData["routine_id"] == nil {
			panic("Routine ID is required!")
		} else {
			routine_entry.RoutineID = int64(reqData["routine_id"].(float64))
		}
		if reqData["checked_on_date"] == nil {
			panic("Checked On Date is required!")
		} else {
			routine_entry.CheckedOnDate = reqData["checked_on_date"].(string)
		}

		// Check if data is already present in database --
		var routine_id_str string = strconv.Itoa(int(routine_entry.RoutineID))
		var checked_on_date_str string = routine_entry.CheckedOnDate
		var exsisted_row_id int64
		var row_exist bool = false
		err = db.QueryRow("SELECT id FROM routine_entries where routine_id = ? and checked_on_date = ?", routine_id_str, checked_on_date_str).Scan(&exsisted_row_id)
		switch {
		case err == sql.ErrNoRows:
			row_exist = false
		case err != nil:
			panic(err)
		default:
			row_exist = true
		}
		if row_exist && exsisted_row_id > 0 {
			// Remove the entries --
			_, err := db.Exec(`DELETE FROM routine_entries WHERE id = ?`, exsisted_row_id)
			if err != nil {
				log.Fatal(err)
			}
			// Return Response --
			var response Response
			response.HasError = false
			response.Message = "Routine has been marked as not done"
			response.Data = nil
			return c.JSON(http.StatusOK, response)
		} else {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = "Entry not found"
			response.Data = routine_entry
			return c.JSON(http.StatusOK, response)
		}
	})

	e.GET("/progress", func(c echo.Context) error {
		var user_id int64 = 1
		var routine_entries []RoutineEntry = getProgress(user_id)
		// Return Response --
		var response Response
		response.HasError = false
		response.Message = "Successfully retrieved today's progress!"
		response.Data = routine_entries
		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func createRoutine(routine Routine) int64 {
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
	var Modes = []string{"Daily", "Weekly", "Monthly", "Yearly"}
	if mode == "" {
		mode = DEFAULT_ROUTINE_MODE
	} else if !slices.Contains(Modes, mode) {
		panic("Invalid routine mode provided")
	}
	var daily_basis_days string = ""
	var weekly_basis_weekday string = ""
	var monthly_basis_date int8 = 0
	var yearly_basis_month_date string = ""
	var Days = []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
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
	lastInsertId, err := dbResponse.LastInsertId()
	if err != nil {
		panic("Unable to retrieve last inserted id")
	}
	return lastInsertId
}

func getRoutines(user_id int64) []Routine {
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

func getProgress(user_id int64) []RoutineEntry {
	// Prepare statement for reading data
	var user_id_str string = strconv.Itoa(int(user_id))
	rows, err := db.Query("SELECT id, user_id, routine_id, checked_on_date, created_at FROM routine_entries where user_id = ?", user_id_str)
	if err != nil {
		panic("Unable to retrieve routine list from Database")
	}
	defer rows.Close()

	// Map to Routine List --
	return mapDBDataToRoutineEntryList(rows)
}

func getRoutineDetailsById(routine_id int64) Routine {
	// Get Routine Details from DB --
	var routine_id_str string = strconv.Itoa(int(routine_id))
	row := db.QueryRow("SELECT id, user_id, slug, routine_title, routine_description, routine_mode, daily_basis_days, weekly_basis_weekday, monthly_basis_date, yearly_basis_month_date, routine_time, is_trash, created_at FROM routines where id = ?", routine_id_str)
	return mapDBDataToRoutineDetails(row)
}

func getRoutineDetailsBySlug(user_id int64, routine_slug string) Routine {
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

func getRequestData(c echo.Context) map[string]interface{} {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return json_map
	} else {
		return json_map
	}
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
	// Check if data is already present in database --
	var user_id_str string = strconv.Itoa(int(user_id))
	var exsisted_row_id int64
	var row_exist bool = false
	err := db.QueryRow("SELECT id FROM routines where user_id = ? and slug = ?", user_id_str, slug).Scan(&exsisted_row_id)
	switch {
	case err == sql.ErrNoRows:
		row_exist = false
	case err != nil:
		panic(err)
	default:
		row_exist = true
	}
	return row_exist
}
