package main

import (
	"encoding/json"
	"net/http"
	"time"

	"routinely/routine"
	"routinely/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Response struct {
	HasError bool
	Message  string
	Data     any
}

func main() {
	e := echo.New()
	// Public Folder --
	e.Static("/images", "images")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlAllowCredentials, echo.HeaderAccessControlAllowOrigin, echo.HeaderContentType},
	}))

	// Create Routine --
	e.POST("/create_account", func(c echo.Context) error {
		// Get Request Data --
		var reqData map[string]any = getRequestData(c)

		existed_token, err := readCookie(c, "token")
		if err == nil && existed_token != "" {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = "Already signed in"
			response.Data = nil
			return c.JSON(http.StatusOK, response)
		}

		// Create User Object
		var userObj user.User
		// fullname --
		if reqData["fullname"] == nil {
			userObj.FullName = ""
		} else {
			userObj.FullName = reqData["fullname"].(string)
		}
		// username --
		if reqData["username"] == nil {
			userObj.Username = ""
		} else {
			userObj.Username = reqData["username"].(string)
		}
		// email --
		if reqData["email"] == nil {
			userObj.Email = ""
		} else {
			userObj.Email = reqData["email"].(string)
		}
		// password --
		if reqData["password"] == nil {
			userObj.SecretPassword = ""
		} else {
			userObj.SecretPassword = reqData["password"].(string)
		}
		last_inserted_id, err := user.CreateUser(userObj)
		if err != nil {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = "Unable to create user"
			response.Data = err.Error()
			return c.JSON(http.StatusOK, response)
		}
		userObj.ID = last_inserted_id

		// Generate JWT Token for user --
		token, err := user.GenerateJWTToken(userObj)
		if err != nil {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = "Error while generating JWT Token"
			response.Data = err.Error()
			return c.JSON(http.StatusOK, response)
		}

		// Set Cookie --
		cookie := new(http.Cookie)
		cookie.Name = "token"
		cookie.Value = token
		cookie.Expires = time.Now().Add(24 * time.Hour)
		cookie.SameSite = http.SameSiteNoneMode
		cookie.Secure = true
		cookie.Path = "/"
		c.SetCookie(cookie)

		user_details, err := user.GetUserDetailsById(last_inserted_id)
		if err != nil {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = "Unable to get user details"
			response.Data = err
			return c.JSON(http.StatusOK, response)
		}

		// Return Response --
		var response Response
		response.HasError = false
		response.Message = "Account has been created sucessfully"
		response.Data = user_details
		return c.JSON(http.StatusOK, response)
	})

	e.POST("/logout", func(c echo.Context) error {
		// Logout by setting token to empty string --
		cookie := new(http.Cookie)
		cookie.Name = "token"
		cookie.Value = ""
		cookie.MaxAge = -1 // Setting negetive value to MaxAge makes the cookie expired, which is another way to delete the cookie
		c.SetCookie(cookie)

		// Return Response --
		var response Response
		response.HasError = false
		response.Message = "Logged Out!"
		response.Data = nil
		return c.JSON(http.StatusOK, response)
	})

	// Change Profile Picture --
	e.POST("/upload_photo", func(c echo.Context) error {
		// Source
		file, err := c.FormFile("file")
		if err != nil {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = "Please provide image"
			response.Data = nil
			return c.JSON(http.StatusOK, response)
		}
		success, err := user.UpdateDisplayPicture(file)
		if err != nil {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = "Error while updating profile picture"
			response.Data = err
			return c.JSON(http.StatusOK, response)
		}

		// Return Response --
		var response Response
		response.HasError = false
		response.Message = "Display picture has been updated"
		response.Data = success
		return c.JSON(http.StatusOK, response)
	})

	// Create Routine --
	e.POST("/create_routine", func(c echo.Context) error {
		var routine_details routine.Routine

		// Get Request Data --
		var reqData map[string]any = getRequestData(c)

		// Create Routine Object
		var routineObj routine.Routine
		routineObj.UserId = 1
		if reqData["title"] == nil {
			routineObj.Title = ""
		} else {
			routineObj.Title = reqData["title"].(string)
		}
		if reqData["description"] == nil {
			routineObj.Description = ""
		} else {
			routineObj.Description = reqData["description"].(string)
		}
		if reqData["mode"] == nil {
			routineObj.Mode = ""
		} else {
			routineObj.Mode = reqData["mode"].(string)
		}
		if reqData["days"] == nil {
			routineObj.DailyBasisDays = ""
		} else {
			routineObj.DailyBasisDays = reqData["days"].(string)
		}
		if reqData["weekday"] == nil {
			routineObj.WeeklyBasisWeekDays = ""
		} else {
			routineObj.WeeklyBasisWeekDays = reqData["weekday"].(string)
		}
		if reqData["monthday"] == nil {
			routineObj.MonthlyBasisDate = 0
		} else {
			routineObj.MonthlyBasisDate = reqData["monthday"].(int8)
		}
		if reqData["yearlymonthdate"] == nil {
			routineObj.YearlyBasisMonthDate = ""
		} else {
			routineObj.YearlyBasisMonthDate = reqData["yearlymonthdate"].(string)
		}
		if reqData["time"] == nil {
			routineObj.Time = ""
		} else {
			routineObj.Time = reqData["time"].(string)
		}
		routineObj.IsTrash = 0
		last_inserted_id, err := routine.CreateRoutine(routineObj)
		if err != nil {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = "Unable to create routine"
			response.Data = nil
			return c.JSON(http.StatusOK, response)
		}
		routine_details = routine.GetRoutineDetailsById(last_inserted_id)

		// Return Response --
		var response Response
		response.HasError = false
		response.Message = "Routine has been created sucessfully"
		response.Data = routine_details
		return c.JSON(http.StatusOK, response)
	})

	// Get Routine Details --
	e.POST("/routine_details", func(c echo.Context) error {
		var user_id int64 = 1
		var slug string

		// Get Request Data --
		var reqData map[string]any = getRequestData(c)

		if reqData["routine_slug"] == nil {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = "slug required"
			response.Data = nil
			return c.JSON(http.StatusOK, response)
		} else {
			slug = reqData["routine_slug"].(string)
		}

		routine_details := routine.GetRoutineDetailsBySlug(user_id, slug)

		// Return Response --
		var response Response
		response.HasError = false
		response.Message = "Successfully retrieved routine data!"
		response.Data = routine_details
		return c.JSON(http.StatusOK, response)
	})

	// Get all routines --
	e.POST("/all_routines", func(c echo.Context) error {
		var user_id int64 = 1
		var routines []routine.Routine = routine.GetRoutines(user_id)

		// Return Response --
		var response Response
		response.HasError = false
		response.Message = "Successfully retrieved list"
		response.Data = routines
		return c.JSON(http.StatusOK, response)
	})

	// Mark a Routine as Done --
	e.POST("/mark_routine_as_done", func(c echo.Context) error {
		// Get Request Data --
		var reqData map[string]any = getRequestData(c)

		var routine_entry routine.RoutineEntry
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

		new_routine_entry, err := routine.MarkRoutineAsDone(routine_entry)
		if err != nil {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = "Unable to mark as done"
			response.Data = err
			return c.JSON(http.StatusOK, response)
		}

		// Return Response --
		var response Response
		response.HasError = false
		response.Message = "Routine has been marked as done"
		response.Data = new_routine_entry
		return c.JSON(http.StatusOK, response)
	})

	// Mark a Routine as Not Done --
	e.POST("/mark_routine_as_not_done", func(c echo.Context) error {
		var routine_entry routine.RoutineEntry

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
		success, err := routine.MarkRoutineAsNotDone(routine_entry)
		if err != nil {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = "Unable to mark the routine as not done"
			response.Data = err
			return c.JSON(http.StatusOK, response)
		}

		// Return Response --
		var response Response
		response.HasError = false
		response.Message = "Routine has been marked as not done"
		response.Data = success
		return c.JSON(http.StatusOK, response)
	})

	// Get Progress --
	e.GET("/progress", func(c echo.Context) error {
		var user_id int64 = 1
		routine_entries, err := routine.GetProgress(user_id)
		if err != nil {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = "Unable to retrieve progress"
			response.Data = routine_entries
			return c.JSON(http.StatusOK, response)
		}
		// Return Response --
		var response Response
		response.HasError = false
		response.Message = "Successfully retrieved today's progress!"
		response.Data = routine_entries
		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":1323"))
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

func readCookie(c echo.Context, name string) (string, error) {
	cookie, err := c.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
