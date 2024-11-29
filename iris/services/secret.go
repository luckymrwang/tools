package services

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"k8s.io/apimachinery/pkg/types"

	"tools/iris/common"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SecretService struct {
}

func (s *SecretService) Add(kubeconfig string) (*corev1.Secret, error) {
	kubeConfig, err := ioutil.ReadFile(getKubeConfig(kubeconfig))
	if err != nil {
		return nil, fmt.Errorf("读取 kube config 失败：%s", err.Error())
	}
	k8sClient, err := common.GetK8sClient(string(kubeConfig))
	if err != nil {
		return nil, fmt.Errorf("获取 k8s 客户端失败：%s", err.Error())
	}

	username := "username"
	password := "password"

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-secret",
			Namespace: "default",
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{
			"username": []byte(username),
			"password": []byte(password),
		},
	}

	resultSecret, err := k8sClient.CoreV1().Secrets("default").Create(context.TODO(), secret, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		log.Fatalf("Error creating secret: %v", err)
	}
	fmt.Println(common.JSONEncode(resultSecret))

	createdSecret, err := k8sClient.CoreV1().Secrets("default").Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Error creating secret: %v", err)
	}

	secret.Data["password"] = []byte("xxxxx")
	createdSecret.StringData = map[string]string{"username": "xxxxx"}
	patch := fmt.Sprintf(`{"data":%s,"stringData":%s}`, common.JSONEncode(secret.Data), common.JSONEncode(createdSecret.StringData))
	patchSecret, err := k8sClient.CoreV1().Secrets("default").Patch(context.TODO(), createdSecret.Name, types.MergePatchType, []byte(patch), metav1.PatchOptions{})
	if err != nil {
		log.Fatalf("Error creating secret: %v", err)
	}

	return patchSecret, nil
}

func (s *SecretService) Dry(kubeconfig string) (*corev1.Secret, error) {
	kubeConfig, err := ioutil.ReadFile(getKubeConfig(kubeconfig))
	if err != nil {
		return nil, fmt.Errorf("读取 kube config 失败：%s", err.Error())
	}
	k8sClient, err := common.GetK8sClient(string(kubeConfig))
	if err != nil {
		return nil, fmt.Errorf("获取 k8s 客户端失败：%s", err.Error())
	}

	username := "username"
	password := "password"

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "my-secret",
			Namespace: "default",
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{
			"username": []byte(username),
			"password": []byte(password),
		},
	}

	resultSecret, err := k8sClient.CoreV1().Secrets("default").Create(context.TODO(), secret, metav1.CreateOptions{DryRun: []string{metav1.DryRunAll}})
	if err != nil {
		log.Fatalf("Error creating secret: %v", err)
	}
	fmt.Println(common.JSONEncode(resultSecret))

	createdSecret, err := k8sClient.CoreV1().Secrets("default").Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Error creating secret: %v", err)
	}

	return createdSecret, nil
}
