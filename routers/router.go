package routers

import (
	"jobBoardApi/data"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// JobPostingRoutes - All CRUD operations for job postings
func JobPostingRoutes() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	// Routes
	e.POST("/job/save", data.SavePosting)
	e.GET("/job/getAllPostings", data.GetPostings)
	e.PUT("/job/:id", data.UpdatePosting)
	e.DELETE("/job/:id", data.DeletePosting)

	return e
}
