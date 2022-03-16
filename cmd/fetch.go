package cmd

import (
	"errors"
	"fmt"
	"github.com/HamzaZo/helm-val/internal/helm"
	"github.com/HamzaZo/helm-val/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli/output"
	"helm.sh/helm/v3/pkg/release"
	"io"
	"strings"
)

type EnvSettings struct {
	KubeConfigFile string
	KubeContext    string
	Namespace      string
}

const valHelp = `
Fetch values from previous release with(or without) specific revision using '--revision'.

Examples:
	
    $ %[1]s val fetch RELEASE-NAME -r/--revision 1

    $ %[1]s val fetch RELEASE-NAME -n/--namespace <ns>
 
    $ %[1]s val fetch RELEASE-NAME -c/--kube-context <ctx>
 
    $ %[1]s val fetch RELEASE-NAME -k/--kubeconfig <kcfg>
`

var (
	revision int
	outfm    output.Format
)

//NewFetchCmd create a fetch command
func NewFetchCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch RELEASE-NAME",
		Short: "Fetch values from previous release",
		Long:  fmt.Sprintf(valHelp, "helm"),
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("release name is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			relValues, err := runFetch(args)
			if err != nil {
				return err
			}
			err = outfm.Write(out, &utils.ValuesWriter{Values: relValues})
			if err != nil {
				return err
			}

			return nil
		},
	}

	flags := cmd.Flags()
	settings.AddFlags(flags)
	flags.IntVarP(&revision, "revision", "r", 0, "number of revision release")
	flags.VarP(utils.ValuesPrinter(output.Table, &outfm), "output", "o",
		fmt.Sprintf("prints the outputs in the specified fomat. Allowed format: %s", strings.Join(output.Formats(), ",")))

	return cmd
}

// AddFlags binds flags to the given flagset
func (e *EnvSettings) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&e.KubeConfigFile, "kubeconfig", "k", "", "path to the kubeconfig file")
	fs.StringVarP(&e.KubeContext, "kube-context", "c", e.KubeContext, "name of the kubeconfig context to use")
	fs.StringVarP(&e.Namespace, "namespace", "n", e.Namespace, "namespace scope for this request")

}

//getLastRelease fetch the latest release
func getLastRelease(release string, ac *action.Get) (*release.Release, error) {
	rel, err := ac.Run(release)
	if err != nil {
		return nil, err
	}
	return rel, nil
}

//getHelmClient returns action configuration based on Helm env
func getHelmClient() (*action.Configuration, error) {
	kubConfig := helm.KubConfigSetup{
		Context:        settings.KubeContext,
		KubeConfigFile: settings.KubeConfigFile,
		Namespace:      settings.Namespace,
	}
	ac, err := helm.NewClient(kubConfig.Namespace, kubConfig)
	if err != nil {
		return nil, err
	}
	return ac, nil
}

//runFetch retrieve values for a given old release
func runFetch(args []string) (map[string]interface{}, error) {
	releaseName := args[0]

	ac, err := getHelmClient()
	if err != nil {
		return nil, err
	}
	p := action.NewGet(ac)
	rel, err := getLastRelease(releaseName, p)
	if err != nil {
		return nil, err
	}

	var previousRelease int
	if revision == 0 {
		previousRelease = rel.Version - 1
	} else {
		previousRelease = revision
	}
	gVal := action.NewGetValues(ac)
	gVal.Version = previousRelease
	gVal.AllValues = true

	relVal, err := gVal.Run(rel.Name)
	if err != nil {
		return nil, err
	}

	return relVal, nil
}
