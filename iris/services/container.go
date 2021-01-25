package services

import (
	"fmt"
	"io/ioutil"
	"tools/iris/common"

	"github.com/emicklei/go-restful"

	"github.com/kataras/iris/v12"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeConfigPath = "/Users/sino/.kube/config"

type ContainerService struct {
	Ctx iris.Context
}

func GetContainerService(ctx iris.Context) *ContainerService {
	return &ContainerService{Ctx: ctx}
}

func (s *ContainerService) ExecShell(namespace, pod, container string) (string, error) {
	kubeConfig, err := ioutil.ReadFile(kubeConfigPath)
	if err != nil {
		return "", fmt.Errorf("读取 kube config 失败：%s", err.Error())
	}
	restConfig, err := clientcmd.RESTConfigFromKubeConfig(kubeConfig)
	if err != nil {
		return "", fmt.Errorf("获取 k8s config 失败：%s", err.Error())
	}
	k8sClient, err := common.GetK8sClient(string(kubeConfig))
	if err != nil {
		return "", fmt.Errorf("获取 k8s 客户端失败：%s", err.Error())
	}
	request := restful.NewRequest(nil)
	request.SetAttribute("namespace", namespace)
	request.SetAttribute("pod", pod)
	request.SetAttribute("container", container)
	return common.ExecShell(k8sClient, restConfig, request)
}

func (s *ContainerService) CopyFromPod(namespace, pod, container, srcPath string) (string, error) {
	destPath := "./uploads"
	prefix := `require('skyapm-nodejs').start({ serviceName: 'nodejs-demo2', directServers: '10.48.51.135:21594' });`

	kubeConfig, err := ioutil.ReadFile(kubeConfigPath)
	if err != nil {
		return "", fmt.Errorf("读取 kube config 失败：%s", err.Error())
	}
	restConfig, err := clientcmd.RESTConfigFromKubeConfig(kubeConfig)
	if err != nil {
		return "", fmt.Errorf("获取 k8s config 失败：%s", err.Error())
	}
	k8sClient, err := common.GetK8sClient(string(kubeConfig))
	if err != nil {
		return "", fmt.Errorf("获取 k8s 客户端失败：%s", err.Error())
	}
	err = common.CopyFromPod(k8sClient, restConfig, namespace, pod, container, srcPath, destPath)
	if err != nil {
		return "", err
	}
	fileInfo := common.Pathinfo(srcPath)
	localPath := fmt.Sprintf("%s/%v.%v", destPath, fileInfo["filename"], fileInfo["extension"])
	content := common.ReadFile(localPath)
	mergeContent := prefix + "\n" + content
	err = common.WriteFile(localPath, []byte(mergeContent))
	if err != nil {
		return "", err
	}

	return "", nil
}
