package handler

import (
	"github.com/williamzhou8818/gifts-miaoshia/database"
)

var (
	orderCh = make(chan database.Order, 10000)
	stopCh  = make(chan struct{}, 1)
)

func InitChannel() {
	go func() {
		<-stopCh
		close(orderCh)
	}()
}

// 目的是想关闭orderCh，该函数可以反复调用
func CloseChannel() {
	//
	select {
	case stopCh <- struct{}{}: //这了不让函数阻塞在本行代码，外面套一个select
	default:
	}
}
