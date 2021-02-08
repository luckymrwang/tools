package services

import (
	"context"
	"fmt"
	"io/ioutil"
	"tools/iris/common"

	"github.com/kataras/iris/v12"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentService struct {
	Ctx iris.Context
}

func GetDeploymentService(ctx iris.Context) *DeploymentService {
	return &DeploymentService{Ctx: ctx}
}

func (s *DeploymentService) InjectNodejs(kubeconfig, namespace, name, srcPath string) (*v1.Deployment, error) {
	kubeConfig, err := ioutil.ReadFile(getKubeConfig(kubeconfig))
	if err != nil {
		return nil, fmt.Errorf("读取 kube config 失败：%s", err.Error())
	}
	k8sClient, err := common.GetK8sClient(string(kubeConfig))
	if err != nil {
		return nil, fmt.Errorf("获取 k8s 客户端失败：%s", err.Error())
	}
	deploy, err := k8sClient.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	containerImage := "skyapm:v2.0.1"
	imagePullPolicy := "IfNotPresent"
	if !common.IsEmpty(kubeconfig) {
		containerImage = fmt.Sprintf("10.48.51.135:5000/com.inspur/incloudos-docker/%s", containerImage)
		imagePullPolicy = "Always"
	}
	initContainer := corev1.Container{
		Name:    "sidecar",
		Image:   containerImage,
		Command: []string{"/bin/sh", "-c", "cp -r /node_modules/* /node/modules"},
		VolumeMounts: []corev1.VolumeMount{{
			Name:      "sidecar",
			MountPath: "/node/modules",
		}},
		ImagePullPolicy: corev1.PullPolicy(imagePullPolicy),
	}
	podSpec := deploy.Spec.Template.Spec
	if len(podSpec.InitContainers) == 0 {
		podSpec.InitContainers = []corev1.Container{initContainer}
	} else {
		podSpec.InitContainers = append(podSpec.InitContainers, initContainer)
	}

	// skyapm
	volume := corev1.Volume{
		Name: "sidecar",
		VolumeSource: corev1.VolumeSource{
			EmptyDir: &corev1.EmptyDirVolumeSource{},
		},
	}
	if len(podSpec.Volumes) == 0 {
		podSpec.Volumes = []corev1.Volume{volume}
	} else {
		podSpec.Volumes = append(podSpec.Volumes, volume)
	}
	volumeMount := corev1.VolumeMount{
		Name:      "sidecar",
		MountPath: "/node_modules",
	}
	// TODO:more than one container
	// volumeMounts add sidecar
	if len(podSpec.Containers[0].VolumeMounts) == 0 {
		podSpec.Containers[0].VolumeMounts = []corev1.VolumeMount{volumeMount}
	} else {
		podSpec.Containers[0].VolumeMounts = append(podSpec.Containers[0].VolumeMounts, volumeMount)
	}

	fileInfo := common.Pathinfo(srcPath)
	fileName := fmt.Sprintf("%v.%v", fileInfo["filename"], fileInfo["extension"])
	configmapName := fmt.Sprintf("%s-nodejs", name)
	// volumeMounts add configmap
	podSpec.Containers[0].VolumeMounts = append(podSpec.Containers[0].VolumeMounts, corev1.VolumeMount{
		Name:      configmapName,
		MountPath: srcPath,
		SubPath:   fileName,
	})
	mode := int32(420)
	// volumes add cofigmap
	podSpec.Volumes = append(podSpec.Volumes, corev1.Volume{
		Name: configmapName,
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{Name: configmapName},
				Items: []corev1.KeyToPath{{
					Key:  fileName,
					Path: fileName,
				}},
				DefaultMode: &mode,
			},
		},
	})
	deploy.Spec.Template.Spec = podSpec
	deploy, err = k8sClient.AppsV1().Deployments(namespace).Update(context.TODO(), deploy, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	return deploy, nil
}

func (s *DeploymentService) InjectSidecar(kubeconfig, namespace, name, srcPath string) (*v1.Deployment, error) {
	kubeConfig, err := ioutil.ReadFile(getKubeConfig(kubeconfig))
	if err != nil {
		return nil, fmt.Errorf("读取 kube config 失败：%s", err.Error())
	}
	k8sClient, err := common.GetK8sClient(string(kubeConfig))
	if err != nil {
		return nil, fmt.Errorf("获取 k8s 客户端失败：%s", err.Error())
	}
	deploy, err := k8sClient.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	containerImage := "skywalking-agent:8.3.0"
	imagePullPolicy := "IfNotPresent"
	if !common.IsEmpty(kubeconfig) {
		containerImage = fmt.Sprintf("10.48.51.135:5000/com.inspur/incloudos-docker/%s", containerImage)
		imagePullPolicy = "Always"
	}
	initContainer := corev1.Container{
		Name:    "sidecar",
		Image:   containerImage,
		Command: []string{"cp", "-r", "/agent", "/sidecar"},
		VolumeMounts: []corev1.VolumeMount{{
			Name:      "sidecar",
			MountPath: "/sidecar",
		}},
		ImagePullPolicy: corev1.PullPolicy(imagePullPolicy),
	}
	podSpec := deploy.Spec.Template.Spec
	if len(podSpec.InitContainers) == 0 {
		podSpec.InitContainers = []corev1.Container{initContainer}
	} else {
		podSpec.InitContainers = append(podSpec.InitContainers, initContainer)
	}

	// java agent
	volume := corev1.Volume{
		Name: "sidecar",
		VolumeSource: corev1.VolumeSource{
			EmptyDir: &corev1.EmptyDirVolumeSource{},
		},
	}
	if len(podSpec.Volumes) == 0 {
		podSpec.Volumes = []corev1.Volume{volume}
	} else {
		podSpec.Volumes = append(podSpec.Volumes, volume)
	}
	volumeMount := corev1.VolumeMount{
		Name:      "sidecar",
		MountPath: "/sidecar",
	}
	// TODO:more than one container
	// volumeMounts add sidecar
	if len(podSpec.Containers[0].VolumeMounts) == 0 {
		podSpec.Containers[0].VolumeMounts = []corev1.VolumeMount{volumeMount}
	} else {
		podSpec.Containers[0].VolumeMounts = append(podSpec.Containers[0].VolumeMounts, volumeMount)
	}

	// java envs
	envs := []corev1.EnvVar{
		{
			Name:  "JAVA_OPTIONS",
			Value: "-javaagent:/sidecar/agent/skywalking-agent.jar",
		},
		{
			Name:  "SW_AGENT_NAME",
			Value: name,
		},
		{
			Name:  "SW_AGENT_COLLECTOR_BACKEND_SERVICES",
			Value: "skywalking-oap.istio-system:11800",
		},
	}
	// add envs
	if len(podSpec.Containers[0].Env) == 0 {
		podSpec.Containers[0].Env = envs
	} else {
		podSpec.Containers[0].Env = append(podSpec.Containers[0].Env, envs...)
	}

	deploy.Spec.Template.Spec = podSpec
	deploy, err = k8sClient.AppsV1().Deployments(namespace).Update(context.TODO(), deploy, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}

	return deploy, nil
}
