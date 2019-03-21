package main

import (
	"errors"
	"flag"

	"github.com/golang/glog"
	"github.com/hexfusion/dbinit/pkg/dbinit"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:     "init --FLAGS",
		Short:   "Initialize an etcd db",
		Long:    "This command generates a blank etcd db with only a member and cluster bucket.",
		PreRunE: validateOpts,
		RunE:    runInitCmd,
	}

	initOpts struct {
		dbPath string
		dbName string
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().StringVar(&initOpts.dbPath, "path", "", "Path to save db file")
	initCmd.PersistentFlags().StringVar(&initOpts.dbPath, "name", "", "Name of the db file")
}

// validateServeOpts validates the user flag values given to the signer server
func validateOpts(cmd *cobra.Command, args []string) error {
	if initOpts.dbPath == "" || initOpts.dbName == "" {
		return errors.New("both --path and --name are required flags")
	}
	return nil
}

func runInitCmd(cmd *cobra.Command, args []string) error {
	// flag.Set("logtostderr", "true")
	flag.Parse()

	c := dbinit.DBConfig{
		Name: initOpts.dbName,
		Path: initOpts.dbPath,
	}

	if err := dbinit.Create(c); err != nil {
		glog.Fatalf("error rendering dbinit config: %v", err)
		return err
	}
	return nil
}
