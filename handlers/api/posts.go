package api

import (
	"net/http"
	"strconv"

	"github.com/bhavyaunacademy/goCRUD/handlers/utils"
	"github.com/labstack/echo/v4"
)

func GetAllPosts(c echo.Context) error {
	users := utils.ReadAllPosts()
	return c.JSON(http.StatusCreated, users)
}

func GetPostByUID(c echo.Context) error {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return c.JSON(403, "Invalid uid")
	}
	posts := utils.ReadAllPosts()
	uid_set := map[int]bool{}
	for k := range posts.Posts {
		uid_set[posts.Posts[k].UID] = true
	}
	if !uid_set[uid] {
		return c.JSON(404, "Post Not Found")
	}
	return c.JSON(http.StatusCreated, posts.Posts[uid-1])

}

func GetPostByUserUID(c echo.Context) error {
	user_uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return c.JSON(403, "Invalid uid")
	}
	posts := utils.ReadAllPosts()
	var userPosts []utils.Post
	for i := 0; i < len(posts.Posts); i++ {
		if posts.Posts[i].UserUID == user_uid {
			userPosts = append(userPosts, posts.Posts[i])
		}
	}

	if len(userPosts) == 0 {
		return c.JSON(404, "Not Found")
	}

	return c.JSON(http.StatusCreated, userPosts)
}

func CreatePost(c echo.Context) (err error) {

	post := new(utils.Post)
	if err = c.Bind(post); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user_uid := post.UserUID
	users := utils.ReadAllUsers()
	if user_uid > len(users.Users) || user_uid < 1 {
		return c.JSON(http.StatusNotFound, "User Not Found")
	}

	utils.WritePost(*post)

	return c.JSON(http.StatusCreated, post)
}

func UpdatePost(c echo.Context) (err error) {

	post := new(utils.Post)
	if err = c.Bind(post); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	post_uid, err := strconv.Atoi(c.Param("postUID"))
	if err != nil {
		return c.JSON(403, "Invalid post uid")
	}

	post.UID = post_uid
	user_uid := post.UserUID
	users := utils.ReadAllUsers()
	if user_uid > len(users.Users) || user_uid < 1 {
		return c.JSON(404, "User Not Found")
	}

	posts := utils.ReadAllPosts()
	uid_set := map[int]bool{}
	for k := range posts.Posts {
		uid_set[posts.Posts[k].UID] = true
	}
	if !uid_set[post_uid] {
		return c.JSON(404, "Post Not Found")
	}

	for i := 0; i < len(posts.Posts); i++ {
		if posts.Posts[i].UID == post.UID && posts.Posts[i].UserUID == post.UserUID {
			utils.WritePost(*post)
			return c.JSON(http.StatusCreated, post)
		}
	}

	return c.JSON(http.StatusNotFound, "Post Does not belong to User")

}

func DeletePost(c echo.Context) (err error) {

	post := new(utils.Post)
	if err = c.Bind(post); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	post_uid, err := strconv.Atoi(c.Param("postUID"))
	if err != nil {
		return c.JSON(403, "Invalid post uid")
	}

	post.UID = post_uid

	user_uid := post.UserUID
	users := utils.ReadAllUsers()
	if user_uid > len(users.Users) || user_uid < 1 {
		return c.JSON(404, "User Not Found")
	}

	posts := utils.ReadAllPosts()
	uid_set := map[int]bool{}
	for k := range posts.Posts {
		uid_set[posts.Posts[k].UID] = true
	}
	if !uid_set[post_uid] {
		return c.JSON(404, "Post Not Found")
	}

	for i := 0; i < len(posts.Posts); i++ {
		if posts.Posts[i].UID == post.UID && posts.Posts[i].UserUID == post.UserUID {
			utils.RemovePost(*post)
			return c.JSON(http.StatusCreated, post)
		}
	}

	return c.JSON(404, "Post Does not belong to User")
}
