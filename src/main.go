package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
	"strings"
	"./joinservice"
)

//	Main function:
//	First it creates Client and asks for address to connect to.
//	After receiving address it creates a server and starts it.
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
	var address string

	if !root {
		fmt.Printf("Enter IP address of a computer in network: ")
		ip, _ = reader.ReadString('\n')
		ip = ip[:len(ip)-1]

		cli := joinservice.NewClient(myip,int(capacity),ip,root)
		address, err = cli.Connect()

		if err != nil {
			log.Fatal(err)
		}
	}

	server := joinservice.NewServer(myip, address, int(capacity), root)
	err = server.Start()

	if err != nil {
		log.Fatal(err)
	}
}
