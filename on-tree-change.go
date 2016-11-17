package ontreechange

import (
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
)

func OnTreeChange(target string, filter func(string, os.FileInfo) bool, callback func(string)) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op == fsnotify.Write {
					callback(event.Name)
				} else if event.Op == fsnotify.Create {
					fileInfo, _ := os.Stat(event.Name)
					if filter(event.Name, fileInfo) {
						watcher.Add(event.Name)
					}
				} else if event.Op == fsnotify.Remove {
					watcher.Remove(event.Name)
				} else if event.Op == fsnotify.Rename {
					fileInfo, _ := os.Stat(event.Name)
					if filter(event.Name, fileInfo) {
						watcher.Add(event.Name)
					}
				}

			case err := <-watcher.Errors:
				if err != nil {
					panic(err)
				}
			}
		}
	}()

	if err := watcher.Add(target); err != nil {
		return err
	}

	filepath.Walk(target, func(name string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err)
		}

		if filter(name, info) {
			watcher.Add(name)
		}

		return nil
	})

	<-done
	return nil
}
