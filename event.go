package filew

import "fmt"

type event struct {
	path, op string
	err      error
}

func (ev event) String() string {
	if ev.err != nil {
		return ev.err.Error()
	}
	return fmt.Sprintf("%s %s", ev.path, ev.op)
}
