package common

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetK8sClient(credential string) (*kubernetes.Clientset, error) {
	restConfig, err := genRestConfig(credential)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("get clientset err: ", err)
	}
	return clientset, nil
}

func genRestConfig(credential string) (*rest.Config, error) {
	restConfig := &rest.Config{}
	var err error
	if json.Valid([]byte(credential)) {
		var m map[string]string
		err := json.Unmarshal([]byte(credential), &m)
		if err != nil {
			return nil, fmt.Errorf("json unmarshal err: ", err)
		}
		restConfig.Username = m["userName"]
		restConfig.Password = m["password"]
		encode, err := base64.StdEncoding.DecodeString(m["caCerData"])
		if err != nil {
			return nil, fmt.Errorf("base64 encode err: ", err)
		}
		restConfig.TLSClientConfig.CAData = encode
		restConfig.Host = m["apiServerUrl"]
	}
	restConfig, err = clientcmd.RESTConfigFromKubeConfig([]byte(credential))
	if err != nil {
		return nil, fmt.Errorf("build client config err: ", err)
	}
	return restConfig, nil
}

func GetDynamicClient(credential string) (dynamic.Interface, error) {
	restConfig, err := genRestConfig(credential)
	if err != nil {
		return nil, err
	}
	restConfig.APIPath = "/apis"
	dynamicClient, err := dynamic.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("get dynamic client err: ", err)
	}
	return dynamicClient, nil
}
