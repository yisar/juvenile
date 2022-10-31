package api

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func DockerV() string {

	cmd := exec.Command("docker", "-v")
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

func Health(w http.ResponseWriter, r *http.Request) {
	output := DockerV()

	GlobalChan.Messages <- output

	w.Write([]byte(output))
}

// b.Messages <- output
