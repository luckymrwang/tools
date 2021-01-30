package services

import (
	"context"
	"fmt"
	"io/ioutil"

	"tools/iris/common"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConfigMapService struct {
}

func (s *ConfigMapService) Add(kubeconfig, namespace, name, key, data string) (*corev1.ConfigMap, error) {
	kubeConfig, err := ioutil.ReadFile(getKubeConfig(kubeconfig))
	if err != nil {
		return nil, fmt.Errorf("读取 kube config 失败：%s", err.Error())
	}
	k8sClient, err := common.GetK8sClient(string(kubeConfig))
	if err != nil {
		return nil, fmt.Errorf("获取 k8s 客户端失败：%s", err.Error())
	}
	configmap := &corev1.ConfigMap{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{Name: name},
		Immutable:  nil,
		Data:       map[string]string{key: data},
		BinaryData: nil,
	}
	configmap, err = k8sClient.CoreV1().ConfigMaps(namespace).Create(context.TODO(), configmap, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return configmap, nil
}

func (s *ConfigMapService) Delete(kubeconfig, namespace, name string) error {
	kubeConfig, err := ioutil.ReadFile(getKubeConfig(kubeconfig))
	if err != nil {
		return fmt.Errorf("读取 kube config 失败：%s", err.Error())
	}
	k8sClient, err := common.GetK8sClient(string(kubeConfig))
	if err != nil {
		return fmt.Errorf("获取 k8s 客户端失败：%s", err.Error())
	}
	err = k8sClient.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}

	return nil
}
