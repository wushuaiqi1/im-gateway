package tests

import (
	"fmt"
	"im/broadcast"
	"net/http"
	"testing"
)

func TestAllUserAPi(t *testing.T) {
	fmt.Println("run broadcast project demo~~~")
	http.HandleFunc("/broadcast", broadcast.AllUserApi)
	err := http.ListenAndServe("localhost:8001", nil)
	if err != nil {
		return
	}
}
