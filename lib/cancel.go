package lib

import (
	"os"
	"os/signal"
	"syscall"
)

func MakeCancelChan() chan struct{} {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	return func(sigCh chan os.Signal) chan struct{} {
		cancel := make(chan struct{})
		go func() {
			<-sigCh
			close(cancel)
		}()
		return cancel
	}(sig)
}
