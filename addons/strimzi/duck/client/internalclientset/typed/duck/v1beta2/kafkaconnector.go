/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta2

import (
	"context"
	"time"

	scheme "github.com/lendi-au/camel-k/addons/strimzi/duck/client/internalclientset/scheme"
	v1beta2 "github.com/lendi-au/camel-k/addons/strimzi/duck/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KafkaConnectorsGetter has a method to return a KafkaConnectorInterface.
// A group's client should implement this interface.
type KafkaConnectorsGetter interface {
	KafkaConnectors(namespace string) KafkaConnectorInterface
}

// KafkaConnectorInterface has methods to work with KafkaConnector resources.
type KafkaConnectorInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta2.KafkaConnector, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta2.KafkaConnectorList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	KafkaConnectorExpansion
}

// kafkaConnectors implements KafkaConnectorInterface
type kafkaConnectors struct {
	client rest.Interface
	ns     string
}

// newKafkaConnectors returns a KafkaConnectors
func newKafkaConnectors(c *KafkaV1beta2Client, namespace string) *kafkaConnectors {
	return &kafkaConnectors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kafkaConnector, and returns the corresponding kafkaConnector object, and an error if there is any.
func (c *kafkaConnectors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.KafkaConnector, err error) {
	result = &v1beta2.KafkaConnector{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkaconnectors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KafkaConnectors that match those selectors.
func (c *kafkaConnectors) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.KafkaConnectorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta2.KafkaConnectorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kafkaconnectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kafkaConnectors.
func (c *kafkaConnectors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kafkaconnectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}
