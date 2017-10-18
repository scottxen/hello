package main

import (
	"net/http"
	"text/template"
)

func main() {
	// http.Handle("/", &MyHandler{})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		templates := template.New("templateName")
		templates.New("test").Parse(doc)
		templates.New("header").Parse(header)
		templates.New("footer").Parse(footer)
		content := Content{
			[3]string{"Lemon", "Orange", "Apple"},
			"Fruits",
		}
		templates.Lookup("test").Execute(w, content)
	})
	http.ListenAndServe(":8000", nil)
}

const doc = `
{{template "header" .Title}}
	<body>
		<h1>List of {{.Title}}</h1>
		<ul>
			{{range .Fruit}}
				<li><h2>{{.}}</h2></li>
			{{end}}
		</ul>	
	</body>
{{template "footer"}}
`

const header = `
<!DOCTYPE html>
	<html>
		<head><title>{{.}}</title></head>
`

const footer = `
	{{printf "%q" (print "This " "is " "the " "end")}}
</html>
`

type Content struct {
	Fruit [3]string
	Title string
}

// // MyHandler is an object implementing http.Handler
// type MyHandler struct {
// 	http.Handler
// }

// // ServeHTTP implements the http.Handler interface
// func (h *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	path := "public/" + req.URL.Path
// 	f, err := os.Open(path)

// 	if err == nil {
// 		bufferedReader := bufio.NewReader(f)
// 		var contentType string
// 		if strings.HasSuffix(path, ".css") {
// 			contentType = "text/css"
// 		} else if strings.HasSuffix(path, ".html") {
// 			contentType = "text/html"
// 		} else if strings.HasSuffix(path, ".js") {
// 			contentType = "application/javascript"
// 		} else if strings.HasSuffix(path, ".png") {
// 			contentType = "image/png"
// 		} else if strings.HasSuffix(path, ".mp4") {
// 			contentType = "video/mp4"
// 		} else {
// 			contentType = "text/plain"
// 		}
// 		w.Header().Add("Content-Type", contentType)
// 		bufferedReader.WriteTo(w)
// 	} else {
// 		w.WriteHeader(404)
// 		w.Write([]byte("404 - " + http.StatusText(404)))
// 	}
// }
