package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"sync"
)

func main() {
	var (
		wg = sync.WaitGroup{}
		ch = make(chan struct{})
	)
	wg.Add(3000)
	for i := 0; i < 3000; i++ {
		go func() {
			<-ch
			g.Log().Print("abcdefghijklmnopqrstuvwxyz1234567890")
			wg.Done()
		}()
	}
	close(ch)
	wg.Wait()
}
