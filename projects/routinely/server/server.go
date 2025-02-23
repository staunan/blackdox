package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

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
	ID          int64
	UserId      int64
	Title       string
	Description string
	Mode        string
	Time        string
	IsTrash     string
	CreatedAt   string
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

	e.POST("/create_routine", func(c echo.Context) error {
		var routine Routine
		var routine_details Routine
		routine.UserId = 1
		routine.Title = "Demo Title"
		routine.Description = "Demo Description"

		var last_inserted_id int64 = createRoutine(routine)
		if err != nil {
			panic("Unable to prepare query")
		}
		if last_inserted_id == 0 {
			panic("Unable to retrieve last return id")
		} else {
			routine_details = getRoutineDetails(last_inserted_id)
		}

		// Return Response --
		var response Response
		response.HasError = false
		response.Message = "Successfully Retrieved Routine Details!"
		response.Data = routine_details
		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func createRoutine(routine Routine) int64 {
	// Preparing SQL statement --
	query := "INSERT INTO `routines` (user_id, routine_title, routine_description, routine_mode, routine_time, is_trash) VALUES (?, ?, ?, ?, ?, ?);"
	insert, err := db.Prepare(query)
	if err != nil {
		panic("Unable to prepare query")
	}

	// Finalizing DB Column values --
	var user_id int64 = routine.UserId
	if user_id == 0 {
		panic("UserID should be present")
	}
	var title string = routine.Title
	if title == "" {
		panic("Please provide title")
	}
	var description string = routine.Description
	var mode string = routine.Mode
	if mode == "" {
		mode = DEFAULT_ROUTINE_MODE
	}
	var time string = routine.Time
	if time == "" {
		time = DEFAULT_ROUTINE_TIME
	}
	var is_trash int = 0
	if routine.IsTrash == "Yes" {
		is_trash = 1
	}

	// Execute DB Query --
	dbResponse, err := insert.Exec(user_id, title, description, mode, time, is_trash)
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
	rows, err := db.Query("SELECT id, user_id, routine_title, routine_description, routine_mode, routine_time, is_trash, created_at FROM routines where user_id = ?", user_id_str)
	if err != nil {
		panic("Unable to retrieve routine list from Database")
	}
	defer rows.Close()

	// Map to Routine List --
	return mapDBDataToRoutineList(rows)
}

func getRoutineDetails(routine_id int64) Routine {
	// Get Routine Details from DB --
	var routine_id_str string = strconv.Itoa(int(routine_id))
	row := db.QueryRow("SELECT id, user_id, routine_title, routine_description, routine_mode, routine_time, is_trash, created_at FROM routines where id = ?", routine_id_str)
	var routine Routine = mapDBDataToRoutineDetails(row)
	return routine
}

func mapDBDataToRoutineDetails(row *sql.Row) Routine {
	// Declaring variable --
	var routine_id int64
	var user_id int64
	var routine_title string
	var routine_description string
	var routine_mode string
	var routine_time string
	var is_trash int64
	var created_at string

	// Scan fields --
	err := row.Scan(&routine_id, &user_id, &routine_title, &routine_description, &routine_mode, &routine_time, &is_trash, &created_at)
	if err != nil {
		panic(err)
	}

	// Create a routine object --
	var routine Routine
	routine.ID = routine_id
	routine.UserId = user_id
	routine.Title = routine_title
	routine.Description = routine_description
	routine.Mode = routine_mode
	routine.Time = routine_time
	if is_trash == 1 {
		routine.IsTrash = "Yes"
	} else {
		routine.IsTrash = "No"
	}
	routine.CreatedAt = created_at
	return routine
}

func mapDBDataToRoutineList(rows *sql.Rows) []Routine {
	var routines []Routine
	var routine_id int64
	var user_id int64
	var title string
	var description string
	var routine_mode string
	var routine_time string
	var is_trash int8
	var created_at string

	for rows.Next() {
		var routine Routine
		if err := rows.Scan(&routine_id, &user_id, &title, &description, &routine_mode, &routine_time, &is_trash, &created_at); err != nil {
			panic("Error while scaning routines")
		}
		routine.ID = routine_id
		routine.UserId = user_id
		routine.Title = title
		routine.Description = description
		routine.Mode = routine_mode
		routine.Time = routine_time
		if is_trash == 1 {
			routine.IsTrash = "Yes"
		} else {
			routine.IsTrash = "No"
		}
		routine.CreatedAt = created_at
		routines = append(routines, routine)
	}

	return routines
}
