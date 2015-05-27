package main
import "fmt"
import "time"

const LOOP = 15

func main() {
	for g:=0; g<LOOP; g++ {
		go vanilla(g)
	}
	for s:=0; s<LOOP * LOOP; s++ {
		fmt.Println(fmt.Sprintf("main sleep(%d)",s))
		time.Sleep(time.Second)
	}
}
func vanilla(a int) {
	for i:=0; i<LOOP; i++ {
		fmt.Println(fmt.Sprintf("(%d)goroutine : step %d", a, i))
		for s:=0; s<a; s++ {
			time.Sleep(time.Second)
		}
	}
	fmt.Println(fmt.Sprintf("(%d)goroutine : finished", a))
}
