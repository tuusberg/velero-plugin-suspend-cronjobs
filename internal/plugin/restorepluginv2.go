package plugin

import (
	"github.com/sirupsen/logrus"
	v1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	"github.com/vmware-tanzu/velero/pkg/plugin/velero"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// RestorePluginV2 is a restore item action plugin for Velero
type RestorePluginV2 struct {
	log logrus.FieldLogger
}

// NewRestorePluginV2 instantiates a v2 RestorePlugin.
func NewRestorePluginV2(log logrus.FieldLogger) *RestorePluginV2 {
	return &RestorePluginV2{log: log}
}

// Name is required to implement the interface, but the Velero pod does not delegate this
// method -- it's used to tell velero what name it was registered under. The plugin implementation
// must define it, but it will never actually be called.
func (p *RestorePluginV2) Name() string {
	return "velero-plugin-suspend-cronjobs"
}

// AppliesTo returns information about which resources this action should be invoked for.
// The IncludedResources and ExcludedResources slices can include both resources
// and resources with group names. These work: "ingresses", "ingresses.extensions".
// A RestoreItemAction's Execute function will only be invoked on items that match the returned
// selector. A zero-valued ResourceSelector matches all resources.
func (p *RestorePluginV2) AppliesTo() (velero.ResourceSelector, error) {
	return velero.ResourceSelector{
		IncludedResources: []string{"cronjobs"},
	}, nil
}

// Execute allows the RestorePlugin to perform arbitrary logic with the item being restored,
// in this case, suspending the CronJob being restored.
func (p *RestorePluginV2) Execute(input *velero.RestoreItemActionExecuteInput) (*velero.RestoreItemActionExecuteOutput, error) {
	// Since the resource selector we defined in AppliesTo() only matches
	// CronJob resources, we can safely cast the input.Item to a *v1.CronJob.
	var cronJob batchv1.CronJob
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(input.Item.UnstructuredContent(), &cronJob)
	if err != nil {
		return nil, err
	}

	// This is where the job gets suspended:
	p.log.Infof("CronJob %s/%s will be suspended", cronJob.Namespace, cronJob.Name)
	cronJob.Spec.Suspend = Pointer(true)

	// Convert the CronJob back to unstructured data:
	cronJobUnstructured, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&cronJob)
	if err != nil {
		return nil, err
	}
	input.Item.SetUnstructuredContent(cronJobUnstructured)
	return velero.NewRestoreItemActionExecuteOutput(input.Item), nil
}

func (p *RestorePluginV2) Progress(_ string, _ *v1.Restore) (velero.OperationProgress, error) {
	return velero.OperationProgress{Completed: true}, nil
}

func (p *RestorePluginV2) Cancel(operationID string, restore *v1.Restore) error {
	return nil
}

func (p *RestorePluginV2) AreAdditionalItemsReady(additionalItems []velero.ResourceIdentifier, restore *v1.Restore) (bool, error) {
	return true, nil
}
