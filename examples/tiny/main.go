package main

import (
	"encoding/hex"
	"fmt"

	"github.com/go-compile/tinytime"
)

func main() {
	d := tinytime.Now().Days()

	// store in as small format as possible as HEX
	fmt.Println("Smallest Possible:", tinytime.ToBytes(d))
	fmt.Println("Smallest Possible (Hex):", hex.EncodeToString(tinytime.ToBytes(d)))

	// Sometimes it is appropriate to have a fixed length array when creating wire
	// protocols or binary formats.
	fmt.Println("\nFixed Length (3):", tinytime.ToFixed(tinytime.ToBytes(d), 3))
}
