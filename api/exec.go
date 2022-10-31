package api

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func Exec(a string, b ...string) string {

	cmd := exec.Command(a, b...)
	cmd.Stdin = os.Stdin

	var output string

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(stdout)
	cmd.Start()
	for scanner.Scan() {
		GlobalChan.Messages <- scanner.Text()
	}
	cmd.Wait()

	return output
}
