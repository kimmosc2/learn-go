package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type HttpHandle func(w http.ResponseWriter, r *http.Request) error

func ErrorWrapper(handle HttpHandle) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handle(writer, request)
		if err != nil {
			log.Println("error:", err.Error())
		}
	}
}

func main() {
	http.HandleFunc("/list/", ErrorWrapper(viewFile))
	log.Fatal(http.ListenAndServe(":8989", nil))
}

func viewFile(w http.ResponseWriter, r *http.Request) error {
	filename := r.URL.Path[len("/list/"):]
	open, err := os.Open(filename)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(open)
	if err != nil {
		return err
	}
	w.Write(b)
	return nil
}
