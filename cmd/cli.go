package main

import (
	"io"

	"github.com/cyberbliss/kuberang/pkg/config"
	"github.com/cyberbliss/kuberang/pkg/kuberang"
	"github.com/spf13/cobra"
)

// NewKismaticCommand creates the kismatic command
func NewKuberangCommand(version string, in io.Reader, out io.Writer) *cobra.Command {
	var skipCleanup bool
	cmd := &cobra.Command{
		Use:   "kuberang",
		Short: "kuberang tests your kubernetes cluster using kubectl",
		RunE: func(cmd *cobra.Command, args []string) error {
			return doCheckKubernetes(skipCleanup)
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.PersistentFlags().StringVarP(&config.Namespace, "namespace", "n", "",
		"Kubernetes namespace in which kuberang will operate. Defaults to 'default' if not specified.")
	cmd.PersistentFlags().StringVar(&config.RegistryURL, "registry-url", "",
		"Override the default Docker Hub URL to use a local offline registry for required Docker images.")
	cmd.Flags().BoolVar(&skipCleanup, "skip-cleanup", false, "Don't clean up. Leave all deployed artifacts running on the cluster.")

	return cmd
}

func doCheckKubernetes(skipCleanup bool) error {
	return kuberang.CheckKubernetes(skipCleanup)
}
