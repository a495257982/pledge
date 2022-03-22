package main

import (
	"fmt"
	"os"
)

func main() {

	shellPath := "/root/go/src/new_rust/lotus-miner"
	argv := make([]string, 2)
	argv[0] = "sealing"
	argv[1] = "jobs"
	attr := new(os.ProcAttr)
	newProcess, err := os.StartProcess(shellPath, argv, attr) //运行脚本
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Process PID", newProcess.Pid)
	processState, err := newProcess.Wait() //等待命令执行完
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("processState PID:", processState.Pid()) //获取PID
	fmt.Println("ProcessExit:", processState.Exited())   //获取进程是否退出
	fmt.Println("ProcessExit:", processState.String())

}
