package cmd

import (
   "fmt"
   "os"

   "github.com/gregfolker/honey/internal/hbox"
   "github.com/spf13/cobra"
   "github.com/pkg/errors"
)

// Application Flags
var (
   path string
   hboxFile string
   verbosity string
)

type RootCmd struct {
   TargetDir string
   HboxFile string
}

var rootCmd = &cobra.Command {
      Use: "honey",
      Short: "Honey is a utility for converting video files from HBOX to MP4",
      PreRunE: func (cmd *cobra.Command, args []string) error {
         return verify()
      },
      RunE: func(cmd *cobra.Command, args []string) error {
         return run()
      },
}

func ExecuteRootCmd() {
   if err := rootCmd.Execute(); err != nil {
      fmt.Fprintln(os.Stderr, err)
      os.Exit(1)
   }
}

func init() {
   rootCmd.PersistentFlags().StringVarP(&verbosity, "log-level", "", "none", "Sets the logging level of the tool")
   rootCmd.PersistentFlags().StringVarP(&path, "target-dir", "t", "", "Path to the directory that contains the *.hbox and *.pll files")
}

func verify() error {
   if path == "" {
      return errors.New(fmt.Sprintf("No target directory specified\n"))
   }

   return nil
}

func run() error {
   var filename string
   var err error

   if err = hbox.HboxToMp4Present(path); err != nil {
      return err
   }

   if filename, err = hbox.GetHboxFilename(path); err != nil {
      return err
   }

   fmt.Println(filename)

   return nil
}
