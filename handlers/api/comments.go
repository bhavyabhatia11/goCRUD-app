package api

import (
	"net/http"
	"strconv"

	"github.com/bhavyaunacademy/goCRUD/handlers/utils"
	"github.com/labstack/echo/v4"
)

func CreateComment(c echo.Context) (err error) {
	comment := new(utils.Comment)
	if err = c.Bind(comment); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user_uid := comment.UserUID
	users := utils.ReadAllUsers()
	if user_uid > len(users.Users) || user_uid < 1 {
		return c.JSON(http.StatusNotFound, "User Not Found")
	}

	post_uid, err := strconv.Atoi(c.Param("postUID"))
	if err != nil {
		return c.JSON(403, "Invalid post uid")
	}

	posts := utils.ReadAllPosts()
	for i := 0; i < len(posts.Posts); i++ {
		if posts.Posts[i].UID == post_uid && posts.Posts[i].UserUID == user_uid {
			utils.WriteComment(*comment, posts.Posts[i])
			return c.JSON(http.StatusCreated, comment)
		}
	}

	return c.JSON(404, "Comment Not Found")
}

func UpdateComment(c echo.Context) (err error) {
	comment := new(utils.Comment)
	if err = c.Bind(comment); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if comment.UID == 0 {
		return c.JSON(403, "Please Give comment uid")
	}
	user_uid := comment.UserUID
	users := utils.ReadAllUsers()
	if user_uid > len(users.Users) || user_uid < 1 {
		return c.JSON(http.StatusNotFound, "User Not Found")
	}

	post_uid, err := strconv.Atoi(c.Param("postUID"))
	if err != nil {
		return c.JSON(403, "Invalid post uid")
	}

	posts := utils.ReadAllPosts()
	for i := 0; i < len(posts.Posts); i++ {
		if posts.Posts[i].UID == post_uid && posts.Posts[i].Comments != nil {
			for j := 0; j < len(posts.Posts[i].Comments); j++ {
				if posts.Posts[i].Comments[j].UserUID == user_uid && posts.Posts[i].Comments[j].UID == comment.UID {
					utils.WriteComment(*comment, posts.Posts[i])
					return c.JSON(http.StatusCreated, comment)
				}
			}
		}
	}
	return c.JSON(404, "Comment Not Found")
}

func DeleteComment(c echo.Context) (err error) {
	comment := new(utils.Comment)
	if err = c.Bind(comment); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if comment.UID == 0 {
		return c.JSON(403, "Please Give comment uid")
	}

	user_uid := comment.UserUID
	users := utils.ReadAllUsers()
	if user_uid > len(users.Users) || user_uid < 1 {
		return c.JSON(http.StatusNotFound, "User Not Found")
	}

	post_uid, err := strconv.Atoi(c.Param("postUID"))
	if err != nil {
		return c.JSON(403, "Invalid post uid")
	}

	posts := utils.ReadAllPosts()
	for i := 0; i < len(posts.Posts); i++ {
		if posts.Posts[i].UID == post_uid && posts.Posts[i].Comments != nil {
			for j := 0; j < len(posts.Posts[i].Comments); j++ {
				if posts.Posts[i].Comments[j].UserUID == user_uid && posts.Posts[i].Comments[j].UID == comment.UID {
					utils.RemoveComment(*comment, posts.Posts[i])
					return c.JSON(http.StatusCreated, comment)
				}
			}
		}
	}
	return c.JSON(404, "Comment Not Found")
}
