package main

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"sync"
)

var html = `
<html>
  <body>
    <h1>Bench marker</h1>
    <form action="." method="POST">
      Your IP Address: <input type="text" name="ip" placeholder="0.0.0.0">
      <input type="submit" value="start!">
    </form>
  </body>
</html>
`

func attack_w_goroutine(url string) {
	count := 0
	var wg sync.WaitGroup

	for count < 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req, _ := http.NewRequest("GET", url, nil)

			client := new(http.Client)
			resp, _ := client.Do(req)
			defer resp.Body.Close()

		}()
		count++
	}
	wg.Wait()
}

func validate_ip(ip string) bool {
	start_ip := net.ParseIP("192.168.15.0")
	end_ip := net.ParseIP("192.168.15.255")
	trial := net.ParseIP(ip)

	if trial.To4() == nil {
		return false
	}
	if bytes.Compare(trial, start_ip) >= 0 && bytes.Compare(trial, end_ip) <= 0 {
		return true
	}

	return false
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//t, _ := template.ParseFiles("views/index.html")
		//t.Execute(w, nil)
		io.WriteString(w, html)
	} else if r.Method == "POST" {
		r.ParseForm()
		input_url := r.Form["ip"][0]
		_, err := url.Parse(input_url)
		if err != nil {
			io.WriteString(w, "invaild URI!")
		} else {
			go attack_w_goroutine("http://" + r.Form["ip"][0] + ":80")

			io.WriteString(w, "OK!")
		}
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":80", nil)
}
