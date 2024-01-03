// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/GoogleCloudPlatform/prometheus-engine/pkg/operator/apis/monitoring/v1"
	scheme "github.com/GoogleCloudPlatform/prometheus-engine/pkg/operator/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OperatorConfigsGetter has a method to return a OperatorConfigInterface.
// A group's client should implement this interface.
type OperatorConfigsGetter interface {
	OperatorConfigs(namespace string) OperatorConfigInterface
}

// OperatorConfigInterface has methods to work with OperatorConfig resources.
type OperatorConfigInterface interface {
	Create(ctx context.Context, operatorConfig *v1.OperatorConfig, opts metav1.CreateOptions) (*v1.OperatorConfig, error)
	Update(ctx context.Context, operatorConfig *v1.OperatorConfig, opts metav1.UpdateOptions) (*v1.OperatorConfig, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.OperatorConfig, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.OperatorConfigList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OperatorConfig, err error)
	OperatorConfigExpansion
}

// operatorConfigs implements OperatorConfigInterface
type operatorConfigs struct {
	client rest.Interface
	ns     string
}

// newOperatorConfigs returns a OperatorConfigs
func newOperatorConfigs(c *MonitoringV1Client, namespace string) *operatorConfigs {
	return &operatorConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the operatorConfig, and returns the corresponding operatorConfig object, and an error if there is any.
func (c *operatorConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.OperatorConfig, err error) {
	result = &v1.OperatorConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("operatorconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OperatorConfigs that match those selectors.
func (c *operatorConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.OperatorConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.OperatorConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("operatorconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested operatorConfigs.
func (c *operatorConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("operatorconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a operatorConfig and creates it.  Returns the server's representation of the operatorConfig, and an error, if there is any.
func (c *operatorConfigs) Create(ctx context.Context, operatorConfig *v1.OperatorConfig, opts metav1.CreateOptions) (result *v1.OperatorConfig, err error) {
	result = &v1.OperatorConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("operatorconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operatorConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a operatorConfig and updates it. Returns the server's representation of the operatorConfig, and an error, if there is any.
func (c *operatorConfigs) Update(ctx context.Context, operatorConfig *v1.OperatorConfig, opts metav1.UpdateOptions) (result *v1.OperatorConfig, err error) {
	result = &v1.OperatorConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("operatorconfigs").
		Name(operatorConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operatorConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the operatorConfig and deletes it. Returns an error if one occurs.
func (c *operatorConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("operatorconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *operatorConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("operatorconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched operatorConfig.
func (c *operatorConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OperatorConfig, err error) {
	result = &v1.OperatorConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("operatorconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
