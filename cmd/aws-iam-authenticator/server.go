/*
Copyright 2017 by the contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/kubernetes-sigs/aws-iam-authenticator/pkg/server"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// DefaultPort is the default localhost port (chosen randomly).
const DefaultPort = 21362

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run a webhook validation server suitable that validates tokens using AWS IAM",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := getConfig()

		if err != nil {
			logrus.Fatalf("%s", err)
		}

		server.New(config).Run()
	},
}

func init() {
	viper.SetDefault("server.port", DefaultPort)

	serverCmd.Flags().String("generate-kubeconfig",
		"/etc/kubernetes/aws-iam-authenticator/kubeconfig.yaml",
		"Output `path` where a generated webhook kubeconfig (for `--authentication-token-webhook-config-file`) will be stored (should be a hostPath mount).")
	viper.BindPFlag("server.generateKubeconfig", serverCmd.Flags().Lookup("generate-kubeconfig"))

	serverCmd.Flags().Bool("kubeconfig-pregenerated",
		false,
		"set to `true` when a webhook kubeconfig is pre-generated by running the `init` command, and therefore the `server` shouldn't unnecessarily re-generate a new one.")
	viper.BindPFlag("server.kubeconfigPregenerated", serverCmd.Flags().Lookup("kubeconfig-pregenerated"))

	serverCmd.Flags().String("state-dir",
		"/var/aws-iam-authenticator",
		"State `directory` for generated certificate and private key (should be a hostPath mount).")
	viper.BindPFlag("server.stateDir", serverCmd.Flags().Lookup("state-dir"))

	serverCmd.Flags().StringP(
		"bind",
		"b",
		"127.0.0.1",
		"IP Address to bind the server to listen to. (should be a 127.0.0.1 or 0.0.0.0)")
	viper.BindPFlag("server.bind", serverCmd.Flags().Lookup("bind"))

	rootCmd.AddCommand(serverCmd)
}
