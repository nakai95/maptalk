package router

import (
	"maptalk/internal/infrastructure/datastore"
	"maptalk/internal/interface/controller"
	"maptalk/internal/interface/presenter"
	"maptalk/internal/interface/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	datastore := datastore.NewDataStore("map-talk-432405")
	userPresenter := presenter.NewUserPresenter()
	userRepository := repository.NewUserRepository(datastore)
	userController := controller.NewUserController(userPresenter, userRepository)

	e.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		user, err := userController.GetUserByID(id)
		if err != nil {
			return c.JSON(500, err)
		}
		return c.JSON(200, user)
	})

	e.POST("/users", func(c echo.Context) error {
		ctx := c.Request().Context()
		input := new(controller.UserInputData)
		if err := c.Bind(input); err != nil {
			return err
		}
		user, err := userController.Save(*input, ctx)
		if err != nil {
			return c.JSON(500, err)
		}
		return c.JSON(200, user)
	})

	return e
}
