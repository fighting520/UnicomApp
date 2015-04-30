// UnicomApp project main.go
package main

import (
	"UnicomApp/db"
	"fmt"

	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var realPath *string

func main() {

	http.HandleFunc("/interface/payment", paymentHandler)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println(r.Form)
	if r.Method == "GET" {

	} else if r.Method == "POST" {
		result, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", result)
		var s db.Goods
		json.Unmarshal([]byte(result), &s)
		fmt.Println(s)
		err := db.Save(s)
		if err != nil {
			log.Fatalln(err)
		}

	}
	w.Write([]byte("ok"))

}

func StaticResource(w http.ResponseWriter, r *http.Request) {
	log.Println(realPath)
	path := r.URL.Path
	request_type := path[strings.LastIndex(path, "."):]
	switch request_type {
	case ".css":
		w.Header().Set("content-type", "text/css")
	case ".js":
		w.Header().Set("content-type", "text/javascript")
	default:
	}
	fin, err := os.Open(*realPath + path)
	defer fin.Close()
	if err != nil {
		log.Fatal("static resource:", err)
	}
	fd, _ := ioutil.ReadAll(fin)
	w.Write(fd)

}
