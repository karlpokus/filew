![The mascot](filew_gopher.png)

# filew
A crude file watcher in go based on polling file size change (a rather inexpensive operation). The only time we're opening a file is to read a dir. I'm hoping to collect some cpu-, and memory usage data and see how it relates to number of files watched.

# features
- recursive watcher
- delivers file change events (updated, removed, created) on a channel
- new files added will be watched
- can be used as a lib or a cli

# install
Grab the cli from the [releases tab](https://github.com/karlpokus/filew/releases). Install the lib with the usual `$ go get github.com/karlpokus/filew`

# usage
cli
```bash
# filew writes to stdout so just use a pipe
$ filew | your-script.sh
```

lib
```go
import "github.com/karlpokus/filew"

fpath := "path/to/dir"
events, err := filew.Watch(fpath, nil)
if err != nil {
  // handle err
}
for ev := range events {
  // handle ev
}
```

# test
```bash
$ go test -bench=.
```

# todo
- [x] tests and mocking
- [x] concurrent diff
- [x] definition of change
- [ ] consider more data points for a file change event
- [x] fpath flag
- [x] install and usage docs
- [ ] godoc
- [x] pre-release binary
- [x] version flag
- [x] measure mem
- [x] recursive watch
- [ ] try [epoll](https://golang.org/pkg/syscall/#EpollCreate)

# license
MIT
