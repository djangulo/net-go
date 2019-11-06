package main

import (
	"fmt"
	"encoding/asn1"
)

const i = 13

func main() {
	mdata, err := asn1.Marshal(i)
	if err != nil {panic(err)}

	var n int
	_, err = asn1.Unmarshal(mdata, &n)
	if err != nil {panic(err)}

	fmt.Printf("Before marshal: %v %[1]T\n", i)
	fmt.Printf("After unmarshal: %v %[1]T\n", n)
}
