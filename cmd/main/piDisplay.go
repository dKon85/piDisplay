package main

import (
	"fmt"
	"piDisplay/cmd/restServer"
	"piDisplay/cmd/tinker"
)

const ADDR string = "192.168.178.41:4223"
const UID string = "7Zw"
const version string = "0.1"

func main() {

	fmt.Printf("Starting main V%s", version)
	tinker.InitDisplay(ADDR, UID)

	defer fmt.Println("Disconnected.")
	defer tinker.DisconnectDisplayManager()
	defer tinker.DeactivateDisplay()

	restServer.StartRESTServer(":8089")

	fmt.Print("Press enter to exit.")
	fmt.Scanln()

}


