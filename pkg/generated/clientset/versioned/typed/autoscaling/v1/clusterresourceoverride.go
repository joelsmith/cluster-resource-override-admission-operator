/*
Copyright 2024 Red Hat, Inc.

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

package v1

import (
	"context"

	v1 "github.com/openshift/cluster-resource-override-admission-operator/pkg/apis/autoscaling/v1"
	scheme "github.com/openshift/cluster-resource-override-admission-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ClusterResourceOverridesGetter has a method to return a ClusterResourceOverrideInterface.
// A group's client should implement this interface.
type ClusterResourceOverridesGetter interface {
	ClusterResourceOverrides() ClusterResourceOverrideInterface
}

// ClusterResourceOverrideInterface has methods to work with ClusterResourceOverride resources.
type ClusterResourceOverrideInterface interface {
	Create(ctx context.Context, clusterResourceOverride *v1.ClusterResourceOverride, opts metav1.CreateOptions) (*v1.ClusterResourceOverride, error)
	Update(ctx context.Context, clusterResourceOverride *v1.ClusterResourceOverride, opts metav1.UpdateOptions) (*v1.ClusterResourceOverride, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, clusterResourceOverride *v1.ClusterResourceOverride, opts metav1.UpdateOptions) (*v1.ClusterResourceOverride, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ClusterResourceOverride, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ClusterResourceOverrideList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterResourceOverride, err error)
	ClusterResourceOverrideExpansion
}

// clusterResourceOverrides implements ClusterResourceOverrideInterface
type clusterResourceOverrides struct {
	*gentype.ClientWithList[*v1.ClusterResourceOverride, *v1.ClusterResourceOverrideList]
}

// newClusterResourceOverrides returns a ClusterResourceOverrides
func newClusterResourceOverrides(c *AutoscalingV1Client) *clusterResourceOverrides {
	return &clusterResourceOverrides{
		gentype.NewClientWithList[*v1.ClusterResourceOverride, *v1.ClusterResourceOverrideList](
			"clusterresourceoverrides",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.ClusterResourceOverride { return &v1.ClusterResourceOverride{} },
			func() *v1.ClusterResourceOverrideList { return &v1.ClusterResourceOverrideList{} }),
	}
}
