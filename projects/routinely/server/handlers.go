package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"routinely/routine"
	"routinely/user"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Response struct {
	HasError bool
	Message  string
	Data     any
}

const (
	JWT_SIGNING_SECRET = "STAUNAN@ROUTINELY"
)

func anyToInt64(value any) int64 {
	return int64(value.(float64))
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

func getLoggedInUserId(c echo.Context) int64 {
	user_id := echo.Context.Get(c, "user_id")
	return user_id.(int64)
}

func readCookie(c echo.Context, name string) (string, error) {
	cookie, err := c.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func setAuthenticationCookie(c echo.Context, token string) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.SameSite = http.SameSiteNoneMode
	cookie.Secure = true
	cookie.Path = "/"
	c.SetCookie(cookie)
}

func verifyAndParseToken(tokenString string) (user.User, error) {
	var loggedin_user user.User
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(JWT_SIGNING_SECRET), nil
	})
	if err != nil {
		return loggedin_user, err
	}
	for key, val := range claims {
		if key == "fullname" {
			loggedin_user.FullName = val.(string)
		} else if key == "username" {
			loggedin_user.Username = val.(string)
		} else if key == "email" {
			loggedin_user.Email = val.(string)
		} else if key == "id" {
			loggedin_user.ID = int64(val.(float64))
		}
	}
	return loggedin_user, nil
}

func jwtTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	var protected_routes []string = []string{"/user_details", "/upload_photo_registration_step", "/skip_upload_photo_in_registration_step", "/logout", "/create_routine", "/routine_details", "/all_routines", "/mark_routine_as_done", "/mark_routine_as_not_done", "/progress"}

	return func(c echo.Context) error {
		// Retrieve the token from the cookie
		existed_token, err := readCookie(c, "token")
		if err != nil {
			// Token not found --
			url_path := c.Request().URL.Path
			if slices.Contains(protected_routes, url_path) {
				// Return Response --
				var response Response
				response.HasError = true
				response.Message = "token not found"
				response.Data = nil
				return c.JSON(http.StatusUnauthorized, response)
			}
			return next(c)
		}
		// Verify the token
		loggedin_user, err := verifyAndParseToken(existed_token)
		if err != nil {
			// Token verification failed --
			url_path := c.Request().URL.Path
			if slices.Contains(protected_routes, url_path) {
				// Return Response --
				var response Response
				response.HasError = true
				response.Message = "token verification failed"
				response.Data = nil
				return c.JSON(http.StatusUnauthorized, response)
			}
			return next(c)
		}
		// Set user id to context --
		echo.Context.Set(c, "user_id", loggedin_user.ID)
		return next(c)
	}
}

func CreateAccountHandler(c echo.Context) error {
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
	setAuthenticationCookie(c, token)

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
}

func UploadPhotoInRegistrationStepHandler(c echo.Context) error {
	var user_id int64 = 1
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
		response.Message = err.Error()
		response.Data = err
		return c.JSON(http.StatusOK, response)
	}
	if success {
		update_success, err := user.UpdateRegistrationStep(user_id, 2)
		if err != nil {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = err.Error()
			response.Data = err
			return c.JSON(http.StatusOK, response)
		}
		// Return Response --
		var response Response
		response.HasError = update_success
		response.Message = "Display picture has been updated"
		response.Data = success
		return c.JSON(http.StatusOK, response)
	} else {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = "Something went wrong!"
		response.Data = nil
		return c.JSON(http.StatusOK, response)
	}
}

func SkipUploadPhotoInRegistrationStepHandler(c echo.Context) error {
	var user_id int64 = 1
	update_success, err := user.UpdateRegistrationStep(user_id, 2)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = err.Error()
		response.Data = err
		return c.JSON(http.StatusOK, response)
	}
	// Return Response --
	var response Response
	response.HasError = update_success
	response.Message = "Successfully skipped the step: 'upload display photo'. registration successful!"
	response.Data = nil
	return c.JSON(http.StatusOK, response)
}

func UserDetailsHandler(c echo.Context) error {
	user_id := getLoggedInUserId(c)
	if user_id == 0 {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = "Login required"
		response.Data = nil
		return c.JSON(http.StatusOK, response)
	}

	user_details, err := user.GetUserDetailsById(user_id)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = err.Error()
		response.Data = nil
		return c.JSON(http.StatusOK, response)
	}
	// Return Response --
	var response Response
	response.HasError = false
	response.Message = "Successfully retrieved user details"
	response.Data = user_details
	return c.JSON(http.StatusOK, response)
}

