package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	username   string = "root"
	password   string = "secret"
	ip         string = "localhost"
	port       string = "3307"
	database   string = "blog"
	connection string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, ip, port, database)
)

type Post struct {
	id      int
	content string
	author  string
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
}

func main() {
	// post := Post{content: "Hello World!", author: "bork bork"}
	// fmt.Println(post)
	// err := post.Create()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(post)

	// readPost, _ := GetPost(5)
	// readPost.DeletePost()
	// readPost, _ = GetPost(2)
	// fmt.Println(readPost)
	// readPost, _ = GetPost(3)
	// fmt.Println(readPost)

	// GetPostCount()
	allposts, err := GetAllPosts(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(allposts)
	// DeletePost(5)
	// num := GetPostCount()
	// InsertPost("I am a robot beep boop", "Not a human")
	// DeletePost(num)
	// GetPostCount()

	// DeleteLatestPost()

	defer db.Close()
}

func (post *Post) Create() (err error) {
	stmt := "insert into posts (content, author) values (?, ?)"
	res, err := db.Exec(stmt, post.content, post.author) // .Scan adds post id to the struct
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	post.id = int(lastId)
	return
}

func (post *Post) Update() (err error) {
	stmt := "UPDATE posts SET content = ?, author = ? WHERE id = ?"
	_, err = db.Exec(stmt, post.content, post.author, post.id) // .Scan adds post id to the struct
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (post *Post) Delete() (err error) {
	DeleteById(post.id)
	return
}

func DeleteById(id int) (err error) {
	stmt := "delete from posts where id = ?"
	_, err = db.Exec(stmt, id)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}

	stmt := "select id, content, author from posts where id = ?"
	err = db.QueryRow(stmt, id).Scan(&post.id, &post.content, &post.author) // selects part of db
	if err != nil {
		return
	}
	return
}

func GetPostCount() (count int) {
	err := db.QueryRow("select count(*) from posts").Scan(&count) // selects part of db
	if err != nil {
		log.Fatal(err)
		return 0
	}
	// fmt.Printf("Number of rows are %d\n", count)
	return
}

//  Gathers all posts from database and puts them in an array of posts
//  USAGE: variable, err := GetAllPosts()
func GetAllPosts(limit int) (allposts []Post, err error) { // NO LIMIT BECAUSE SMALL SCALE
	post := Post{}
	if limit == 0 {
		limit = GetPostCount()
	}

	rows, err := db.Query("select id, content, author from posts limit ?", limit) // selects part of db
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // IMPORTANT

	for rows.Next() {
		err := rows.Scan(&post.id, &post.content, &post.author) // scans, if err, err variable will change
		if err != nil {                                         // else rows will be modified by scan
			log.Fatal(err)
		}
		allposts = append(allposts, post)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(allposts[1].content)
	return
}

func (post *Post) DeletePost() (err error) {
	_, err = db.Exec("DELETE FROM posts WHERE id = ?", post.id)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func DeleteLatestPost() {
	all, err := GetAllPosts(0)
	if err != nil {
		log.Fatal(err)
	}

	// DeletePost(all[len(all)-1].id)
	fmt.Println(all[:len(all)-1])
}
