package cmds

import (
	"os"

	"github.com/spf13/cobra"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "ksd",
		Short:             "Kubernetes secret database",
		DisableAutoGenTag: true,
	}

	stopCh := genericapiserver.SetupSignalHandler()
	cmd.AddCommand(NewCmdRun(os.Stdout, os.Stderr, stopCh))

	return cmd
}
