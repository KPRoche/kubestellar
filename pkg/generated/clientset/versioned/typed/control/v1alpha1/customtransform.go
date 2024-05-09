/*
Copyright The KubeStellar Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/kubestellar/kubestellar/api/control/v1alpha1"
	scheme "github.com/kubestellar/kubestellar/pkg/generated/clientset/versioned/scheme"
)

// CustomTransformsGetter has a method to return a CustomTransformInterface.
// A group's client should implement this interface.
type CustomTransformsGetter interface {
	CustomTransforms() CustomTransformInterface
}

// CustomTransformInterface has methods to work with CustomTransform resources.
type CustomTransformInterface interface {
	Create(ctx context.Context, customTransform *v1alpha1.CustomTransform, opts v1.CreateOptions) (*v1alpha1.CustomTransform, error)
	Update(ctx context.Context, customTransform *v1alpha1.CustomTransform, opts v1.UpdateOptions) (*v1alpha1.CustomTransform, error)
	UpdateStatus(ctx context.Context, customTransform *v1alpha1.CustomTransform, opts v1.UpdateOptions) (*v1alpha1.CustomTransform, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CustomTransform, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CustomTransformList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CustomTransform, err error)
	CustomTransformExpansion
}

// customTransforms implements CustomTransformInterface
type customTransforms struct {
	client rest.Interface
}

// newCustomTransforms returns a CustomTransforms
func newCustomTransforms(c *ControlV1alpha1Client) *customTransforms {
	return &customTransforms{
		client: c.RESTClient(),
	}
}

// Get takes name of the customTransform, and returns the corresponding customTransform object, and an error if there is any.
func (c *customTransforms) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CustomTransform, err error) {
	result = &v1alpha1.CustomTransform{}
	err = c.client.Get().
		Resource("customtransforms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CustomTransforms that match those selectors.
func (c *customTransforms) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CustomTransformList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CustomTransformList{}
	err = c.client.Get().
		Resource("customtransforms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested customTransforms.
func (c *customTransforms) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("customtransforms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a customTransform and creates it.  Returns the server's representation of the customTransform, and an error, if there is any.
func (c *customTransforms) Create(ctx context.Context, customTransform *v1alpha1.CustomTransform, opts v1.CreateOptions) (result *v1alpha1.CustomTransform, err error) {
	result = &v1alpha1.CustomTransform{}
	err = c.client.Post().
		Resource("customtransforms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(customTransform).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a customTransform and updates it. Returns the server's representation of the customTransform, and an error, if there is any.
func (c *customTransforms) Update(ctx context.Context, customTransform *v1alpha1.CustomTransform, opts v1.UpdateOptions) (result *v1alpha1.CustomTransform, err error) {
	result = &v1alpha1.CustomTransform{}
	err = c.client.Put().
		Resource("customtransforms").
		Name(customTransform.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(customTransform).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *customTransforms) UpdateStatus(ctx context.Context, customTransform *v1alpha1.CustomTransform, opts v1.UpdateOptions) (result *v1alpha1.CustomTransform, err error) {
	result = &v1alpha1.CustomTransform{}
	err = c.client.Put().
		Resource("customtransforms").
		Name(customTransform.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(customTransform).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the customTransform and deletes it. Returns an error if one occurs.
func (c *customTransforms) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("customtransforms").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *customTransforms) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("customtransforms").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched customTransform.
func (c *customTransforms) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CustomTransform, err error) {
	result = &v1alpha1.CustomTransform{}
	err = c.client.Patch(pt).
		Resource("customtransforms").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}