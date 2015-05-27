package main
import "fmt"
import "time"

const LOOP = 5

func main() {
	for g:=0; g<5; g++ {
		go vanilla(g)
	}
	for s:=0; s<25; s++ {
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
