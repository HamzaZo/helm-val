package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"io"
	"os"
)

const rootCmdUsage = `
The val plugin helps you to get values from an old release
`

var settings *EnvSettings

//NewRootCmd creates a root cmd
func NewRootCmd(out io.Writer, args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "val",
		Short:        "Fetch values from previous release",
		Long:         rootCmdUsage,
		SilenceUsage: true,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 1 {
				return errors.New("no argument accepted")
			}
			return nil
		},
	}
	flags := cmd.PersistentFlags()
	flags.Parse(args)

	settings = new(EnvSettings)
	if ctx := os.Getenv("HELM_KUBECONTEXT"); ctx != "" {
		settings.KubeContext = ctx
	}

	cmd.AddCommand(NewFetchCmd(out))

	return cmd
}
