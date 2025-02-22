package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

const (
	username = "root"
	password = "0000"
	hostname = "127.0.0.1:3306"
	dbname   = "routinely"
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

var db *sql.DB

func main() {
	e := echo.New()
	// Create a DB Connection --
	cfg := mysql.Config{
		User:   username,
		Passwd: password,
		Net:    "tcp",
		Addr:   hostname,
		DBName: dbname,
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
		return c.JSON(http.StatusOK, getRoutines(user_id))
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func getRoutines(user_id int64) []Routine {
	var routine_id int64
	var title string
	var description string
	var routine_mode string
	var routine_time string
	var is_trash int8
	var created_at string

	var routines []Routine

	// Prepare statement for reading data
	var user_id_str string = strconv.Itoa(int(user_id))
	rows, err := db.Query("SELECT id, routine_title, routine_description, routine_mode, routine_time, is_trash, created_at FROM routines where user_id = ?", user_id_str)
	if err != nil {
		return routines
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var routine Routine
		if err := rows.Scan(&routine_id, &title, &description, &routine_mode, &routine_time, &is_trash, &created_at); err != nil {
			return routines
		}
		routine.ID = routine_id
		routine.UserId = user_id
		routine.Title = title
		routine.Description = description
		routine.Mode = routine_mode
		routine.Time = routine_time
		if is_trash == 0 {
			routine.IsTrash = "No"
		} else {
			routine.IsTrash = "Yes"
		}
		routine.CreatedAt = created_at
		routines = append(routines, routine)
	}
	return routines
}
