package v1

import (
	"context"

	v1 "github.com/aanogueira/kubernetes-crds/api/types/kafka/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type KafkaTopicInterface interface {
	List(opts metav1.ListOptions) (*v1.KafkaTopicList, error)
	Get(name string, options metav1.GetOptions) (*v1.KafkaTopic, error)
	Create(*v1.KafkaTopic) (*v1.KafkaTopic, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
}

type kafkaTopicClient struct {
	restClient rest.Interface
	ns         string
}

func (c *kafkaTopicClient) List(opts metav1.ListOptions) (*v1.KafkaTopicList, error) {
	result := v1.KafkaTopicList{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("kafkatopics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *kafkaTopicClient) Get(name string, opts metav1.GetOptions) (*v1.KafkaTopic, error) {
	result := v1.KafkaTopic{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource("kafkatopics").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *kafkaTopicClient) Create(project *v1.KafkaTopic) (*v1.KafkaTopic, error) {
	result := v1.KafkaTopic{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource("kafkatopics").
		Body(project).
		Do(context.Background()).
		Into(&result)

	return &result, err
}

func (c *kafkaTopicClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.ns).
		Resource("kafkatopics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.Background())
}

func (c *kafkaTopicClient) Delete(name string, options *metav1.DeleteOptions) error {
	err := c.restClient.
		Delete().
		Namespace(c.ns).
		Resource("kafkatopics").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()

	return err
}

func (c *kafkaTopicClient) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	err := c.restClient.
		Delete().
		Namespace(c.ns).
		Resource("kafkatopics").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do(context.Background()).
		Error()

	return err
}
