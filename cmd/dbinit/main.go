package main

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dbinit",
		Short: "etcd database initalizer",
		Long:  "",
	}
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		glog.Exitf("Error executing dbinit: %v", err)
	}

}
