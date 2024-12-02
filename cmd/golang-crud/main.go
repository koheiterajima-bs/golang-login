// HTTPのハンドリング
package main

import (
	"log"
	"net/http"

	"github.com/koheiterajima-bs/golang-crud/internal/article"
)

func main() {
	http.HandleFunc("/", article.Index)
	http.HandleFunc("/show", article.Show)
	http.HandleFunc("/create", article.Create)
	http.HandleFunc("/edit", article.Edit)
	http.HandleFunc("/delete", article.Delete)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
