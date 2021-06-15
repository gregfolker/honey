package cmd

import (
   "fmt"
   "os"

   "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
   Use:"honey",
   Short:"Honey is a utility to convert video files from HBOX format to MP4 format",
   Run: func(cmd *cobra.Command, args []string) {
      // Do stuff here
   },
}

func Execute() {
   if err := rootCmd.Execute(); err != nil {
      fmt.Fprintln(os.Stderr, err)
      os.Exit(1)
   }
}
