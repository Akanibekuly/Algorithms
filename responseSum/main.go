package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	var uri string
	fmt.Scan(&uri)
	var port, a, b string
	fmt.Scan(&port, &a, &b)
	//curl "http://127.0.0.1:7777?a=2&b=4"
	uri += ":" + port + "?a=" + a + "&b=" + b
	//fmt.Println(uri)
	res, err := http.Get(uri)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	de := json.NewDecoder(res.Body)
	var arr []int
	err = de.Decode(&arr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sum(arr))
}

func sum(arr []int) int {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}
