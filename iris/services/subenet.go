package services

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/kataras/iris/v12"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"tools/iris/common"
)

type SubentService struct {
	Ctx iris.Context
}

func GetSubentService(ctx iris.Context) *SubentService {
	return &SubentService{Ctx: ctx}
}

func (s *SubentService) GetClusterCIDR(kubeconfig string) (common.P, error) {
	kubeConfig, err := os.ReadFile(getKubeConfig(kubeconfig))
	if err != nil {
		return nil, fmt.Errorf("读取 kube config 失败：%s", err.Error())
	}
	k8sClient, err := common.GetK8sClient(string(kubeConfig))
	if err != nil {
		return nil, fmt.Errorf("获取 k8s 客户端失败：%s", err.Error())
	}

	configmap, err := k8sClient.CoreV1().ConfigMaps("kube-system").Get(context.TODO(), "kube-proxy", metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(bytes.NewBufferString(configmap.Data["config.conf"]))
	if err != nil {
		return nil, err
	}
	clusterCIDRv4, clusterCIDRv6 := s.getClusterCIDR(viper.GetString("clusterCIDR"))
	return common.P{
		"clusterCIDRv4": clusterCIDRv4,
		"clusterCIDRv6": clusterCIDRv6,
	}, nil
}

func (s *SubentService) getClusterCIDR(address string) (clusterCIDRv4, clusterCIDRv6 string) {
	ips := strings.Split(address, ",")
	if len(ips) == 2 {
		v4IP := net.ParseIP(strings.Split(ips[0], "/")[0])
		v6IP := net.ParseIP(strings.Split(ips[1], "/")[0])
		if v4IP.To4() != nil && v6IP.To16() != nil {
			clusterCIDRv4 = ips[0]
			clusterCIDRv6 = ips[1]
			return
		}
		v4IP = net.ParseIP(strings.Split(ips[1], "/")[0])
		v6IP = net.ParseIP(strings.Split(ips[0], "/")[0])
		if v4IP.To4() != nil && v6IP.To16() != nil {
			clusterCIDRv4 = ips[1]
			clusterCIDRv6 = ips[0]
			return
		}
		return
	}
	ip := net.ParseIP(strings.Split(address, "/")[0])
	if ip.To4() != nil {
		clusterCIDRv4 = address
	} else if ip.To16() != nil {
		clusterCIDRv6 = address
	}
	return
}

func (s *SubentService) WhetherContains(subnetA, subnetB string) bool {
	return common.ContainsCIDR(subnetA, subnetB)
}
