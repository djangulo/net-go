package main

import (
	"fmt"
	"net"
	"os"
	"text/tabwriter"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	dotAddr := os.Args[1]

	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	
	const format = "%v\t%v\n"
	tw := tabwriter.NewWriter(os.Stdout, 2, 8, 0, '\t', 0)
	fmt.Fprintf(tw, format, "Address", addr.String())
fmt.Fprintf(tw, format, "Default mask length", bits)
fmt.Fprintf(tw, format, "Leading ones count", ones )
fmt.Fprintf(tw, format, "Mask (hex)", mask) 
fmt.Fprintf(tw, format, "Network", network)
	tw.Flush()
	os.Exit(0)
}
