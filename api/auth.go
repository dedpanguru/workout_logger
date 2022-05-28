package api

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/dedpanguru/workout_logger/database"
	"github.com/labstack/echo/v4"
)

func register(ctx echo.Context) error {
	// validate JSON body
	var input database.User
	if err = json.NewDecoder(ctx.Request().Body).Decode(&input); err != nil {
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "Invalid JSON body",
		}
	}

	// hash password
	hasher := sha512.New()
	hasher.Write([]byte(input.Password))
	hashedPassword := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	// store in DB
	if _, err = db.Exec(ctx.Request().Context(), "insert into users(name, password) values($1, $2);", input.Name, hashedPassword); err != nil {
		if strings.Split(err.Error(), " ")[1] == "duplicate" {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: "Username Taken!",
			}
		}
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return ctx.JSON(http.StatusCreated, echo.Map{
		"message": "Account Successfully Created!",
	})
}

func getKey() []byte {
	if os.Getenv("SECRET_KEY") != "" {
		return []byte(os.Getenv("SECRET_KEY"))
	}
	return []byte("secretsecretsecret")
}
