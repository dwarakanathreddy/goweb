package main

import (
	"errors"
	"net/http"

	"github.com/dwarakanathreddy/goweb/pkg/handlers"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	_, err := fmt.Fprintf(w, "hello")
	// 	if err != nil {
	// 		fmt.Println("err")
	// 	}

	// })

	_ = http.ListenAndServe(":8082", nil)
}

func Divide(a int, b int) (int, error) {
	if b <= 0 {
		return 0, errors.New("error")
	}
	return a / b, nil
}
