package exception

import(
	"runtime"
	"os"
)

func Try(fun func(), catch func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()
	fun()
}

func PrintStack(all bool) {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, all)
	os.Stdout.Write(buf[:n])
}