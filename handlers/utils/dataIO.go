package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type User struct {
	UID   int    `json:"uid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type Users struct {
	Users []User `json:"users"`
}

type Vote struct {
	UpVotes   int `json:"upvotes"`
	DownVotes int `json:"downvotes"`
}

type Comment struct {
	UID     int    `json:"uid"`
	UserUID int    `json:"user_uid"`
	Content string `json:"content"`
}

type Post struct {
	UID      int       `json:"uid"`
	UserUID  int       `json:"user_uid"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Votes    Vote      `json:"votes"`
	Comments []Comment `json:"comments"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}

func ReadAllUsers() Users {

	content, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload Users

	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return payload
}

func ReadAllPosts() Posts {

	content, err := ioutil.ReadFile("posts.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload Posts

	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return payload
}

func WritePost(post Post) {
	posts := ReadAllPosts()

	if post.UID == 0 {
		post.UID = posts.Posts[len(posts.Posts)-1].UID + 1
		posts.Posts = append(posts.Posts, post)
	} else {
		for i := 0; i < len(posts.Posts); i++ {
			if post.UID == posts.Posts[i].UID {
				posts.Posts[i].Title = post.Title
				posts.Posts[i].Content = post.Content
				posts.Posts[i].Comments = post.Comments
			}
		}
	}

	file, err := json.MarshalIndent(posts, "", "")
	if err != nil {
		log.Fatal("Error during Marshal(): ", err)
	}
	ioutil.WriteFile("posts.json", file, 0644)
}

func RemovePost(post Post) {
	posts := ReadAllPosts()

	for i := 0; i < len(posts.Posts); i++ {
		if post.UID == posts.Posts[i].UID {
			posts.Posts = append(posts.Posts[:i], posts.Posts[i+1:]...)
			break
		}
	}

	file, err := json.MarshalIndent(posts, "", "")
	if err != nil {
		log.Fatal("Error during Marshal(): ", err)
	}
	ioutil.WriteFile("posts.json", file, 0644)
}

func WriteComment(comment Comment, post Post) {

	if comment.UID == 0 {
		if post.Comments == nil {
			post.Comments = make([]Comment, 0)
			comment.UID = 1
		} else {
			comment.UID = post.Comments[len(post.Comments)-1].UID + 1
		}

		post.Comments = append(post.Comments, comment)
	} else {
		for i := 0; i < len(post.Comments); i++ {
			if comment.UID == post.Comments[i].UID {
				post.Comments[i] = comment
			}
		}
	}

	WritePost(post)
}

func RemoveComment(comment Comment, post Post) {

	for i := 0; i < len(post.Comments); i++ {
		if comment.UID == post.Comments[i].UID {
			post.Comments = append(post.Comments[:i], post.Comments[i+1:]...)
			break
		}
	}

	WritePost(post)
}