func LoginHandler(c echo.Context) error {
	// Check if user is already logged in --
	existed_token, err := readCookie(c, "token")
	if err == nil && existed_token != "" {
		// Check if token is valid --
		loggedInUser, err := verifyAndParseToken(existed_token)
		if err != nil {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = "Invalid token, please login"
			response.Data = nil
			return c.JSON(http.StatusOK, response)
		}
		if loggedInUser.ID > 0 {
			// Return Response --
			var response Response
			response.HasError = true
			response.Message = "User already logged in"
			response.Data = nil
			return c.JSON(http.StatusOK, response)
		}
	}

	// Get Request Data --
	var reqData map[string]any = getRequestData(c)
	var login_user user.User
	if reqData["email"] == nil {
		panic("Email is required!")
	} else {
		login_user.Email = reqData["email"].(string)
	}
	if reqData["password"] == nil {
		panic("Password is required!")
	} else {
		login_user.SecretPassword = reqData["password"].(string)
	}

	// Validate Login Data --
	result_user, err := user.ValidateLoginUser(login_user)
	if err != nil {
		// Return Response --
		var response Response
		response.Message = "Error while validating login data"
		if err.Error() == "user not found" {
			response.Message = "User :'" + login_user.Email + "' doesn't exists"
		} else if err.Error() == "incorrect password" {
			response.Message = "Your password is incorrect, please check and try again!"
		} else if err.Error() == "invalid email format" {
			response.Message = "Email :'" + login_user.Email + "' is invalid, please provide a valid email address!"
		} else if err.Error() == "invalid password format" {
			response.Message = "Password format is invalid"
		}
		response.HasError = true
		response.Data = err.Error()
		return c.JSON(http.StatusOK, response)
	}
	// Get User Details --
	user_details, err := user.GetUserDetailsById(result_user.ID)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = "Error while getting user details"
		response.Data = err.Error()
		return c.JSON(http.StatusOK, response)
	}
	// Generate JWT Token --
	token, err := user.GenerateJWTToken(user_details)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = "Error while generating jwt token for user"
		response.Data = err.Error()
		return c.JSON(http.StatusOK, response)
	}
	// Store token to cookie and user object--
	setAuthenticationCookie(c, token)
	user_details.JWTToken = token

	// Return Response --
	var response Response
	response.HasError = false
	response.Message = "Logged In!"
	response.Data = user_details
	return c.JSON(http.StatusOK, response)
}

func LogoutHandler(c echo.Context) error {
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
}

func CreateRoutineHandler(c echo.Context) error {
	var routine_details routine.Routine

	// Get Request Data --
	var reqData map[string]any = getRequestData(c)

	// Create Routine Object
	var routineObj routine.Routine
	routineObj.UserId = getLoggedInUserId(c)
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
	routineObj.Status = routine.ROUTINE_STATUS_ACTIVE
	last_inserted_id, err := routine.CreateRoutine(routineObj)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = err.Error()
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
}

func UpdateRoutineHandler(c echo.Context) error {
	// Get Request Data --
	var reqData map[string]any = getRequestData(c)

	// Create Routine Object
	var routineObj routine.Routine
	routineObj.UserId = getLoggedInUserId(c)
	if reqData["id"] == nil {
		routineObj.ID = 0
	} else {
		routineObj.ID = anyToInt64(reqData["id"])
	}
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
	updated_routine_details, err := routine.UpdateRoutine(routineObj)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = err.Error()
		response.Data = nil
		return c.JSON(http.StatusOK, response)
	}

	// Return Response --
	var response Response
	response.HasError = false
	response.Message = "Routine has been updated sucessfully"
	response.Data = updated_routine_details
	return c.JSON(http.StatusOK, response)
}

