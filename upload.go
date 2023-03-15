package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("We're running...")
	http.HandleFunc("/upload", upload)
}

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		_, err := io.WriteString(h, strconv.FormatInt(crutime, 10))
		if err != nil {
			return
		}
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		err = t.Execute(w, token)
		if err != nil {
			return
		}
	} else {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			return
		}
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {

			}
		}(file)
		_, err = fmt.Fprintf(w, "%v", handler.Header)
		if err != nil {
			return
		}
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)
		_, err = io.Copy(f, file)
		if err != nil {
			return
		}
	}
}
