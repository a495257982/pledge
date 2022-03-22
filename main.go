package main

import (
	"log"
	"os/exec"
)

func main() {

	cmnd := exec.Command("main.sh", "2")
	//cmnd.Run() // and wait
	cmnd.Start()
	log.Println("log")
}