func UpdateRoutineStatusHandler(c echo.Context) error {
	// Get Request Data --
	var reqData map[string]any = getRequestData(c)

	// Create Routine Object
	var routineObj routine.Routine
	routineObj.UserId = getLoggedInUserId(c)

	if reqData["id"] == nil {
		routineObj.ID = 0
	} else {
		routineObj.ID = anyToInt64(reqData["id"])
	}
	if reqData["status"] == nil {
		routineObj.Status = routine.ROUTINE_STATUS_ACTIVE
	} else {
		routineObj.Status = reqData["status"].(string)
	}

	success, err := routine.UpdateRoutineStatus(routineObj)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = err.Error()
		response.Data = nil
		return c.JSON(http.StatusOK, response)
	}

	// Return Response --
	var response Response
	response.HasError = false
	response.Message = "Status Updated!"
	response.Data = success
	return c.JSON(http.StatusOK, response)
}

func MoveToTrashHandler(c echo.Context) error {
	// Get Request Data --
	var reqData map[string]any = getRequestData(c)

	// Create Routine Object
	var routineObj routine.Routine
	routineObj.UserId = getLoggedInUserId(c)

	if reqData["id"] == nil {
		routineObj.ID = 0
	} else {
		routineObj.ID = anyToInt64(reqData["id"])
	}

	success, err := routine.MoveToTrash(routineObj)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = err.Error()
		response.Data = nil
		return c.JSON(http.StatusOK, response)
	}

	// Return Response --
	var response Response
	response.HasError = false
	response.Message = "Routine has been successfully moved to trash!"
	response.Data = success
	return c.JSON(http.StatusOK, response)
}

func RestoreFromTrashHandler(c echo.Context) error {
	// Get Request Data --
	var reqData map[string]any = getRequestData(c)

	// Create Routine Object
	var routineObj routine.Routine
	routineObj.UserId = getLoggedInUserId(c)

	if reqData["id"] == nil {
		routineObj.ID = 0
	} else {
		routineObj.ID = anyToInt64(reqData["id"])
	}

	success, err := routine.RestoreFromTrash(routineObj)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = err.Error()
		response.Data = nil
		return c.JSON(http.StatusOK, response)
	}

	// Return Response --
	var response Response
	response.HasError = false
	response.Message = "Routine has been successfully restored from trash!"
	response.Data = success
	return c.JSON(http.StatusOK, response)
}

func DeleteRoutineForeverHandler(c echo.Context) error {
	// Get Request Data --
	var reqData map[string]any = getRequestData(c)

	// Create Routine Object
	var routineObj routine.Routine
	routineObj.UserId = getLoggedInUserId(c)

	if reqData["id"] == nil {
		routineObj.ID = 0
	} else {
		routineObj.ID = anyToInt64(reqData["id"])
	}

	success, err := routine.DeleteRoutineForever(routineObj)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = err.Error()
		response.Data = nil
		return c.JSON(http.StatusOK, response)
	}

	// Return Response --
	var response Response
	response.HasError = false
	response.Message = "Routine has been successfully removed!"
	response.Data = success
	return c.JSON(http.StatusOK, response)
}

func VerifyRoutineTitleHandler(c echo.Context) error {
	// Get Request Data --
	var reqData map[string]any = getRequestData(c)
	var title string = reqData["title"].(string)
	user_id := getLoggedInUserId(c)

	slug_exists, err := routine.VerifyRoutineTitle(title, user_id)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = "Something went wrong!"
		response.Data = err.Error()
		return c.JSON(http.StatusOK, response)
	}

	// Return Response --
	var response Response
	response.HasError = false
	response.Message = ""
	response.Data = slug_exists
	return c.JSON(http.StatusOK, response)
}

func RoutineDetailsHandler(c echo.Context) error {
	var user_id int64 = getLoggedInUserId(c)
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
}

func RoutineHistoryHandler(c echo.Context) error {
	// Get Request Data --
	var reqData map[string]any = getRequestData(c)
	var routine_id int64 = 0

	var user_id int64 = getLoggedInUserId(c)
	if reqData["id"] == nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = "Routine ID is required"
		response.Data = nil
		return c.JSON(http.StatusOK, response)
	} else {
		routine_id = anyToInt64(reqData["id"])
	}

	histories, err := routine.GetRoutineHistories(user_id, routine_id)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = err.Error()
		response.Data = nil
		return c.JSON(http.StatusOK, response)
	}

	// Return Response --
	var response Response
	response.HasError = false
	response.Message = "Successfully retrieved routine history!"
	response.Data = histories
	return c.JSON(http.StatusOK, response)
}

