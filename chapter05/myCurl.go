package chapter05

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter ./xxxx  <url>")
		os.Exit(-1)
	}
}

func Test0504() {
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}

func Test0505() {
	var b bytes.Buffer

	b.Write([]byte("hello"))

	fmt.Fprintf(&b, "world~")

	io.Copy(os.Stdout, &b)
}
