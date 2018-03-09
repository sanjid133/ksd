package cmds

import (
	"github.com/sanjid133/ksd/pkg/cmd/server"
	"github.com/spf13/cobra"
	"io"
)

func NewCmdRun(out, errOut io.Writer, stopCh <-chan struct{}) *cobra.Command {
	opt := server.NewKsdServerOptions(out, errOut)
	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Launch ksd server",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opt.Complete(); err != nil {
				return err
			}
			if err := opt.Validate(args); err != nil {
				return err
			}
			if err := opt.RunKsdServer(stopCh); err != nil {
				return err
			}

			return nil
		},
	}
	opt.AddFlags(cmd.Flags())

	return cmd
}
