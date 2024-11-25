package tests

import (
	"fmt"
	"im/broadcast"
	"im/higo"
	"net/http"
	"testing"
)

func TestAllUserAPi(t *testing.T) {
	fmt.Println("run broadcast project demo~~~")
	http.HandleFunc("/broadcast", broadcast.AllUserApi)
	http.HandleFunc("/higo", higo.Hello)
	err := http.ListenAndServe(":18081", nil)
	if err != nil {
		return
	}
}
