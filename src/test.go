package main

import "./protocols/sip"
import "fmt"

func main() {
	msg := new(sip.Message)
	fmt.Println(msg.ToString());
}