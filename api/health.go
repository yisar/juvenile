package api

import (
"net/http"
)


func Health(w http.ResponseWriter, r *http.Request) {
	Exec("docker", "-v")

	// GlobalChan.Messages <- output

	// w.Write([]byte(output))
}

// b.Messages <- output
