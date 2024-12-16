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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	autoscalingv1 "github.com/openshift/cluster-resource-override-admission-operator/pkg/apis/autoscaling/v1"
	versioned "github.com/openshift/cluster-resource-override-admission-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/openshift/cluster-resource-override-admission-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/cluster-resource-override-admission-operator/pkg/generated/listers/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterResourceOverrideInformer provides access to a shared informer and lister for
// ClusterResourceOverrides.
type ClusterResourceOverrideInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterResourceOverrideLister
}

type clusterResourceOverrideInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterResourceOverrideInformer constructs a new informer for ClusterResourceOverride type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterResourceOverrideInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterResourceOverrideInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterResourceOverrideInformer constructs a new informer for ClusterResourceOverride type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterResourceOverrideInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1().ClusterResourceOverrides().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1().ClusterResourceOverrides().Watch(context.TODO(), options)
			},
		},
		&autoscalingv1.ClusterResourceOverride{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterResourceOverrideInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterResourceOverrideInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterResourceOverrideInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&autoscalingv1.ClusterResourceOverride{}, f.defaultInformer)
}

func (f *clusterResourceOverrideInformer) Lister() v1.ClusterResourceOverrideLister {
	return v1.NewClusterResourceOverrideLister(f.Informer().GetIndexer())
}
