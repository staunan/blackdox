package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"routinely/mysqldb"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
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
	ID             int64
	FullName       string
	Email          string
	Username       string
	SecretPassword string
	CreatedAt      string
}

func CreateUser(user User) (int64, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		return 0, errors.New("unable to connect to db")
	}

	// Preparing SQL statement --
	query := "INSERT INTO `users` (full_name, email, username, secret_password) VALUES (?, ?, ?, ?);"
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

	// Check if user already exists --
	user_exists := checkIfUserAlreadyExists(email)
	if user_exists {
		return 0, errors.New("user already exists")
	}
	// Execute DB Query --
	dbResponse, err := insert.Exec(fullname, email, username, secret_password)
	if err != nil {
		return 0, errors.New("unable to execute query")
	}
	insert.Close()

	// Return last inserted ID --
	return dbResponse.LastInsertId()
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
		return user, errors.New("Unable to connect to db")
	}

	// Get User Details from DB --
	var user_id_str string = strconv.Itoa(int(user_id))
	row := db.QueryRow("SELECT id, full_name, email, username, secret_password, created_at FROM users where id = ?", user_id_str)
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

	// Scan fields --
	err := row.Scan(&id, &full_name, &email, &username, &secret_password, &created_at)
	if err != nil {
		return user, err
	}
	// Create a routine object --
	user.ID = id
	user.FullName = full_name
	user.Email = email
	user.Username = username
	user.SecretPassword = secret_password
	user.CreatedAt = created_at
	return user, nil
}

func Login(c echo.Context) (string, error) {
	// Connect to db --
	db, err := mysqldb.ConnectMySQL()
	if err != nil {
		panic("Unable to connect to db")
	}

	// Get Request Data --
	var reqData map[string]any = getRequestData(c)
	var email string = ""
	var password string = ""
	if reqData["email"] == nil {
		panic("Email is required!")
	} else {
		email = reqData["email"].(string)
	}
	if reqData["password"] == nil {
		panic("Password is required!")
	} else {
		password = reqData["password"].(string)
	}

	// Check if user present in database --
	var user_id int64
	var fullname string
	var username string
	var hashed_password string
	query_err := db.QueryRow("SELECT id, full_name, username, secret_password FROM users where email = ?", email).Scan(&user_id, &fullname, &username, &hashed_password)
	switch {
	case query_err == sql.ErrNoRows:
		return "", errors.New("user not found")
	case query_err != nil:
		return "", query_err
	default:
		if !Verify(hashed_password, password) {
			return "", errors.New("incorrect password")
		}
		// Login user --
		// Set custom claims
		claims := &jwtCustomClaims{
			user_id,
			fullname,
			username,
			email,
			true,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}
		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(JWT_SIGNING_SECRET))
		if err != nil {
			return "", err
		}
		return t, nil
		// return c.JSON(http.StatusOK, echo.Map{
		// 	"token": t,
		// })
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

func getRequestData(c echo.Context) map[string]any {
	json_map := make(map[string]any)
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return json_map
	} else {
		return json_map
	}
}
