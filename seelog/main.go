package main
import "fmt"
import "time"
import log "github.com/cihub/seelog"

const LOOP = 15

func configLogger() {
    config := `
<seelog type="sync">
  <outputs formatid="common">
  <console />

  <filter levels="trace,debug">
    <file path="./logs/debug.log" />
  </filter>
  <filter levels="info">
    <file path="./logs/info.log" />
  </filter>
  <filter levels="error,critical">
	<file path="./logs/error.log" />
  </filter>
  </outputs>
  <formats>
    <format id="common" format='{"level":"%LEVEL","time":"%UTCDateT%UTCTime","msg":"%Msg"}%n' />
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
	log.Debug(s)
}
func log_info(s string) {
	log.Info(s)
}
func log_warn(s string) {
	log.Warn(s)
}
func log_error(s string) {
	log.Error(s)
}

