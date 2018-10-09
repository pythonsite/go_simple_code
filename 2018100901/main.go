package main

import (
	"fmt"
	"log"
	"time"

	"github.com/radovskyb/watcher"
)

func main() {
	w := watcher.New()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Create, watcher.Write, watcher.Chmod)
	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Add("/app/test"); err != nil {
		log.Fatalln(err)
	}
	if err := w.AddRecursive("/app/test"); err != nil {
		log.Fatalln(err)
	}

	go func() {
		w.Wait()
		w.TriggerEvent(watcher.Create, nil)
		w.TriggerEvent(watcher.Write, nil)
		w.TriggerEvent(watcher.Chmod, nil)
	}()
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}

}
