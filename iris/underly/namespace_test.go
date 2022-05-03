package underly

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"tools/iris/common"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestPatchMergeNamespace(t *testing.T) {
	namespace := "calico"
	kubeConfig := common.ReadFile("/Users/sino/.kube/config")

	patchData := `{"metadata":{"annotations":{"cni.projectcalico.org/ipv4pools":"[\"calicoupgrade\"]"}}}`

	k8sClient, err := common.GetK8sClient(kubeConfig)
	if err != nil {
		t.Error(err)
	}
	underlyNamespace, err := k8sClient.CoreV1().Namespaces().Patch(context.TODO(), namespace, types.MergePatchType, []byte(patchData), metav1.PatchOptions{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(underlyNamespace)
}

func TestNamespaceJoin(t *testing.T) {
	namespaces := []string{
		"a",
		"b",
	}
	fmt.Println(strings.Join(namespaces, `,\"`))
}

func TestMarshal(t *testing.T) {
	a := []string{
		"b",
		"c",
	}
	ret, _ := json.Marshal(a)
	fmt.Println(string(ret))

	var aa []string
	aaa := string(ret)
	err := json.Unmarshal([]byte(aaa), &aa)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(aa)
}

func TestPatchRemoveNamespace(t *testing.T) {
	namespace := "calico"
	kubeConfig := common.ReadFile("/Users/sino/.kube/config")

	patchData := `[{"op":"remove","path":"/metadata/annotations/cni.projectcalico.org~1ipv4pools"}]`
	k8sClient, err := common.GetK8sClient(kubeConfig)
	if err != nil {
		t.Error(err)
	}
	underlyNamespace, err := k8sClient.CoreV1().Namespaces().Patch(context.TODO(), namespace, types.JSONPatchType, []byte(patchData), metav1.PatchOptions{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(underlyNamespace)
}

func TestPatchAddNamespace(t *testing.T) {
	namespace := "calico"
	kubeConfig := common.ReadFile("/Users/sino/.kube/config")

	a := []string{"bar"}
	str, _ := json.Marshal(a)
	fmt.Println(str)
	patchData := `[{"op":"add","path":"/metadata/annotations","value":{"cni.projectcalico.org/ipv4pools":"[\"aaax\"]"}}]`
	k8sClient, err := common.GetK8sClient(kubeConfig)
	if err != nil {
		t.Error(err)
	}
	underlyNamespace, err := k8sClient.CoreV1().Namespaces().Patch(context.TODO(), namespace, types.JSONPatchType, []byte(patchData), metav1.PatchOptions{})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(underlyNamespace)
}
