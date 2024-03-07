package k8s

import (
	"context"
	"kubevision/internal/model"

	"github.com/BytemanD/easygo/pkg/global/logging"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type K8sClient struct {
	restConfig rest.Config
	apiConfig  api.Config
	client     *kubernetes.Clientset
}

func (c K8sClient) GetClusterInfo() (*model.Cluster, error) {
	serverVersion, err := c.GetServerVersion()
	if err != nil {
		return nil, err
	}

	return &model.Cluster{
		Host:           c.restConfig.Host,
		ApiPath:        c.restConfig.APIPath,
		ServerVersion:  serverVersion,
		CurrentContext: c.apiConfig.CurrentContext,
	}, nil
}

func (c K8sClient) GetServerVersion() (*version.Info, error) {
	return c.client.ServerVersion()
}

func (c K8sClient) ListNamespaces() ([]model.Namespace, error) {
	nsList, err := c.client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	namespaces := []model.Namespace{}
	for _, item := range nsList.Items {
		namespaces = append(namespaces, model.ParseV1Namespce(item))
	}
	return namespaces, nil
}
func (c K8sClient) ListNodes() ([]model.Node, error) {
	items, err := c.client.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	nodes := []model.Node{}
	for _, item := range items.Items {
		nodes = append(nodes, model.ParseV1Node(item))
	}
	return nodes, nil
}
func (c K8sClient) ListPods(namespace string) ([]model.Pod, error) {
	posList, err := c.client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	pods := []model.Pod{}
	for _, item := range posList.Items {
		pods = append(pods, model.ParseV1Pod(item))
	}
	return pods, nil
}
func (c K8sClient) ListDaemonsets(namespace string) ([]model.Daemonset, error) {
	items, err := c.client.AppsV1().DaemonSets(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	daemonsets := []model.Daemonset{}
	for _, item := range items.Items {
		daemonsets = append(daemonsets, model.ParseV1Daemonset(item))
	}
	return daemonsets, nil
}
func (c K8sClient) GetDaemonsets(namespace string, name string) (*model.Daemonset, error) {
	item, err := c.client.AppsV1().DaemonSets(namespace).Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	ds := model.ParseV1Daemonset(*item)
	return &ds, nil
}
func (c K8sClient) ListDeployments(namepace string) ([]model.Deployment, error) {
	items, err := c.client.AppsV1().Deployments(namepace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	nodes := []model.Deployment{}
	for _, item := range items.Items {
		nodes = append(nodes, model.ParseV1Deployment(item))
	}
	return nodes, nil
}
func (c K8sClient) ListServices(namepace string) ([]model.Service, error) {
	items, err := c.client.CoreV1().Services(namepace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	services := []model.Service{}
	for _, item := range items.Items {
		services = append(services, model.ParseV1Service(item))
	}
	return services, nil
}
func (c K8sClient) ListStatefulSets(namepace string) ([]model.StatefulSet, error) {
	items, err := c.client.AppsV1().StatefulSets(namepace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	statefulSets := []model.StatefulSet{}
	for _, item := range items.Items {
		statefulSets = append(statefulSets, model.ParseV1StatefulSet(item))
	}
	return statefulSets, nil
}
func (c K8sClient) ListJobs(namepace string) ([]model.Job, error) {
	items, err := c.client.BatchV1().Jobs(namepace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	jobs := []model.Job{}
	for _, item := range items.Items {
		jobs = append(jobs, model.ParseV1Job(item))
	}
	return jobs, nil
}
func (c K8sClient) ListBatchBetaV1CronJobs(namepace string) ([]model.CronJob, error) {
	cronJobs := []model.CronJob{}
	items1, err := c.client.BatchV1beta1().CronJobs(namepace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range items1.Items {
		cronJobs = append(cronJobs, model.ParseV1betaCronJob(item))
	}
	return cronJobs, nil
}
func (c K8sClient) ListBatchV1CronJobs(namepace string) ([]model.CronJob, error) {
	cronJobs := []model.CronJob{}
	items1, err := c.client.BatchV1beta1().CronJobs(namepace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range items1.Items {
		cronJobs = append(cronJobs, model.ParseV1betaCronJob(item))
	}
	return cronJobs, nil
}
func (c K8sClient) ListCronJobs(namepace string) ([]model.CronJob, error) {
	cronJobs, err := c.ListBatchBetaV1CronJobs(namepace)
	if err != nil {
		cronJobs, err = c.ListBatchV1CronJobs(namepace)
	}
	if err != nil {
		return nil, err
	}
	return cronJobs, nil
}
func (c K8sClient) ListEvents(namepace string) ([]model.Event, error) {
	items, err := c.client.CoreV1().Events(namepace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	events := []model.Event{}
	for _, item := range items.Items {
		events = append(events, model.ParseV1Event(item))
	}
	return events, nil
}
func (c K8sClient) ListConfigMaps(namepace string) ([]model.ConfigMap, error) {
	items, err := c.client.CoreV1().ConfigMaps(namepace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	cms := []model.ConfigMap{}
	for _, item := range items.Items {
		cms = append(cms, model.ParseV1ConfigMap(item))
	}
	return cms, nil
}
func (c K8sClient) ListSecrets(namepace string) ([]model.Secret, error) {
	items, err := c.client.CoreV1().Secrets(namepace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	events := []model.Secret{}
	for _, item := range items.Items {
		events = append(events, model.ParseV1Secret(item))
	}
	return events, nil
}

var k8sClient *K8sClient

func GetClient() (*K8sClient, error) {
	if k8sClient == nil {
		configPath := "/etc/kubevision/k8sconfig"
		logging.Info("init client")
		config, err := clientcmd.BuildConfigFromFlags("", configPath)
		if err != nil {
			return nil, err
		}
		apiConfig, err := clientcmd.LoadFromFile(configPath)
		if err != nil {
			return nil, err
		}
		clientSet, err := kubernetes.NewForConfig(config)
		if err != nil {
			return nil, err
		}
		k8sClient = &K8sClient{
			restConfig: *config,
			apiConfig:  *apiConfig,
			client:     clientSet}
	}
	return k8sClient, nil
}
