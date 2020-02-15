package main

import (
	Config "EcderGo/app/config"
	"fmt"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("temp/index.html", "temp/header.html", "temp/footer.html")

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", IndexHandler)

	conf := Config.InitConfig()

	port := conf.ListeningPort

	fmt.Printf("Listening on port: %s \n", port)
	fmt.Printf("http://localhost:%s", port)

	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

