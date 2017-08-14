package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type LineOfLog struct {
	RemoteAddr  string
	ContentType string
	Path        string
	Query       string
	Method      string
	Body        string
}

// var TemplateOfLog = `Remote address:{{.RemoteAddr}} Content-Type:{{.ContentType}} HTTP method:{{.Method}}`

func log_http(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//		bufbody := new(bytes.Buffer)
		//		bufbody.ReadFrom(r.Body)
		//		body := bufbody.String()
		//
		//		line := LineOfLog{
		//			r.RemoteAddr,
		//			r.Header.Get("Content-Type"),
		//			r.URL.Path,
		//			r.URL.RawQuery,
		//			r.Method, body,
		//		}
		//		tmpl, err := template.New("line").Parse(TemplateOfLog)
		//		if err != nil {
		//			panic(err)
		//		}
		//
		//		bufline := new(bytes.Buffer)
		//		err = tmpl.Execute(bufline, line)
		//		if err != nil {
		//			panic(err)
		//		}

		handler.ServeHTTP(w, r)
	})
}

func sleep_princess(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	fmt.Fprintf(w, "Good morning!")
}

func main() {
	http.HandleFunc("/", sleep_princess)

	if err := http.ListenAndServe(":8080", log_http(http.DefaultServeMux)); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
