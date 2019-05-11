package filew

import (
	"sort"
	"testing"

	"github.com/karlpokus/filew/internal/mockw"
  "github.com/karlpokus/filew/internal/tree"
)

func TestWatch(t *testing.T) {
	mock := &mockw.Mock{}
	events, err := Watch("whatever", mock)
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}
	var got []event
	for ev := range events { // will close after sending the err
		got = append(got, ev)
	}
	sort.Slice(got, func(i, j int) bool {
		return got[i].path < got[j].path
	})
	want := []event{
    {"", "", mockw.PathErr},
		{"a", "updated", nil},
		{"b", "removed", nil},
		{"c", "created", nil},
	}
	if !match(got, want) {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func match(got, want []event) bool {
  if len(got) != len(want) {
    return false
  }
  for i := range want {
    if got[i] != want[i] {
      return false
    }
  }
  return true
}

func BenchmarkWalk(b *testing.B) {
  w := fswalk{}
  base := make(tree.Tree)
  b.ReportAllocs()
  b.ResetTimer()
  for i := 0; i < b.N; i++ {
    w.Walk("testdata/express", base) // ~300 files
  }

}
