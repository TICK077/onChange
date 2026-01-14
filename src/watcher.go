package main

import (
	"time"

	"github.com/fsnotify/fsnotify"
)

func Watch(cfg *Config, trigger chan struct{}) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	if err := watcher.Add(cfg.WatchDir); err != nil {
		return err
	}

	go func() {
		defer watcher.Close()

		var lastEvent time.Time
		debounce := 200 * time.Millisecond

		for {
			select {
			case <-watcher.Events:
				now := time.Now()

				// Ignore events fired too close together
				if now.Sub(lastEvent) < debounce {
					continue
				}
				lastEvent = now

				select {
				case trigger <- struct{}{}:
				default:
				}

			case err := <-watcher.Errors:
				Error(err.Error())
			}
		}
	}()

	Info("Watching directory: " + cfg.WatchDir)
	select {}
}
