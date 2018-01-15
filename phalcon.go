package phalcon

import (
	"net/http"
	"io"
)

func Hello(res http.ResponseWriter, req *http.Request) {

	res.Header().Set(
		"Content-Type", "text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
		<html>
		<head>
		<title>Hello World</title>
		</head>
		<body>
		Hello World!
		</body>
		</html>`,
	)
}

func Listen(port int) {
	//http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}