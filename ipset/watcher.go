package ipset

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/radovskyb/watcher"
)

func StartWatcher() {
	w := watcher.New()
	w.SetMaxEvents(1)

	w.FilterOps(watcher.Write, watcher.Create, watcher.Remove)

	r := regexp.MustCompile(".(ip|net)set$")
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case event := <-w.Event:
				fmt.Println(event) // Print the event's info.
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.AddRecursive("/tmp/update-iplist"); err != nil {
		log.Fatalln(err)
	}

	for path, f := range w.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}

	fmt.Println()

	go func() {
		w.Wait()
		w.TriggerEvent(watcher.Write, nil)
		w.TriggerEvent(watcher.Create, nil)
		w.TriggerEvent(watcher.Remove, nil)
	}()

	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
