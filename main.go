package main

import "fmt"

func main() {
	test := make(chan string)
	server(test)
	select {
	case t, ok := <-test:
		if ok != false {
			fmt.Println(t)
		}
		fmt.Println("ERROR")
	default:
		fmt.Println("error")
	}
}
