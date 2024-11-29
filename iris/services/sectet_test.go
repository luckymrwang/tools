package services

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	secret, err := new(SecretService).Add("")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(secret)
}
