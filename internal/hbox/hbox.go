package hbox

import (
   "fmt"
   "path/filepath"
   "io/fs"

   "github.com/pkg/errors"
)

const (
   HboxExt = ".hbox"
   HboxToMp4 = "HboxToMp4.exe"
)

func HboxToMp4Present(d string) error {
   found := false

   err := filepath.Walk(d, func(path string, f fs.FileInfo, err error) error {
      if f.IsDir() {
         return filepath.SkipDir
      }

      if f.Name() == HboxToMp4 {
         if !found {
            found = true
            return nil
         } else {
            return errors.New(fmt.Sprintf("found multiple instances of %s in %s\n", HboxToMp4, d))
         }
      }

      return nil
   })

   return err
}
