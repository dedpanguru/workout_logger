package api

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	dbpool *pgxpool.Pool
	ctx    context.Context
	err    error
)

func NewRouter(pool *pgxpool.Pool, c context.Context) *echo.Echo {
	dbpool = pool
	ctx = c
	router := echo.New()
	router.Use(middleware.Recover())
	router.Use(middleware.Logger())
	router.GET("/", homeHandler)
	router.POST("/register", register, HasJSONHeader)
	return router
}
func homeHandler(c echo.Context) error {
	defer dbpool.Close()
	var greeting string
	err = dbpool.QueryRow(ctx, "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		http.Error(c.Response().Writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": greeting,
	})
}
