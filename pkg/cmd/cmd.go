package cmd

import (
   "fmt"
   "os"

   "honey/internal/hbox"

   "github.com/spf13/cobra"
   "github.com/pkg/errors"
)

type RootCmd struct {
   TargetDir string
   Verbose bool
   HboxFile string
}

func NewRootCmd() *cobra.Command{
   r := &RootCmd{}

   cmd := &cobra.Command{
      Use: "honey",
      Short: "Honey is a utility for converting video files from HBOX to MP4",
      PreRunE: func (cmd *cobra.Command, args []string) error {
         return r.VerifyInputs()
      },
      RunE: func(cmd *cobra.Command, args []string) error {
         // Do stuff here
         return nil
      },
   }

   r.AddFlags(cmd)

   return cmd
}

func ExecuteRootCmd() {
   cmd := NewRootCmd()

   if err := cmd.Execute(); err != nil {
      fmt.Fprintln(os.Stderr, err)
      os.Exit(1)
   }
}

func (r *RootCmd) AddFlags(cmd *cobra.Command) {
   cmd.Flags().StringVarP(&r.TargetDir, "target-dir", "t", "", "Path to the directory that contains the *.hbox and *.pll files")
   cmd.Flags().BoolVarP(&r.Verbose, "verbose", "v", false, "Enable verbose logging")
}

func (cmd *RootCmd) VerifyInputs() error {
   if cmd.TargetDir == "" {
      return errors.New(fmt.Sprintf("No target directory specified\n"))
   }

   if err := hbox.HboxToMp4Present(cmd.TargetDir); err != nil {
      return err
   }

   if file, err := hbox.GetHboxFilename(cmd.TargetDir); err != nil {
      return err
   } else {
      cmd.HboxFile = file
   }

   return nil
}
