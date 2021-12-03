package helm

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"os"
)

var (
	settings = cli.New()
)

//NewClient return a new helm client with provided config
func NewClient(namespace string, kfcg KubConfigSetup) (*action.Configuration, error) {
	actionConfig := new(action.Configuration)

	return func(ac *action.Configuration) (*action.Configuration, error) {
		settings.KubeContext = kfcg.Context
		settings.KubeConfig = kfcg.KubeConfigFile
		if namespace == "" {
			namespace = settings.Namespace()
		} else {
			kfcg.Namespace = settings.Namespace()
		}
		err := actionConfig.Init(settings.RESTClientGetter(), namespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
			log.Debug(fmt.Sprintf(format, v))
		})
		if err != nil {
			return nil, err
		}
		return ac, nil

	}(actionConfig)
}
