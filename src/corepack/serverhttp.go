package corepack

import (
	"fmt"
	"io"
	"net/http"
)

func doServerHttp() {
	fmt.Println("serverHttp start")
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9998", nil)
	http.Handle("/assets/", http.StripPrefix(
		"/assets/",
		http.FileServer(http.Dir("assets")),
	),
	)
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	html := "<!doctype html>"
	html += "<html>"
	html += "	<head>"
	html += "		<title>Say Hello</title>"
	html += "	</head>"
	html += "	<body>"
	html += "		<h1>Hello World!</h1>"
	html += "	</body>"
	html += "</html>"
	io.WriteString(res, html)
}
