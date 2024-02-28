package controller

import (
	"context"
	"kubevision/internal/model"

	"github.com/BytemanD/easygo/pkg/global/logging"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sClient struct {
	client *kubernetes.Clientset
}

func (c K8sClient) ListNamespaces() ([]model.Namespace, error) {
	logging.Info("list namespaces")
	namespaces := []model.Namespace{}
	nsList, err := c.client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range nsList.Items {
		namespaces = append(namespaces, model.ParseV1Namespce(item))
	}
	return namespaces, nil
}
func (c K8sClient) ListNodes() ([]model.Node, error) {
	nodes := []model.Node{}
	items, err := c.client.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range items.Items {
		nodes = append(nodes, model.ParseV1Node(item))
	}
	return nodes, nil
}
func (c K8sClient) ListPods(namespace string) ([]model.Pod, error) {
	logging.Info("list pods with namespace: %s", namespace)
	pods := []model.Pod{}
	posList, err := c.client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range posList.Items {
		pods = append(pods, model.ParseV1Pod(item))
	}
	return pods, nil
}
func (c K8sClient) ListDaemonsets(namespace string) ([]model.Daemonset, error) {
	daemonsets := []model.Daemonset{}
	items, err := c.client.AppsV1().DaemonSets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range items.Items {
		daemonsets = append(daemonsets, model.ParseV1Daemonset(item))
	}
	return daemonsets, nil
}

func (c K8sClient) ListDeployments(namepace string) ([]model.Deployment, error) {
	nodes := []model.Deployment{}
	items, err := c.client.AppsV1().Deployments(namepace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range items.Items {
		nodes = append(nodes, model.ParseV1Deployment(item))
	}
	return nodes, nil
}

var k8sClient *K8sClient

func NewClient() (*K8sClient, error) {
	if k8sClient == nil {
		logging.Info("init client")
		config, err := clientcmd.BuildConfigFromFlags("", "/etc/kubevision/k8sconfig")
		if err != nil {
			return nil, err
		}
		clientSet, err := kubernetes.NewForConfig(config)
		if err != nil {
			return nil, err
		}
		k8sClient = &K8sClient{client: clientSet}
	}
	return k8sClient, nil
}
