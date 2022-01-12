package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

//go:generate controller-gen object paths=$GOFILE

type KafkaTopicSpec struct {
	Name       string             `json:"name,omitempty"`
	Partitions int32              `json:"partitions,omitempty"`
	Replicas   int16              `json:"replicas,omitempty"`
	Config     map[string]*string `json:"config,omitempty"`
}

type KafkaTopicStatus struct {
	Name       string             `json:"name,omitempty"`
	Partitions int32              `json:"partitions,omitempty"`
	Replicas   int16              `json:"replicas,omitempty"`
	Error      string             `json:"error,omitempty"`
	Config     map[string]*string `json:"config,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KafkaTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KafkaTopicSpec   `json:"spec,omitempty"`
	Status KafkaTopicStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KafkaTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaTopic `json:"items"`
}
