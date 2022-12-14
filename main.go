package main

import (
  "fmt"
  "os"
  "path/filepath"
  "sync"
  "strings"
)


/*func renameFilesAndFolders(path string, oldStr string, newStr string) error {
  newPath := strings.Replace(path, oldStr, newStr, -1)
  os.Rename(path, newPath)
  return nil
}*/


func renameFilesAndFolders(path string, oldStr string, newStr string) error {
  newPath := strings.Replace(path, oldStr, newStr, -1)
  err := os.RemoveAll(path)
  if err != nil {
    fmt.Println(err)
  }
  errMkdir := os.MkdirAll(newPath, 0750)
	if errMkdir != nil && !os.IsExist(errMkdir) {
		fmt.Println(errMkdir)
	}
  return nil
}

func main () {

  var wg sync.WaitGroup

  // directory, string to be replaced and string to replace it
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
    for _, path := range slice {
      renameFilesAndFolders(path, oldStr, newStr)
    }
  }
}
