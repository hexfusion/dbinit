package main

import (
	"errors"
	"flag"

	"github.com/golang/glog"
	"github.com/hexfusion/dbutil/pkg"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:     "init --FLAGS",
		Short:   "Initialize a blank etcd db",
		Long:    "This command generates a blank etcd db with only a member and cluster bucket.",
		PreRunE: validateOpts,
		RunE:    runInitCmd,
	}

	initOpts struct {
		dbPath string
		dbName string
	}
	discoveryCmd = &cobra.Command{
		Use:     "discovery --FLAGS",
		Short:   "Populate etcd memberlist with YAML",
		Long:    "This command populates the etcd member bucket with YAML config file.",
		PreRunE: validateOpts,
		RunE:    runInitCmd,
	}
)

func init() {
	rootCmd.AddCommand(dicoveryCmd)
	initCmd.PersistentFlags().StringVar(&initOpts.dbPath, "path", "", "Path to save db file")
	initCmd.PersistentFlags().StringVar(&initOpts.dbPath, "name", "", "Name of the db file")

	rootCmd.AddCommand(
		pkg.NewInitCmd(),
		pkg.NewDiscoveryCommand(),
	)
}

func init() {
	cobra.EnablePrefixMatching = true
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

	c := dbutil.DBConfig{
		Name: initOpts.dbName,
		Path: initOpts.dbPath,
	}

	if err := dbutil.Create(c); err != nil {
		glog.Fatalf("error rendering dbutil config: %v", err)
		return err
	}
	return nil
}
