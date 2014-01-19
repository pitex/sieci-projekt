package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
	"strings"
	// "net"
	"./joinservice"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter your IP address: ")
	myip, _ := reader.ReadString('\n')
	myip = myip[:len(myip)-1]

	

	fmt.Printf("Enter connection limit for this computer: ")
	temp, _ := reader.ReadString('\n')
	temp = temp[:len(temp)-1]

	capacity, err := strconv.ParseInt(temp, 10, 0)

	if err != nil {
		log.Fatal(err)
	}

	if capacity < 2 {
		capacity = 2
	}



	fmt.Printf("Is this the first computer in network?: ")
	ans, _ := reader.ReadString('\n')
	ans = strings.ToLower(ans[:len(ans)-1])

	root := ans == "y" || ans == "yes"

	var ip string

	if !root {
		fmt.Printf("Enter IP address of a computer in network: ")
		ip, _ = reader.ReadString('\n')
		ip = ip[:len(ip)-1]
	}

	cli := joinservice.NewClient(myip,int(capacity),ip,root)

	cli.Connect()
}