# filew
A crude file watcher in go based on polling file size change (a rather inexpensive operation). The only time we're opening a file is to read a dir. I'm hoping to collect some cpu-, and memory usage data and see how it relates to number of files watched.

# features
- recursive watcher
- delivers file change events on a channel
- new files added will be watched

# usage
See `cmd/filew/main.go`

# test
```bash
$ go test -bench=.
```

# todo
- [x] tests and mocking
- [x] concurrent diff
- [x] definition of change
- [ ] consider more data points for a file change event
- [ ] fpath flag
- [x] measure mem
- [x] recursive watch
- [ ] try [epoll](https://golang.org/pkg/syscall/#EpollCreate)

# license
MIT
