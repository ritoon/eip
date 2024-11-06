package main

func main() {
	// Start a goroutine
	go func() {
		println("Hello from a goroutine")
	}()
	println("Hello from main")
}
