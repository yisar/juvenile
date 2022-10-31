package api

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Exec(a string, b string) string {

	cmd := exec.Command(a, b)
	cmd.Stdin = os.Stdin

	var output string

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	readout := bufio.NewReader(stdout)
	go func() {
		output += GetOutput(readout)
	}()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	readerr := bufio.NewReader(stderr)
	go func() {
		output += GetOutput(readerr)
	}()

	cmd.Run()

	return output
}

func GetOutput(reader *bufio.Reader) string {
	var sumOutput string
	outputBytes := make([]byte, 200)
	for {
		n, err := reader.Read(outputBytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			sumOutput += err.Error()
		}
		output := string(outputBytes[:n])
		fmt.Print(output) //输出屏幕内容
		sumOutput += output

	}
	return sumOutput
}

// b.Messages <- output
