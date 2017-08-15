package main

import (
	"io"
	"net/http"
	"runtime"
	"sync"
)

var html = `
<html>
  <head>
    <style>
      html, body, main {
        height: 100%;
        background: #660066;
      }

      main {
        display: flex;
        justify-content: center;
        align-items: center;
      }

      .submit {
        border: solid 1px #ccc;
        padding: 10px 30px;
        margin: 0 0 20px;
        font-size: 1.2in;
        cursor: pointer;
      }
    </style>
  </head>
  <body>
    <main>
      <form action="." method="POST">
        <input type="submit"class="submit" value="力を行使する">
      </form>
    </main>
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

//func validate_ip(ip string) bool {
//	start_ip := net.ParseIP("192.168.15.0")
//	end_ip := net.ParseIP("192.168.15.255")
//	trial := net.ParseIP(ip)
//
//	if trial.To4() == nil {
//		return false
//	}
//	if bytes.Compare(trial, start_ip) >= 0 && bytes.Compare(trial, end_ip) <= 0 {
//		return true
//	}
//
//	return false
//}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		io.WriteString(w, html)
	} else if r.Method == "POST" {
		//r.ParseForm()
		//input_url := r.Form["ip"][0]
		//_, err_parse := url.Parse(input_url)
		//err_vaild := validate_ip(input_url)
		//if err_parse != nil {
		//	io.WriteString(w, "invaild URI!")
		//} else if err_vaild != true {
		//	io.WriteString(w, "invaild IP address")
		//} else {
		go attack_w_goroutine("http://192.168.15.5:80")

		io.WriteString(w, "力はふるわれた…")
		//}
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":80", nil)
}
