package main

import (
	Config "EcderGo/app/config"
	"EcderGo/app/models"
	"fmt"
	"html/template"
	"net/http"
)
// Posts
var posts map[string]*models.Post

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("temp/index.html", "temp/header.html", "temp/footer.html")

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}

	err = t.ExecuteTemplate(w, "index", posts)
}

func AddNewPostHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("temp/write.html", "temp/header.html", "temp/footer.html")

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}

	_ = t.ExecuteTemplate(w, "write", nil)
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("temp/write.html", "temp/header.html", "temp/footer.html")

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}

	id := r.FormValue("id")
	post, found := posts[id]

	if !found {
		http.NotFound(w, r)
	} else {
		_ = t.ExecuteTemplate(w, "write", post)
	}

	http.Redirect(w, r, "/", 302)
}

func SaveNewPostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	text := r.FormValue("content")

	var post *models.Post
	if id != "" {
		post = posts[id]
		post.Title = title
		post.Text = text
	} else {
		id := GenerateId()
		post := models.NewPost(id, title, text)
		posts[post.Id] = post
	}

	http.Redirect(w, r, "/", 302)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		http.NotFound(w, r)
	}

	delete(posts, id)

	http.Redirect(w, r, "/", 302)
}

func main() {
	// For correct css load
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	posts = make(map[string]*models.Post, 0)

	// Handler func
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/write", AddNewPostHandler)
	http.HandleFunc("/edit", EditPostHandler)
	http.HandleFunc("/delete", DeletePostHandler)
	http.HandleFunc("/SavePost", SaveNewPostHandler)

	conf := Config.InitConfig()
	port := conf.ListeningPort

	fmt.Printf("Listening on port: %s \n", port)
	fmt.Printf("http://localhost:%s \n", port)

	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}