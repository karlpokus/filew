package filew

import (
	"time"

	"github.com/karlpokus/filew/internal/tree"
)

// Watch watches files and returns a channel with file change events.
// note: the channel might also send an error if Walker.Walk fails. This also
// closes the channel.
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
		defer close(events)
		for _ = range ticker.C { // poll
			fs := make(tree.Tree)
			err := w.Walk(root, fs)
			if err != nil {
				events <- event{err: err}
				return
			}
			//go diff(tree.Copy(base), tree.Copy(fs), events)
			diff(base, fs, events)
			base = fs
		}
	}()
	return
}
