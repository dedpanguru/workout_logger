package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func HasJSONHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("Content-Type") == "application/json" {
			//call next
			if err = next(c); err != nil {
				c.Error(err)
			}
		}
		return nil
	}
}

func register(c echo.Context) error {
	defer dbpool.Close()
	// read json body
	var data map[string]any
	if err = json.NewDecoder(c.Request().Body).Decode(&data); err != nil {
		http.Error(c.Response().Writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	// verify request structure
	if data["username"].(string) == "" && data["password"].(string) == "" {
		http.Error(c.Response().Writer, err.Error(), http.StatusUnprocessableEntity)
		return err
	}
	// check if username is unique
	var usernames []string
	if err = dbpool.QueryRow(ctx, "select username from users").Scan(&usernames); err != nil {
		http.Error(c.Response().Writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	for _, name := range usernames {
		if name == data["username"].(string) {
			http.Error(c.Response().Writer, "Username Taken!", http.StatusBadRequest)
			return err
		}
	}

	// credentials are now valid!
	// hash password before storing
	key := []byte("secretsecretsecret")
	if os.Getenv("SECRET_KEY") != "" {
		key = []byte(os.Getenv("SECRET_KEY"))
	}
}
