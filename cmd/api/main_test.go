package main

import (
	"os"
	"testing"
	"time"
)

func TestMain_Startup(t *testing.T) {
	os.Setenv("GEMINI_API_KEY", "dummy-test-key")
	os.Setenv("PORT", "8082")
	os.Setenv("ENV", "test")
	os.Setenv("LOG_LEVEL", "debug")

	done := make(chan struct{})
	go func() {
		defer close(done)
		main()
	}()

	time.Sleep(500 * time.Millisecond)

	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf("failed to find process: %v", err)
	}
	if err := p.Signal(os.Interrupt); err != nil {
		t.Fatalf("failed to send interrupt signal: %v", err)
	}

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("main() did not exit after SIGTERM")
	}
}
