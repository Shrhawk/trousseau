package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/oleiade/trousseau/pkg/trousseau"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes an encrypted data store in the current working directory",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("unable to determine current working directory; reason: %s\n", err.Error())
		}

		storepath := filepath.Join(cwd, ".trousseau")
		if _, err := os.Stat(storepath); !os.IsNotExist(err) {
			log.Fatalf("A trousseau store already exists in %s\n", storepath)
		}

		s := trousseau.NewStore()
		err = s.WriteTo(filepath.Join(cwd, ".trousseau"))
		if err != nil {
			log.Fatal(err)
		}
	},
}
