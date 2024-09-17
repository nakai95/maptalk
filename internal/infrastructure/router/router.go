package router

import (
	"encoding/json"
	"fmt"
	"log"
	"maptalk/internal/infrastructure/datastore"
	"maptalk/internal/interface/controller"
	"maptalk/internal/interface/presenter"
	"maptalk/internal/interface/repository"
	"time"

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

	e.GET("/sse", func(c echo.Context) error {
		log.Printf("SSE client connected, ip: %v", c.RealIP())

		w := c.Response()
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		ch, err := postRepository.Listener(c.Request().Context())
		if err != nil {
			return c.JSON(500, err)
		}

		for {
			select {
			case <-c.Request().Context().Done():
				log.Printf("SSE client disconnected, ip: %v", c.RealIP())
				return nil

			case post := <-ch:
				if post.ID == "" {
					continue
				}
				output, err := postPresenter.PresentPost(post)
				if err != nil {
					return c.JSON(500, err)
				}

				bytes, _ := json.Marshal(output)

				fmt.Fprintf(w, "data: %s\n\n", bytes)
				w.Flush()
				// case <-ticker.C:
				// 	post := <-ch

				// 	if post.ID == "" {
				// 		continue
				// 	}

				// 	output, err := postPresenter.PresentPost(post)
				// 	if err != nil {
				// 		return c.JSON(500, err)
				// 	}

				// 	fmt.Fprintf(w, "data: %#v\n\n", output)
				// 	w.Flush()

			}
		}
	})

	return e
}
