package main

import (
	"fmt"
	"synckv"
)

const PORT = 3792

func main() {
	synckv.StartServer(PORT)

	synckv.ClientPut("UofT", "20", PORT)
	fmt.Println(synckv.ClientGet("UofT", PORT))
}
