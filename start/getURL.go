package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	for _, url := range os.Args[1:] {
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 			os.Exit(1)
// 		}

// 		// 确保 Body 最终一定会被关闭
// 		defer resp.Body.Close()

// 		// 检查 HTTP 状态码是否成功
// 		if resp.StatusCode != http.StatusOK {
// 			fmt.Fprintf(os.Stderr, "fetch: %s returned status %d\n", url, resp.StatusCode)
// 			os.Exit(1)
// 		}

// 		// b, err := ioutil.ReadAll(resp.Body)
// 		// 替换掉已弃用的 ioutil.ReadAll
// 		b, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
// 			os.Exit(1)
// 		}

// 		fmt.Printf("%s", b)
// 	}
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
// 	for k, v := range r.Header {
// 		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
// 	}
// 	fmt.Fprintf(w, "Host = %q\n", r.Host)
// 	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
// 	if err := r.ParseForm(); err != nil {
// 		log.Print(err)
// 	}
// 	for k, v := range r.Form {
// 		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
// 	}
// }

// // func handler(w http.ResponseWriter, r *http.Request) {
// // 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// // }

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Fatal(http.ListenAndServe("localhost:8000", nil))
// }

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
