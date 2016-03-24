package core

import (
	"os"
	"os/signal"
	"syscall"
)

//初始化停止信号
func InitSignal() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)
	return c
}

//处理信号
func HandleSignal(c chan os.Signal) {
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
		case syscall.SIGHUP:
		default:
			return

		}
	}
}
