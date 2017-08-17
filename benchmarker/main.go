package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/sugyan/termburn"
	"golang.org/x/sync/errgroup"
)

func make_a_magic(url string) error {
	var g errgroup.Group
	c := make(chan bool, 700)

	for count := 0; count < 1000; count++ {
		c <- true
		g.Go(func() error {
			defer func() { <-c }()
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return errors.New("案山子が燃えてしまった…")
			}
			req.Header.Set("User-Agent", "Super Ultra Hyper Miracle Eccentric Wonder Mighty Ultimate Magic")

			client := new(http.Client)
			resp, err := client.Do(req)
			if err != nil {
				return errors.New("案山子が燃えてしまった…")
			}
			defer resp.Body.Close()

			if resp.StatusCode >= 400 {
				return errors.New("案山子が燃えてしまった…")
			}

			return nil
		})
	}

	errG := g.Wait()
	if errG != nil {
		return errG
	}

	return nil
}

func main() {
	fmt.Println("力を見せよう…")
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := make_a_magic("http://192.168.15.5:80")
	if err != nil {
		termburn.Run()
		fmt.Println(err)
		os.Exit(255)
	}

	fmt.Println("魔法は正しく実行された！")
}
