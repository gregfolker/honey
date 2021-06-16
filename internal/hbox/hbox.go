package hbox

import (
   "fmt"
   "path/filepath"
   "io/fs"
   "os/exec"
   "strings"

   "github.com/gregfolker/honey/pkg/log"
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
         log.Trace(fmt.Sprintf("Found %s at %s", f.Name(), path))
         if !found {
            found = true
            return nil
         } else {
            return errors.New(fmt.Sprintf("found multiple instances of %s in %s\n", f.Name(), path))
         }
      }

      return nil
   })

   return err
}

func GetHboxFilename(d string) (string, error) {
   file := ""

   err := filepath.Walk(d, func(path string, f fs.FileInfo, err error) error {
      if f.IsDir() {
         return filepath.SkipDir
      }

      if filepath.Ext(f.Name()) == HboxExt {
         log.Trace(fmt.Sprintf("Found %s at %s", f.Name(), path))
         if file == "" {
            file = f.Name()
            return nil
         } else {
            return errors.New(fmt.Sprintf("found multiple *%s files in %s\n", HboxExt, path))
         }
      }

      return nil
   })

   return file, err
}

func RunHboxToMp4(d string, f string) error {
   cmdstr := strings.Join([]string{HboxToMp4, f, "out"}, " ")
   cmd := exec.Command("sh", "-c", cmdstr)

   log.Debug(fmt.Sprintf("Running %s...", cmdstr))

   if stdout, err := cmd.Output(); err != nil {
      return err
   } else {
      fmt.Println(string(stdout))
   }

   return nil
}
