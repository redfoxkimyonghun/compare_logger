package main
import "fmt"
func main() {
	go func() {
		fmt.Println("another goroutine")
	}()
	fmt.Println("main goroutine")
}
