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

	datastore := datastore.NewDataStore("test-project")

	// Users
	userPresenter := presenter.NewUserPresenter()
	userRepository := repository.NewUserRepository(datastore)
	userController := controller.NewUserController(userPresenter, userRepository)

	// Posts
	postPresenter := presenter.NewPostPresenter()
	postRepository := repository.NewPostRepository(datastore)
	postController := controller.NewPostController(postPresenter, postRepository)

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

	e.POST("/posts", func(c echo.Context) error {
		ctx := c.Request().Context()
		input := new(controller.PostInputData)
		if err := c.Bind(input); err != nil {
			return err
		}

		user, err := userController.GetUserByID(input.UserId)
		if err != nil {
			return c.JSON(400, err)
		}

		data := controller.FormattedPost{Message: input.Message}
		data.User.ID = user.ID
		data.User.Name = user.Name
		data.User.Avatar = user.Avatar
		data.Coordinate.Latitude = input.Coordinate.Latitude
		data.Coordinate.Longitude = input.Coordinate.Longitude

		err = postController.Save(data, ctx)
		if err != nil {
			return c.JSON(500, err)
		}
		return c.JSON(200, "OK")
	})

	return e
}
