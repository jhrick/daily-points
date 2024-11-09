package main

import (
	"os"

	"github.com/jhrick/daily-points/internal/storage"
)

func main() {
  args := os.Args[1:]

  if len(args) == 0 {
    return
  }

  storage.CheckStorageFolder()
}
