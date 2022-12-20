//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"context"

	"github.com/kcp-dev/logicalcluster/v3"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	wildwestv1alpha1 "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/apis/wildwest/v1alpha1"
	kcpwildwestv1alpha1 "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/clientset/versioned/cluster/typed/wildwest/v1alpha1"
	wildwestv1alpha1client "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/clientset/versioned/typed/wildwest/v1alpha1"
)

var cowboysResource = schema.GroupVersionResource{Group: "wildwest.dev", Version: "v1alpha1", Resource: "cowboys"}
var cowboysKind = schema.GroupVersionKind{Group: "wildwest.dev", Version: "v1alpha1", Kind: "Cowboy"}

type cowboysClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *cowboysClusterClient) Cluster(clusterPath logicalcluster.Path) kcpwildwestv1alpha1.CowboysNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &cowboysNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of Cowboys that match those selectors across all clusters.
func (c *cowboysClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*wildwestv1alpha1.CowboyList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(cowboysResource, cowboysKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &wildwestv1alpha1.CowboyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &wildwestv1alpha1.CowboyList{ListMeta: obj.(*wildwestv1alpha1.CowboyList).ListMeta}
	for _, item := range obj.(*wildwestv1alpha1.CowboyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested Cowboys across all clusters.
func (c *cowboysClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(cowboysResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type cowboysNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *cowboysNamespacer) Namespace(namespace string) wildwestv1alpha1client.CowboyInterface {
	return &cowboysClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type cowboysClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *cowboysClient) Create(ctx context.Context, cowboy *wildwestv1alpha1.Cowboy, opts metav1.CreateOptions) (*wildwestv1alpha1.Cowboy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(cowboysResource, c.ClusterPath, c.Namespace, cowboy), &wildwestv1alpha1.Cowboy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*wildwestv1alpha1.Cowboy), err
}

func (c *cowboysClient) Update(ctx context.Context, cowboy *wildwestv1alpha1.Cowboy, opts metav1.UpdateOptions) (*wildwestv1alpha1.Cowboy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(cowboysResource, c.ClusterPath, c.Namespace, cowboy), &wildwestv1alpha1.Cowboy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*wildwestv1alpha1.Cowboy), err
}

func (c *cowboysClient) UpdateStatus(ctx context.Context, cowboy *wildwestv1alpha1.Cowboy, opts metav1.UpdateOptions) (*wildwestv1alpha1.Cowboy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(cowboysResource, c.ClusterPath, "status", c.Namespace, cowboy), &wildwestv1alpha1.Cowboy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*wildwestv1alpha1.Cowboy), err
}

func (c *cowboysClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(cowboysResource, c.ClusterPath, c.Namespace, name, opts), &wildwestv1alpha1.Cowboy{})
	return err
}

func (c *cowboysClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(cowboysResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &wildwestv1alpha1.CowboyList{})
	return err
}

func (c *cowboysClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*wildwestv1alpha1.Cowboy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(cowboysResource, c.ClusterPath, c.Namespace, name), &wildwestv1alpha1.Cowboy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*wildwestv1alpha1.Cowboy), err
}

// List takes label and field selectors, and returns the list of Cowboys that match those selectors.
func (c *cowboysClient) List(ctx context.Context, opts metav1.ListOptions) (*wildwestv1alpha1.CowboyList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(cowboysResource, cowboysKind, c.ClusterPath, c.Namespace, opts), &wildwestv1alpha1.CowboyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &wildwestv1alpha1.CowboyList{ListMeta: obj.(*wildwestv1alpha1.CowboyList).ListMeta}
	for _, item := range obj.(*wildwestv1alpha1.CowboyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *cowboysClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(cowboysResource, c.ClusterPath, c.Namespace, opts))
}

func (c *cowboysClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*wildwestv1alpha1.Cowboy, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(cowboysResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &wildwestv1alpha1.Cowboy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*wildwestv1alpha1.Cowboy), err
}
