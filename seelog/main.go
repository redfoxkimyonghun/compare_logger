package main
import "fmt"
import "time"
import log "github.com/cihub/seelog"

const LOOP = 15

func configLogger() {
    config := `
<seelog type="sync">
    <outputs>
        <filter levels="trace,debug,info">
            <console formatid="ltsv"/>
        </filter>
        <filter levels="warn,error,critical">
            <console formatid="ltsv_error"/>
        </filter>
        <file formatid="ltsv" path="result.log"/>
    </outputs>
    <formats>
        <format id="ltsv" format="time:%Date(2006-01-02T15:04:05.000Z07:00)%tlev:%l%tmsg:%Msg%n"/>
        <format id="ltsv_error"
            format="%EscM(31)time:%Date(2006-01-02T15:04:05.000Z07:00)%tlev:%l%tmsg:%Msg%EscM(0)%n"/>
    </formats>
</seelog>`

    logger, err := log.LoggerFromConfigAsBytes([]byte(config))
    if err != nil {
        panic(err)
    }
    log.ReplaceLogger(logger)
}

func main() {
	configLogger()
	for g:=0; g<LOOP; g++ {
		go vanilla(g)
	}
	for s:=0; s<LOOP * LOOP; s++ {
		log.Info(fmt.Sprintf("main sleep(%d)",s))
		time.Sleep(time.Second)
	}
}
func vanilla(a int) {
	for i:=0; i<LOOP; i++ {
		log.Info(fmt.Sprintf("(%d)goroutine : step %d", a, i))
		for s:=0; s<a; s++ {
			time.Sleep(time.Second)
		}
	}
	log.Info(fmt.Sprintf("(%d)goroutine : finished", a))
}
