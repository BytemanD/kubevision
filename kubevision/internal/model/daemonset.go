package model

import (
	appv1 "k8s.io/api/apps/v1"
)

type Daemonset struct {
	BaseModel
	Labels                 map[string]string `json:"labels,omitempty"`
	NumberReady            int32             `json:"number_ready,omitempty"`
	CurrentNumberScheduled int32             `json:"current_number_scheduled,omitempty"`
	DesiredNumberScheduled int32             `json:"desired_number_scheduled,omitempty"`
	NodeSelector           map[string]string `json:"node_selector,omitempty"`
	MatchLabels            map[string]string `json:"selector,omitempty"`
	Containers             []Container       `json:"containers,omitempty"`
	InitContainers         []Container       `json:"init_containers,omitempty"`
}

func ParseV1Daemonset(item appv1.DaemonSet) Daemonset {
	containers := []Container{}
	initContainers := []Container{}
	for _, c := range item.Spec.Template.Spec.Containers {
		containers = append(containers, ParseV1Container(c, nil))
	}

	for _, c := range item.Spec.Template.Spec.InitContainers {
		initContainers = append(containers, ParseV1Container(c, nil))
	}
	return Daemonset{
		BaseModel: BaseModel{
			Name:     item.Name,
			Creation: item.CreationTimestamp.Format("2006-01-02 15:04:05"),
		},
		Labels:                 item.Labels,
		NumberReady:            item.Status.NumberReady,
		CurrentNumberScheduled: item.Status.CurrentNumberScheduled,
		DesiredNumberScheduled: item.Status.DesiredNumberScheduled,

		NodeSelector:   item.Spec.Template.Spec.NodeSelector,
		MatchLabels:    item.Spec.Selector.MatchLabels,
		Containers:     containers,
		InitContainers: initContainers,
	}
}
