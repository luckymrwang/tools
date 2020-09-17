package services

import "testing"

func TestEcho(t *testing.T) {
	dService := GetDemoService(nil)
	dService.Echo()
}
