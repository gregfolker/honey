package cmd

import (
   "fmt"
   "os"

   "github.com/spf13/cobra"
)

var (
   TargetDir string
   Verbose bool
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
