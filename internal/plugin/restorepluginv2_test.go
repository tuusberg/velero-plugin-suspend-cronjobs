package plugin

import (
	"github.com/sirupsen/logrus"
	"github.com/vmware-tanzu/velero/pkg/plugin/velero"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unsturctured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"reflect"
	"testing"
)

func TestRestorePluginV2_AppliesTo(t *testing.T) {
	t.Run("Only applies to CronJobs", func(t *testing.T) {
		plugin := &RestorePluginV2{
			log: logrus.New(),
		}

		want := velero.ResourceSelector{
			IncludedResources: []string{"cronjobs"},
		}
		got, err := plugin.AppliesTo()
		if err != nil {
			t.Errorf("AppliesTo() error = %v", err)
			return
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("AppliesTo() got = %v, want %v", got, want)
		}
	})
}

func TestRestorePluginV2_Execute(t *testing.T) {
	t.Run("Suspends CronJobs", func(t *testing.T) {
		cronJob := batchv1.CronJob{
			ObjectMeta: v1.ObjectMeta{
				Name:      "test-cronjob",
				Namespace: "test-namespace",
			},
			Spec: batchv1.CronJobSpec{
				Schedule: "0 0 * * *",
				Suspend:  boolPointer(false),
			},
		}

		cronJobUnstructured, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&cronJob)
		if err != nil {
			t.Errorf("Error converting CronJob to unstructured: %v", err)
		}
		input := &velero.RestoreItemActionExecuteInput{
			Item: &unsturctured.Unstructured{
				Object: cronJobUnstructured,
			},
		}

		plugin := &RestorePluginV2{
			log: logrus.New(),
		}

		output, err := plugin.Execute(input)
		if err != nil {
			t.Errorf("Error executing plugin: %v", err)
		}

		got := output.UpdatedItem.UnstructuredContent()["spec"].(map[string]interface{})["suspend"]
		want := true
		if got != want {
			t.Errorf("Execute() got = %v, want %v", got, true)
		}
	})
}
