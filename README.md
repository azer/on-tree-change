## on-tree-change

Simple and recursive file watching function based on [fsnotify](https://github.com/fsnotify/fsnotify).

## Install

```bash
$ go get github.com/azer/on-tree-change
```

## Usage

```go
import (
  . "github.com/azer/on-tree-change"
  "os"
)

OnTreeChange("src/", filter, func (filename string) {
    fmt.Println(filename + " updated")
})

func filter (filename string, fileinfo os.FileInfo) bool {
    return info.IsDir()
}
```
