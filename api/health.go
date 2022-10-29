package api

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"
)

func DockerV(b Broker) http.HandlerFunc {

	cmd := exec.Command("docker", "-v")
	cmd.Stdin = os.Stdin

	var wg sync.WaitGroup
	wg.Add(2)

	var output string

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	readout := bufio.NewReader(stdout)
	go func() {
		defer wg.Done()
		output += GetOutput(readout)
	}()

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	readerr := bufio.NewReader(stderr)
	go func() {
		defer wg.Done()
		output += GetOutput(readerr)
	}()

	cmd.Run()
	wg.Wait()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
		go func() {
			for i := 0; ; i++ {
				b.Messages <- output
				time.Sleep(3e9)
			}
		}()
	})
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
