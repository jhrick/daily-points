package storage

import (
	"fmt"
	"io/fs"
	"os"
)

var homePath string = "/home/" + os.Getenv("USER") 

func CheckStorageFolder() {
  os.Chdir(homePath)
 
  _, err := os.Stat(".points")
  if os.IsNotExist(err) {
    setStorage()
    return
  }

  fmt.Println("folder already exist")
}

func setStorage() {
  var filePerm fs.FileMode = 0022

  err := os.Mkdir(".points", filePerm)
  if err != nil {
    panic(fmt.Sprintf("mkdir error: %v", err))
  }

  _, err = os.Stat(".points")
  if os.IsNotExist(err) {
    panic("cannot create folder")
  }

  fmt.Println("created!")
}
