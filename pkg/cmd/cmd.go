package cmd

import (
   "fmt"
   "os"

   "honey/internal/hbox"

   "github.com/spf13/cobra"
)

var (
   TargetDir string
   Verbose bool
   HboxFile string
)

var rootCmd = &cobra.Command{
   Use: "honey",
   Short: "Honey is a utility to convert video files from HBOX format to MP4 format",
   Run: func(cmd *cobra.Command, args []string) {
      // Do stuff here
   },
}

func Execute() {
   AddFlags()

   if err := VerifyInputs(); err != nil {
      fmt.Fprintln(os.Stderr, err)
      os.Exit(1)
   }

   if err := rootCmd.Execute(); err != nil {
      fmt.Fprintln(os.Stderr, err)
      os.Exit(1)
   }
}

func AddFlags() {
   rootCmd.PersistentFlags().StringVarP(&TargetDir, "target-dir", "t", "", "Path to the directory that contains the *.hbox and *.pll files")
   rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Enable verbose logging")

   rootCmd.MarkPersistentFlagRequired("target-dir")
}

func VerifyInputs() error {
   if err := hbox.HboxToMp4Present(TargetDir); err != nil {
      return err
   }

   if file, err := hbox.GetHboxFilename(TargetDir); err != nil {
      return err
   } else {
      HboxFile = file
   }

   return nil
}
