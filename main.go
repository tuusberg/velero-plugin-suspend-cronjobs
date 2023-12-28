package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tuusberg/velero-plugin-suspend-cronjobs/internal/plugin"
	"github.com/vmware-tanzu/velero/pkg/plugin/framework"
)

func main() {
	framework.NewServer().
		RegisterRestoreItemActionV2("example.io/restore-pluginv2", newRestorePluginV2).
		Serve()
}

func newRestorePluginV2(logger logrus.FieldLogger) (interface{}, error) {
	return plugin.NewRestorePluginV2(logger), nil
}
