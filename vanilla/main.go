package main
import "fmt"
import "time"

const LOOP = 15

func main() {
	for g:=0; g<LOOP; g++ {
		go vanilla(g)
	}
	for s:=0; s<LOOP * LOOP; s++ {
		log_warn(fmt.Sprintf("main sleep(%d)",s))
		time.Sleep(time.Second)
	}
}
func vanilla(a int) {
	for i:=0; i<LOOP; i++ {
		log_info(fmt.Sprintf("(id %d)goroutine : step %d", a, i))
		for s:=0; s<a; s++ {
			time.Sleep(time.Second)
		}
	}
	log_warn(fmt.Sprintf("(id %d)goroutine : finished", a))
}
func log_info(s string) {
	s="[INFO]" + s
	fmt.Println(s)
}
func log_warn(s string) {
	s="[WARN]" + s
	fmt.Println(s)
}
func log_error(s string) {
	s="[ERROR]" + s
	fmt.Println(s)
}

