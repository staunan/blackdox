package user

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/mail"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"routinely/mysqldb"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	JWT_SIGNING_SECRET = "STAUNAN@ROUTINELY"
)

type jwtCustomClaims struct {
	ID       int64  `json:"id"`
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Admin    bool   `json:"admin"`
	jwt.RegisteredClaims
}

type User struct {
	ID                         int64
	FullName                   string
	Email                      string
	Username                   string
	SecretPassword             string
	DisplayPictureName         string
	RegistrationStepsCompleted int8
	RegistrationSuccessful     bool
	JWTToken                   string
	CreatedAt                  string
}

func CreateUser(user User) (int64, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return 0, errors.New("unable to connect to db")
	}

	// Preparing SQL statement --
	query := "INSERT INTO `users` (full_name, email, username, secret_password, registration_steps_completed, registration_successful) VALUES (?, ?, ?, ?, ?, ?);"
	insert, err := db.Prepare(query)
	if err != nil {
		return 0, errors.New("unable to prepare query")
	}

	// Finalizing DB Column values --
	// fullname --
	var fullname string = user.FullName
	if fullname == "" {
		return 0, errors.New("fullname should be present")
	} else if len(fullname) > 50 {
		return 0, errors.New("fullname is too big")
	}
	// Email --
	var email string = user.Email
	if email == "" {
		return 0, errors.New("email should be present")
	} else if len(email) > 50 {
		return 0, errors.New("email is too big")
	}
	// Username --
	var username string = user.Username
	if username == "" {
		return 0, errors.New("username should be present")
	} else if len(username) > 50 {
		return 0, errors.New("username is too big")
	}
	// Secret Password --
	var secret_password string = user.SecretPassword
	if secret_password == "" {
		return 0, errors.New("secret password should be present")
	} else if len(secret_password) > 50 {
		return 0, errors.New("secret password is too big")
	}
	hashed_password, err := Hash(secret_password)
	if err != nil {
		return 0, err
	}

	// Check if user already exists --
	user_exists := checkIfUserAlreadyExists(email)
	if user_exists {
		return 0, errors.New("user already exists")
	}
	// Execute DB Query --
	dbResponse, err := insert.Exec(fullname, email, username, hashed_password, 1, 0)
	if err != nil {
		return 0, err
	}
	insert.Close()

	// Return last inserted ID --
	return dbResponse.LastInsertId()
}

func GenerateJWTToken(user_details User) (string, error) {
	// Set custom claims
	claims := &jwtCustomClaims{
		user_details.ID,
		user_details.FullName,
		user_details.Username,
		user_details.Email,
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	return token.SignedString([]byte(JWT_SIGNING_SECRET))
}

func UpdateDisplayPicture(file *multipart.FileHeader, user_id int64) (bool, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return false, errors.New("unable to connect to db")
	}

	src, err := file.Open()
	if err != nil {
		return false, err
	}
	defer src.Close()

	var now_time string = time.Now().Format("20060102150405")
	var file_extension string = getExtensionFromFileName(file.Filename)
	var user_id_str string = strconv.Itoa(int(user_id))
	file_name := "dp_" + now_time + user_id_str + file_extension

	// Destination
	dst, err := os.Create("images/user_profile_pictures/" + file_name)
	if err != nil {
		return false, err
	}
	defer dst.Close()
	fmt.Printf("%#v\n", dst)

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return false, err
	}

	// Delete previous image --
	var prev_file_name string
	prev_dp_row := db.QueryRow("SELECT display_picture_name FROM users where id = ?", user_id_str)
	// Scan fields --
	scan_err := prev_dp_row.Scan(&prev_file_name)
	if scan_err != nil {
		return false, scan_err
	}
	if prev_file_name != "" {
		// Remove image from folder --
		remove_error := os.Remove("images/user_profile_pictures/" + prev_file_name)
		if remove_error != nil {
		}
	}

	// Update database --
	query := "UPDATE users SET display_picture_name = ? WHERE id = ?"
	update_result, err := db.Exec(query, file_name, user_id_str)
	if err != nil {
		return false, err
	}
	rows_updated, err := update_result.RowsAffected()
	if err != nil {
		return false, err
	}
	if rows_updated > 0 {
		return true, nil
	} else {
		return false, errors.New("something went wrong")
	}
}

func getExtensionFromFileName(filename string) string {
	name_arr := strings.Split(filename, ".")
	ext := name_arr[len(name_arr)-1]
	return "." + ext
}

func checkIfUserAlreadyExists(email string) bool {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		panic("Unable to connect to db")
	}

	// Check if data is already present in database --
	var exsisted_row_id int64
	query_err := db.QueryRow("SELECT id FROM users where email = ?", email).Scan(&exsisted_row_id)
	switch {
	case query_err == sql.ErrNoRows:
		return false
	case query_err != nil:
		panic(query_err)
	default:
		return true
	}
}

func GetUserDetailsById(user_id int64) (User, error) {
	var user User
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return user, errors.New("unable to connect to db")
	}

	// Get User Details from DB --
	var user_id_str string = strconv.Itoa(int(user_id))
	row := db.QueryRow("SELECT id, full_name, email, username, secret_password, display_picture_name, registration_steps_completed, registration_successful, created_at FROM users where id = ?", user_id_str)
	return mapDBDataToUserDetails(row)
}

