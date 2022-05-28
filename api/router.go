package api

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	db  *pgxpool.Pool
	err error
)

func NewRouter(pool *pgxpool.Pool) *echo.Echo {
	db = pool
	router := echo.New()
	router.Use(middleware.Recover())
	router.Use(middleware.Logger())
	router.GET("/", homeHandler)
	router.POST("/register", register)
	return router
}
func homeHandler(c echo.Context) error {
	var greeting string
	err = db.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Printf("Problem with connecting to DB: %s\n", err.Error())
		return &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": greeting,
	})
}
