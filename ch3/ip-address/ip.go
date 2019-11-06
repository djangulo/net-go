package main

import (
	"fmt"
	"os"
	"net"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
	}

	name := os.Args[1]

	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	}	else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)

}
