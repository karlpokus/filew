# filew
A crude file watcher in go. Let's see how expensive polling really is.

# usage
See `cmd/filew/main.go`

# test
```bash
$ go test # broken atm
```

# todo
- [ ] tests and mocking
- [ ] definition of change
- [ ] measure mem
- [x] recursive watch
- [ ] try [epoll](https://golang.org/pkg/syscall/#EpollCreate)

# license
MIT
