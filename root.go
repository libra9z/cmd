package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
)

var RootCmd *cobra.Command

var(
	Address			string
	Consul			string
	Token			string
	Debug			bool
	DebugAddr		string
	RecordAddr		string
	HttpAddr		string
	ZipkinAddr		string
	AppdashAddr		string
	LightstepToken	string
)

func init() {
	_, f := filepath.Split(os.Args[0])
	RootCmd = &cobra.Command{Use:f}
	RootCmd.PersistentFlags().StringVar(&Address, "addr", "", "config this service address format(ip:port).")
	RootCmd.PersistentFlags().StringVar(&Consul, "consul", "", "config consul address format(ip:port).")
	RootCmd.PersistentFlags().StringVar(&Token, "token", "", "config consul acl token.")

	RootCmd.PersistentFlags().BoolVar(&Debug,"debug", false, "enable debug mode true or false")
	RootCmd.PersistentFlags().StringVar(&DebugAddr,"debug.addr", ":8080", "Debug and metrics listen address")
	RootCmd.PersistentFlags().StringVar(&HttpAddr,"http.addr", ":8081", "HTTP listen address")
	RootCmd.PersistentFlags().StringVar(&RecordAddr,"record.addr", "", "zipkin recorder address")
	RootCmd.PersistentFlags().StringVar(&ZipkinAddr,"zipkin.addr", "", "Enable Zipkin tracing via a Kafka server host:port")
	RootCmd.PersistentFlags().StringVar(&AppdashAddr,"appdash.addr", "", "Enable Appdash tracing via an Appdash server host:port")
	RootCmd.PersistentFlags().StringVar(&LightstepToken,"lightstep.token", "", "Enable LightStep tracing via a LightStep access token")

	RootCmd.AddCommand(versionCmd)
}



func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
