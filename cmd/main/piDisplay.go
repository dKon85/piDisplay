package main

import (
	"fmt"
)

const ADDR string = "192.168.178.41:4223"
const UID string = "7Zw"
const version string = "0.1"

func main() {
	fmt.Println("Starting main V%s", version)
	InitDisplay()
	StartRESTServer(":8089")

	fmt.Print("Press enter to exit.")
	fmt.Scanln()

	DeactivateDisplay()
	DisconnectDisplayManager()
}


