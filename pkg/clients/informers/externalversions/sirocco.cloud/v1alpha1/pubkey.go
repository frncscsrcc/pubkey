/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	versioned "pubkey/pkg/clients/clientset/versioned"
	internalinterfaces "pubkey/pkg/clients/informers/externalversions/internalinterfaces"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	siroccocloudv1alpha1 "pubkey/pkg/apis/sirocco.cloud/v1alpha1"
	v1alpha1 "pubkey/pkg/clients/listers/sirocco.cloud/v1alpha1"
)

// PubkeyInformer provides access to a shared informer and lister for
// Pubkeys.
type PubkeyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PubkeyLister
}

type pubkeyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPubkeyInformer constructs a new informer for Pubkey type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPubkeyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPubkeyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPubkeyInformer constructs a new informer for Pubkey type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPubkeyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SiroccoV1alpha1().Pubkeys(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SiroccoV1alpha1().Pubkeys(namespace).Watch(options)
			},
		},
		&siroccocloudv1alpha1.Pubkey{},
		resyncPeriod,
		indexers,
	)
}

func (f *pubkeyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPubkeyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pubkeyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&siroccocloudv1alpha1.Pubkey{}, f.defaultInformer)
}

func (f *pubkeyInformer) Lister() v1alpha1.PubkeyLister {
	return v1alpha1.NewPubkeyLister(f.Informer().GetIndexer())
}
