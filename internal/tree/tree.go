package tree

type Tree map[string]int64

func Copy(og Tree) Tree {
	cp := make(Tree)
	for k, v := range og {
		cp[k] = v
	}
	return cp
}
