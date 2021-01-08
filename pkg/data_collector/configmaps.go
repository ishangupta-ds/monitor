package data_collector

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type ConfigMapClient struct {
	Clientset *kubernetes.Clientset
}

type ConfigMapsPerNamespaceDetails struct {
	Namespace      string
	ConfigMapsInfo []ConfigMapInfo
}

type ConfigMapInfo struct {
	Name              string
	UID               types.UID
	Namespace         string
	APIVersion        string
	ClusterName       string
	CreationTimestamp metav1.Time
	DataSize          int
	BinaryDataSize    int
	Labels            map[string]string
	Annotations       map[string]string
}

func GetAllConfigMapByNamespace(cmc *ConfigMapClient, namespace string) *ConfigMapsPerNamespaceDetails {

	configMapList, err := cmc.Clientset.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{})

	if err != nil {
		fmt.Println("Failed to fetch configMapList for namespace: ", namespace)
	}

	cmnsd := ConfigMapsPerNamespaceDetails{}
	cmnsd.Namespace = namespace

	for _, cm := range configMapList.Items {
		configMapInfo := ConfigMapInfo{
			Name:              cm.GetName(),
			UID:               cm.GetUID(),
			Namespace:         cm.GetNamespace(),
			APIVersion:        cm.APIVersion,
			ClusterName:       cm.GetClusterName(),
			CreationTimestamp: cm.GetCreationTimestamp(),
			DataSize:          len(cm.Data),
			BinaryDataSize:    len(cm.BinaryData),
			Labels:            cm.GetLabels(),
			Annotations:       cm.GetAnnotations(),
		}
		cmnsd.ConfigMapsInfo = append(cmnsd.ConfigMapsInfo, configMapInfo)
	}

	return &cmnsd
}

func AllConfigMapsPerNamespace(clientset *kubernetes.Clientset, namespaces []string) map[string]*ConfigMapsPerNamespaceDetails {
	cmnsd := map[string]*ConfigMapsPerNamespaceDetails{}

	for _, ns := range namespaces {
		cmnsd[ns] = GetAllConfigMapByNamespace(&ConfigMapClient{Clientset: clientset}, ns)
	}

	return cmnsd
}
