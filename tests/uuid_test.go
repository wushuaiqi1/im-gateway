package tests

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	fmt.Println(uuid.NewString())
}
