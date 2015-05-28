package main
import "fmt"
import "time"

const LOOP = 15

func main() {
	for g:=0; g<LOOP; g++ {
		go vanilla(g)
	}
	for s:=0; s<LOOP * LOOP; s++ {
		log_info(fmt.Sprintf("main sleep(%d)",s))
		time.Sleep(time.Second)
	}
}
func vanilla(a int) {
	log_info(fmt.Sprintf("(id %d)goroutine : start", a))
	for i:=0; i<LOOP; i++ {
		log_debug(fmt.Sprintf("(id %d)goroutine : step %d", a, i))
		for s:=0; s<a; s++ {
			time.Sleep(time.Second)
		}
		if i % 3 == 0 {
			log_warn(fmt.Sprintf("(id %d)goroutine : some trouble occurred : step %d", a, i))
		}
	}
	log_info(fmt.Sprintf("(id %d)goroutine : finished", a))
}
func log_debug(s string) {
	fmt.Println("[DEBUG]"+s)
}
func log_info(s string) {
	fmt.Println("[INFO]"+s)
}
func log_warn(s string) {
	fmt.Println("[WARN]"+s)
}
func log_error(s string) {
	fmt.Println("[ERROR]"+s)
}