func GetAllRoutinesHandler(c echo.Context) error {
	// Get Request Data --
	var reqData map[string]any = getRequestData(c)
	var page int
	var search string
	var routine_mode string
	if reqData["page"] == nil {
		page = 1
	} else {
		page = int(reqData["page"].(float64))
	}
	if reqData["search"] == nil {
		search = ""
	} else {
		search = reqData["search"].(string)
	}
	if reqData["mode"] == nil {
		routine_mode = ""
	} else {
		routine_mode = reqData["mode"].(string)
	}

	var user_id int64 = getLoggedInUserId(c)
	routines, err := routine.GetRoutinesByPage(user_id, page, search, routine_mode)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = "Unable to retrieve routine list"
		response.Data = nil
		return c.JSON(http.StatusOK, response)
	}

	// Return Response --
	var response Response
	response.HasError = false
	response.Message = "Successfully retrieved list"
	response.Data = routines
	return c.JSON(http.StatusOK, response)
}

func GetTrashedRoutinesHandler(c echo.Context) error {
	// Get Request Data --
	var reqData map[string]any = getRequestData(c)
	var page int
	if reqData["page"] == nil {
		page = 1
	} else {
		page = int(reqData["page"].(float64))
	}

	var user_id int64 = getLoggedInUserId(c)
	routines, err := routine.GetTrashedRoutinesByPage(user_id, page)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = "Unable to retrieve trashed items"
		response.Data = nil
		return c.JSON(http.StatusOK, response)
	}

	// Return Response --
	var response Response
	response.HasError = false
	response.Message = "Successfully retrieved list of trashed items"
	response.Data = routines
	return c.JSON(http.StatusOK, response)
}

func GetInboxesHandler(c echo.Context) error {
	var user_id int64 = getLoggedInUserId(c)
	inboxes, err := routine.GetInboxes(user_id)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = "Unable to retrieve inbox items"
		response.Data = nil
		return c.JSON(http.StatusOK, response)
	}

	// Return Response --
	var response Response
	response.HasError = false
	response.Message = "Successfully retrieved inboxes"
	response.Data = inboxes
	return c.JSON(http.StatusOK, response)
}

func MarkRoutineAsDoneHandler(c echo.Context) error {
	// Get Request Data --
	var reqData map[string]any = getRequestData(c)

	var routine_entry routine.RoutineEntry
	routine_entry.UserID = getLoggedInUserId(c)

	if reqData["routine_id"] == nil {
		routine_entry.RoutineID = 0
	} else {
		routine_entry.RoutineID = int64(reqData["routine_id"].(float64))
	}
	if reqData["checked_on_date"] == nil {
		routine_entry.CheckedOnDate = ""
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
}

func MarkRoutineAsNotDoneHandler(c echo.Context) error {
	// Get Request Data --
	var reqData map[string]any = getRequestData(c)

	var routine_entry routine.RoutineEntry
	routine_entry.UserID = getLoggedInUserId(c)

	if reqData["routine_id"] == nil {
		routine_entry.RoutineID = 0
	} else {
		routine_entry.RoutineID = int64(reqData["routine_id"].(float64))
	}
	if reqData["checked_on_date"] == nil {
		routine_entry.CheckedOnDate = ""
	} else {
		routine_entry.CheckedOnDate = reqData["checked_on_date"].(string)
	}
	result, err := routine.MarkRoutineAsNotDone(routine_entry)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = err.Error()
		response.Data = nil
		return c.JSON(http.StatusOK, response)
	}

	// Return Response --
	var response Response
	response.HasError = false
	response.Message = "Routine has been marked as not done"
	response.Data = result
	return c.JSON(http.StatusOK, response)
}

func GetDayProgressHandler(c echo.Context) error {
	// Get Request Data --
	var reqData map[string]any = getRequestData(c)
	var date string
	if reqData["date"] == nil {
		date = ""
	} else {
		date = reqData["date"].(string)
	}
	var user_id int64 = getLoggedInUserId(c)

	routine_entries, err := routine.GetDayProgress(user_id, date)
	if err != nil {
		// Return Response --
		var response Response
		response.HasError = true
		response.Message = err.Error()
		response.Data = routine_entries
		return c.JSON(http.StatusOK, response)
	}
	// Return Response --
	var response Response
	response.HasError = false
	response.Message = "Successfully retrieved given date's progress!"
	response.Data = routine_entries
	return c.JSON(http.StatusOK, response)
}
