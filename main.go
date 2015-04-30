// UnicomApp project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var realPath *string

func main() {

	http.HandleFunc("/interface/payment", paymentHandler)

	http.ListenAndServe(":9999", nil)
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ok")
	r.ParseForm()
	log.Println(r.Form)
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
