package main

import (
	"net/http"

	"github.com/bhavyaunacademy/goCRUD/handlers/api"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// routes for getting users
	e.GET("/users", api.GetAllUsers)
	e.GET("/users/:uid", api.GetUserByUID)

	// routes for getting posts
	e.GET("/posts", api.GetAllPosts)
	e.GET("/posts/:uid", api.GetPostByUID)
	e.GET("/posts/user/:uid", api.GetPostByUserUID)

	// CRUD routes for Post
	e.POST("/posts/create-post", api.CreatePost)
	e.PUT("/posts/:postUID/update-post", api.UpdatePost)
	e.DELETE("/posts/:postUID/delete-post", api.DeletePost)

	// CRUD routes for Comments
	e.POST("/posts/:postUID/create-comment", api.CreateComment)
	e.PUT("/posts/:postUID/update-comment", api.UpdateComment)
	e.DELETE("/posts/:postUID/delete-comment", api.DeleteComment)

	e.Logger.Fatal(e.Start(":1323"))
}
