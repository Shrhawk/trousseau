package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/oleiade/trousseau/pkg/trousseau"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a key value pair",
	Long:  ``,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("unable to determine current working directory; reason: %s\n", err.Error())
		}

		s, err := trousseau.Open(filepath.Join(cwd, ".trousseau"))
		if err != nil {
			log.Fatalf("unable to open trousseau data store; reason: %s\n", err.Error())
		}

		s.Set(args[0], args[1])
	},
}
