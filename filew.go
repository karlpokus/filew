package filew

import (
	"time"

	"github.com/karlpokus/filew/internal/tree"
)

// Watch watches files and returns a channels with file change events
func Watch(root string, w Walker) (events chan event, err error) {
	if w == nil {
		w = fswalk{}
	}
	base := make(tree.Tree)
	err = w.Walk(root, base)
	if err != nil {
		return
	}
	events = make(chan event)
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		defer close(events)      // TODO: alert future senders about this
		for _ = range ticker.C { // poll
			fs := make(tree.Tree)
			err := w.Walk(root, fs)
			if err != nil {
				events <- event{err: err}
				return
			}
			diff(base, fs, events) // TODO: concurrent op
			base = fs
		}
	}()
	return
}
