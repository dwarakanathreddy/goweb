package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
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

func Home(w http.ResponseWriter, r *http.Request) {
	data := *r
	fmt.Println(data.Body)
	fmt.Println(r)
	fmt.Fprintf(w, "home")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is from about page")
}
