package main

import (
	"log"
	"os/exec"
)

func main() {

	cmnd := exec.Command("test.sh", "2")
	//cmnd.Run() // and wait
	cmnd.Start()
	log.Println("log")
}
