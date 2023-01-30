package services

import (
	"fmt"
	"tools/iris/common"

	"github.com/kataras/iris/v12"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	rootPath     = "kapis"
	groupVersion = "resources.icks/v1"

	kubeAggregation = "ks-apiserver.kube-aggregation.svc.cluster.local"
)

type VerifyChain interface {
	Chain(runtimeCredentialContent string, rs *unstructured.Unstructured) error
}

type BaseService struct {
	chains []VerifyChain
}

func (s *BaseService) Proxy(ctx iris.Context, runtimeId string, conditions ...common.P) (string, error) {
	for _, val := range ctx.URLParams() {
		fmt.Println(runtimeId, val)
	}
	fmt.Println(conditions)
	return "", nil
}

func (s *BaseService) VerifyChainDryRun(runtimeCredentialContent string, rs *unstructured.Unstructured, update bool) error {
	fmt.Println(runtimeCredentialContent, rs, update)
	return nil
}

func (s *BaseService) CreateResource(runtimeCredentialContent string, rs *unstructured.Unstructured) error {
	fmt.Println(runtimeCredentialContent, rs)
	return nil
}

func (s *BaseService) UpdateResource(runtimeCredentialContent string, rs *unstructured.Unstructured) error {
	fmt.Println(runtimeCredentialContent, rs)
	return nil
}

func (s *BaseService) DeleteResource(runtimeCredentialContent string, rs *unstructured.Unstructured) error {
	fmt.Println(runtimeCredentialContent, rs)
	return nil
}

func (s *BaseService) SetAnnotations(rs *unstructured.Unstructured, annotations map[string]string) error {
	fmt.Println(rs, annotations)
	return nil
}
