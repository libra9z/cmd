package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
)

var RootCmd *cobra.Command

var(
	Address	string
	Consul	string
	Token	string
)

func init() {
	_, f := filepath.Split(os.Args[0])
	RootCmd = &cobra.Command{Use:f}
	RootCmd.PersistentFlags().StringVar(&Address, "addr", "", "config this service address format(ip:port).")
	RootCmd.PersistentFlags().StringVar(&Consul, "consul", "", "config consul address format(ip:port).")
	RootCmd.PersistentFlags().StringVar(&Token, "token", "", "config consul acl token.")

	RootCmd.AddCommand(versionCmd)
}



func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
