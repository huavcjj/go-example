package main

import "os"

func main() {

	// panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

// panic: a problem

// goroutine 1 [running]:
// main.main()
// 	/.../learn-go/42-Panic/main.go:7 +0x2c
// exit status 2
