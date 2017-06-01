package main

import "os/exec"
import "os"
import "fmt"
import "bufio"
import "io"

func main() {
	subProcess := exec.Command("cmd")

	stdin, err := subProcess.StdinPipe()

	if err != nil {
		fmt.Println(err)
	}

	defer stdin.Close()

	subProcess.Stdout = os.Stdout
	subProcess.Stderr = os.Stderr

	if err = subProcess.Start(); err != nil {
		fmt.Println("An error occurred: ", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		fmt.Println("You said >>> " + text)
		io.WriteString(stdin, text)
		if subProcess.ProcessState.Exited() {
			return
		}
	}
}
