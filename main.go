package main

import "fmt"

func main() {
	var x Alert = Alert{FATAL, BAD_CERTIFICATE }

	fmt.Print(x)
}