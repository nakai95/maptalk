package router

import (
	"maptalk/internal/domain/usecase"
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

	userPresenter := presenter.NewUserPresenter()
    userRepository := repository.NewUserRepository()
	userUseCase := usecase.NewUserUseCase(userPresenter, userRepository)
	userController := controller.NewUserController(userUseCase)
    

    e.GET("/users/:id", func(c echo.Context) error {
	    id := c.Param("id")
        user, err := userController.GetUserByID(id)
        if err != nil {
            return c.JSON(500, err)
        }
        return c.JSON(200, user)
    })

    return e
}
