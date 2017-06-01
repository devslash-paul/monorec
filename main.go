package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func waitFor(proc *exec.Cmd, ch chan *exec.Cmd) {
	err := proc.Wait()
	if err != nil {
		log.Fatal(err)
	}

	ch <- proc
}

func main() {
	subProcess := exec.Command("cmd")

	stdin, err := subProcess.StdinPipe()

	if err != nil {
		fmt.Println(err)
	}

	defer stdin.Close()

	subProcess.Stdout = os.Stdout
	subProcess.Stderr = os.Stderr

	reader := bufio.NewReader(os.Stdin)
	subProcess.Start()
	ch := make(chan *exec.Cmd)
	go waitFor(subProcess, ch)

	for {
		text, _ := reader.ReadString('\n')
		io.WriteString(stdin, text)
		v := subProcess.ProcessState

		if v != nil && v.Exited() {
			return
		}
	}
}
