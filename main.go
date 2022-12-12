package main

import (
  "fmt"
  "os"
  "path/filepath"
  "sync"
  "strings"
)

func renameFilesAndFolders(path string, oldStr string, newStr string) error {
  newPath := strings.Replace(path, oldStr, newStr, -1)
  return os.Rename(path, newPath)
}

func main () {

  var wg sync.WaitGroup

  root := "tree"
  oldStr := "f"
  newStr := "y"
  slice := make([]string, 0)

  filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    if err != nil {
      fmt.Println(err)
      return err
    }

    if strings.Contains(info.Name(), oldStr) {
      wg.Add(1)
      go func() {
        slice = append(slice, path)
        wg.Done()
      }()
    }

    return nil
  })

  wg.Wait()
  if len(slice) > 0 {
    for i := len(slice) - 1; i >= 0; i-- {
      renameFilesAndFolders(slice[i], oldStr, newStr)
    }
  }
}
