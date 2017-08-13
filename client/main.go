package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"sync"
)

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

			_, _ = ioutil.ReadAll(resp.Body)

			fmt.Println("count :" + strconv.Itoa(count))
			fmt.Println("Goroutine : " + strconv.Itoa(runtime.NumGoroutine()))
		}()
		count++
	}
	wg.Wait()
}

func validate_ip(ip string) bool {
	start_ip := net.ParseIP("192.168.0.0")
	end_ip := net.ParseIP("192.168.255.255")
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
	w.WriteHeader(http.StatusOK)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/index.html")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		input_url := r.Form["ip"][0]
		_, err := url.Parse(input_url)
		fmt.Println(input_url)
		if err != nil {
			io.WriteString(w, "invaild URI!")
		} else {
			go attack_w_goroutine("http://" + r.Form["ip"][0] + ":80")

			io.WriteString(w, "OK!")
		}
	}
}

func main() {
	// url := "http://localhost:8080/"
	runtime.GOMAXPROCS(runtime.NumCPU())

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8001", nil)
}
