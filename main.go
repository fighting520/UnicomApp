// UnicomApp project main.go
package main

import (
	"UnicomApp/db"

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
		g := db.Goods{}
		if len(r.Form["amount"]) > 0 {
			g.Amount = r.Form["amount"][0]
		}
		if len(r.Form["cpTradeId"]) > 0 {
			g.CpTradeId = r.Form["cpTradeId"][0]
		}

		if len(r.Form["money"]) > 0 {
			g.Money = r.Form["money"][0]
		}
		if len(r.Form["payDetailId"]) > 0 {
			g.PayDetailId = r.Form["payDetailId"][0]
		}
		if len(r.Form["payStatus"]) > 0 {
			g.PayStatus = r.Form["payStatus"][0]
		}
		if len(r.Form["payType"]) > 0 {
			g.PayType = r.Form["payType"][0]
		}
		if len(r.Form["price"]) > 0 {
			g.Price = r.Form["price"][0]
		}
		if len(r.Form["productName"]) > 0 {
			g.ProductName = r.Form["productName"][0]
		}
		if len(r.Form["productType"]) > 0 {
			g.ProductType = r.Form["productType"][0]
		}
		if len(r.Form["sign"]) > 0 {
			g.Sign = r.Form["sign"][0]
		}
		err := db.Save(g)
		if err != nil {
			log.Println(err)
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
