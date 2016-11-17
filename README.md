## on-tree-change

Simple and recursive file watching function based on [fsnotify](https://github.com/fsnotify/fsnotify).

## Install

```bash
$ go get github.com/azer/on-change
```

## Usage

Simply call the function with what you wanna watch and the function to be called on change:

```go
import (
  . "github.com/azer/on-change"
)

OnChange("foobar.txt", func (err error) {
    if err != nil {
        panic(err)
    }

  // foobar.txt has been updated
})
```

Optionally, you can pass a filter function to exclude files/folders from the watching. Here is an example;

```go
import (
  . "github.com/azer/on-change"
  "os"
)

OnChange("src/", filter, func (err error) {
    if err != nil {
        panic(err)
    }

  // foobar.txt has been updated
})

func filter (filename string, fileinfo os.FileInfo) bool {
    return info.IsDir()
}
```