func mapDBDataToUserDetails(row *sql.Row) (User, error) {
	var user User
	// Declaring variable --
	var id int64
	var full_name string
	var email string
	var username string
	var secret_password string
	var created_at string
	var display_picture_name string
	var registration_steps_completed int8
	var registration_successful int8

	// Scan fields --
	err := row.Scan(&id, &full_name, &email, &username, &secret_password, &display_picture_name, &registration_steps_completed, &registration_successful, &created_at)
	if err != nil {
		return user, err
	}
	// Create a routine object --
	user.ID = id
	user.FullName = full_name
	user.Email = email
	user.Username = username
	user.SecretPassword = secret_password
	user.DisplayPictureName = display_picture_name
	user.RegistrationStepsCompleted = registration_steps_completed
	if registration_successful == 0 {
		user.RegistrationSuccessful = false
	} else if registration_successful == 1 {
		user.RegistrationSuccessful = true
	}
	user.CreatedAt = created_at
	return user, nil
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validateUsername(username string) (bool, error) {
	if len(username) > 20 {
		return false, errors.New("username should not be greater than 20 character")
	}
	return true, nil
}

func isPasswordValid(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(password) >= 8 {
		hasMinLen = true
	}
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func ValidateLoginUser(user_details User) (User, error) {
	var result_user User
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return result_user, err
	}

	// Validate data --
	if !isEmailValid(user_details.Email) {
		return result_user, errors.New("invalid email format")
	}
	if !isPasswordValid(user_details.SecretPassword) {
		return result_user, errors.New("invalid password format")
	}

	// Check if user present in database --
	var user_id int64
	var hashed_password string
	query_err := db.QueryRow("SELECT id, secret_password FROM users where email = ?", user_details.Email).Scan(&user_id, &hashed_password)
	switch {
	case query_err == sql.ErrNoRows:
		return result_user, errors.New("user not found")
	case query_err != nil:
		return result_user, query_err
	default:
		if !Verify(hashed_password, user_details.SecretPassword) {
			return result_user, errors.New("incorrect password")
		} else {
			result_user.ID = user_id
			return result_user, nil
		}
	}
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Verify(hashed, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

func UpdateRegistrationStep(user_id int64, step_number int) (bool, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return false, errors.New("unable to connect to db")
	}

	query := "UPDATE users SET registration_steps_completed = ?, registration_successful = ? WHERE id = ?"
	var user_id_str string = strconv.Itoa(int(user_id))
	if step_number == 0 {
		return false, errors.New("step 0 not allowed")
	} else if step_number == 1 {
		update_result, err := db.Exec(query, step_number, false, user_id_str)
		if err != nil {
			return false, err
		}
		rows_updated, err := update_result.RowsAffected()
		if err != nil {
			return false, err
		}
		return rows_updated > 0, nil
	} else if step_number == 2 {
		update_result, err := db.Exec(query, step_number, 1, user_id_str)
		if err != nil {
			return false, err
		}
		rows_updated, err := update_result.RowsAffected()
		if err != nil {
			return false, err
		}
		return rows_updated > 0, nil
	} else {
		return false, errors.New("invalid step number")
	}
}

func CheckIfEmailAvailable(user User) (bool, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return false, err
	}

	// Validate data --
	if !isEmailValid(user.Email) {
		return false, errors.New("invalid email format")
	}

	// Check if user present in database --
	var user_id int64
	query_err := db.QueryRow("SELECT id FROM users where email = ?", user.Email).Scan(&user_id)
	switch {
	case query_err == sql.ErrNoRows:
		return true, nil
	case query_err != nil:
		return false, query_err
	default:
		return false, nil
	}
}

func UpdateUserEmail(user User) (bool, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return false, err
	}

	// Validate data --
	if !isEmailValid(user.Email) {
		return false, errors.New("invalid email format")
	}

	emailAvailable, err := CheckIfEmailAvailable(user)
	if err != nil {
		return false, err
	}
	if emailAvailable {
		query := "UPDATE users SET email = ? WHERE id = ?"
		update_result, err := db.Exec(query, user.Email, user.ID)
		if err != nil {
			return false, err
		}
		rows_updated, err := update_result.RowsAffected()
		if err != nil {
			return false, err
		}
		return rows_updated > 0, nil
	} else {
		return false, errors.New("email not available")
	}
}

func CheckIfUsernameAvailable(user User) (bool, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return false, err
	}

	// Validate data --
	valid, err := validateUsername(user.Username)
	if err != nil {
		return false, err
	}

	if valid {
		// Check if user present in database --
		var user_id int64
		query_err := db.QueryRow("SELECT id FROM users where username = ?", user.Username).Scan(&user_id)
		switch {
		case query_err == sql.ErrNoRows:
			return true, nil
		case query_err != nil:
			return false, query_err
		default:
			return false, nil
		}
	} else {
		return false, errors.New("username is not valid")
	}
}
