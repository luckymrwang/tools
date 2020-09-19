package services

import (
	"fmt"
	"tools/iris/common"

	"github.com/emicklei/go-restful"

	"github.com/kataras/iris/v12"
	"k8s.io/client-go/tools/clientcmd"
)

type ContainerService struct {
	Ctx iris.Context
}

func GetContainerService(ctx iris.Context) *ContainerService {
	return &ContainerService{Ctx: ctx}
}

func (s *ContainerService) ExecShell(namespace, pod, container string) (string, error) {
	kubeConfig := ""
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeConfig))
	if err != nil {
		return "", fmt.Errorf("获取 k8s config 失败：%s", err.Error())
	}
	k8sClient, err := common.GetK8sClient(kubeConfig)
	if err != nil {
		return "", fmt.Errorf("获取 k8s 客户端失败：%s", err.Error())
	}
	request := restful.NewRequest(nil)
	request.SetAttribute("namespace", namespace)
	request.SetAttribute("pod", pod)
	request.SetAttribute("container", container)
	return common.ExecShell(k8sClient, restConfig, request)
}
