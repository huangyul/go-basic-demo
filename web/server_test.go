package web

import (
	"testing"
)

func TestServer(t *testing.T) {
	var h = &HTTPServer{}

	h.Start(":8081")
}
