package main

import (
"fmt"
"net/http"
"io/ioutil"
"bytes"
)

func main() {
	resp, err := http.Get("http://www.google.com/")
	if err != nil{
		fmt.Println("error w/ get request")
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("error reading body")
	}
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(b))
	fmt.Println("%q",rdr1)
	fmt.Println("Hello, world.")
}
