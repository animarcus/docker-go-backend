package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var (
	username       string = os.Getenv("MYSQL_USER")
	password       string = os.Getenv("MYSQL_PASSWORD")
	database       string = os.Getenv("MYSQL_DATABASE")
	db_port        string = os.Getenv("LOCAL_DB_PORT")
	webserver_port string = os.Getenv("LOCAL_WEBSERVER_PORT")
	db_ip          string = "db"
	webserver_ip   string = "backend"

	connection string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, db_ip, db_port, database)
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}
}

type Post struct {
	Id      int    `json:"id"`
	Created string `json:"created"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	// c := cors.New(cors.Options{
	// 	AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	// })

	fmt.Println(webserver_ip + ":" + webserver_port)

	server := http.Server{
		Addr: webserver_ip + ":" + webserver_port,
	}
	http.HandleFunc("/post/", handleRequest)

	fmt.Println("Webserver listening on port", webserver_port)

	server.ListenAndServe()

	// http.HandleFunc("/post/", handleRequest)
	// http.ListenAndServe(":"+webserver_port, nil)

	// fmt.Println(database)

	// post, err := GetAllPosts(0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(post)

	// post := Post{}
	// post.Id = 4
	// post.Title = "bork"
	// post.Content = "virus"

	// post.update()

	// post, err := retrieve(4)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// post := Post{}
	// post.Title = "IT WORKED"
	// post.Content = "Fucking finally"
	// post.create()
	// // fmt.Println(post)

	// posts, err := GetAllPosts(0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(posts)

	// defer db.Close()
}

func setupCorsResponse(w http.ResponseWriter, req *http.Request) {
	(w).Header().Set("Content-Type", "text/html; charset=utf-8")
	(w).Header().Set("Access-Control-Allow-Origin", "*")

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	setupCorsResponse(w, r)
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, _ := retrieve(id)
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}
	fmt.Println(post)
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Accept")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	setupCorsResponse(w, r)
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post Post
	json.Unmarshal(body, &post)
	err = post.create()
	if err != nil {
		return
	}
	// fmt.Println(GetAllPosts(0))
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Accept")
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Print("GGGOGGGGGG")
	setupCorsResponse(w, r)
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	err = post.update()
	if err != nil {
		return
	}
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Accept")
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	setupCorsResponse(w, r)
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	post.delete()
	if err != nil {
		return
	}
	fmt.Println("Deleted post with info:", post)
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Accept")
	w.WriteHeader(200)
	return
}

// Code for interfacing with database

func retrieve(id int) (post Post, err error) {
	post = Post{}

	stmt := "select id, date_format(created, '%d/%m %H:%i'), title, content from posts where id = ?"
	err = db.QueryRow(stmt, id).Scan(&post.Id, &post.Created, &post.Title, &post.Content) // selects part of db
	if err != nil {
		return
	}
	return
}

func (post *Post) create() (err error) {
	stmt := "insert into posts (title, content) values (?, ?)"
	res, err := db.Exec(stmt, &post.Title, &post.Content) // .Scan adds post id to the struct
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	post.Id = int(lastId)
	post.Created = "a "
	return
}

func (post *Post) update() (err error) {
	stmt := "UPDATE posts SET title = ?, content = ? WHERE id = ?"
	_, err = db.Exec(stmt, post.Title, post.Content, post.Id)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (post *Post) delete() (err error) {
	err = DeleteById(post.Id)
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

	rows, err := db.Query("select id, created, title, content from posts limit ?", limit) // selects part of db
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() // IMPORTANT

	for rows.Next() {
		err := rows.Scan(&post.Id, &post.Created, &post.Title, &post.Content) // scans, if err, err variable will change
		if err != nil {                                                       // else rows will be modified by scan
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
	_, err = db.Exec("DELETE FROM posts WHERE id = ?", post.Id)
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
