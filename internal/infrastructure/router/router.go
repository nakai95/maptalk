package router

import (
	"maptalk/internal/infrastructure/datastore"
	"maptalk/internal/interface/controller"
	"maptalk/internal/interface/presenter"
	"maptalk/internal/interface/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"context"
)

type Name struct {
	Name string `json:"name"`
}

func NewRouter() *echo.Echo {
    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

	datastore, nil := datastore.NewDataStore(context.Background())
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

	e.POST("/user", func(c echo.Context) error {
		n := new(Name)
		if err := c.Bind(n); err != nil {
			return err
		}
		user, err := userController.Save(n.Name)
		if err != nil {
			return c.JSON(500, err)
		}
		return c.JSON(200, user)
	})

    return e
}
