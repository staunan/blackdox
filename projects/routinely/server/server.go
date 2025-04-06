package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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
	// jwt Middleware --
	e.Use(jwtTokenMiddleware)

	// User Registration --
	e.POST("/create_account", CreateAccountHandler)
	e.POST("/upload_photo_registration_step", UploadPhotoInRegistrationStepHandler)
	e.POST("/skip_upload_photo_in_registration_step", SkipUploadPhotoInRegistrationStepHandler)
	e.GET("/user_details", UserDetailsHandler)

	// User Login --
	e.POST("/login", LoginHandler)

	// User Logout --
	e.POST("/logout", LogoutHandler)

	// Routines --
	e.POST("/create_routine", CreateRoutineHandler)
	e.POST("/verify_routine_title", VerifyRoutineTitleHandler)
	e.POST("/update_routine", UpdateRoutineHandler)
	e.POST("/all_routines", GetAllRoutinesHandler)
	e.POST("/trashed_routines", GetTrashedRoutinesHandler)

	// Routine Details --
	e.POST("/routine_details", RoutineDetailsHandler)
	e.POST("/routine_history", RoutineHistoryHandler)
	e.POST("/update_routine_status", UpdateRoutineStatusHandler)
	e.POST("/move_to_trash", MoveToTrashHandler)
	e.POST("/restore_from_trash", RestoreFromTrashHandler)
	e.POST("/delete_routine_forever", DeleteRoutineForeverHandler)

	// Routine Entries
	e.POST("/mark_routine_as_done", MarkRoutineAsDoneHandler)
	e.POST("/mark_routine_as_not_done", MarkRoutineAsNotDoneHandler)

	// Get Progress --
	e.GET("/progress", GetProgressHandler)
	e.POST("/day_progress", GetDayProgressHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
