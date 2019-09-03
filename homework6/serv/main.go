package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	res.Header().Set("Content-Type", "text/html")
	str := fmt.Sprintf(`<doctype html>
	<html>
		<head>
			<title>Hello World!</title>
		</head>
		<body>
			Hello %s!
		</body>
	</html>`, name)
	io.WriteString(res, str)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":80", nil)
}
