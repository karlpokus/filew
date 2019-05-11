package mockw

import (
	"errors"

	"github.com/karlpokus/filew/internal/tree"
)

var PathErr = errors.New("oops")

type Mock struct {
	i int
}

func (m *Mock) Walk(root string, t tree.Tree) error {
	defer func() { m.i++ }()
	if m.i == 0 { // base
		t["a"] = int64(1)
		t["b"] = int64(1)
		return nil
	}
	if m.i == 1 { // fs
		t["a"] = int64(2)
		t["c"] = int64(1)
		return nil
	}
	return PathErr
}
