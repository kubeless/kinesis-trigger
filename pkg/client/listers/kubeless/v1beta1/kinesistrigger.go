/*
Copyright (c) 2016-2017 Bitnami

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

// This file was automatically generated by lister-gen

package v1beta1

import (
	v1beta1 "github.com/kubeless/kinesis-trigger/pkg/apis/kubeless/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KinesisTriggerLister helps list KinesisTriggers.
type KinesisTriggerLister interface {
	// List lists all KinesisTriggers in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.KinesisTrigger, err error)
	// KinesisTriggers returns an object that can list and get KinesisTriggers.
	KinesisTriggers(namespace string) KinesisTriggerNamespaceLister
	KinesisTriggerListerExpansion
}

// kinesisTriggerLister implements the KinesisTriggerLister interface.
type kinesisTriggerLister struct {
	indexer cache.Indexer
}

// NewKinesisTriggerLister returns a new KinesisTriggerLister.
func NewKinesisTriggerLister(indexer cache.Indexer) KinesisTriggerLister {
	return &kinesisTriggerLister{indexer: indexer}
}

// List lists all KinesisTriggers in the indexer.
func (s *kinesisTriggerLister) List(selector labels.Selector) (ret []*v1beta1.KinesisTrigger, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.KinesisTrigger))
	})
	return ret, err
}

// KinesisTriggers returns an object that can list and get KinesisTriggers.
func (s *kinesisTriggerLister) KinesisTriggers(namespace string) KinesisTriggerNamespaceLister {
	return kinesisTriggerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KinesisTriggerNamespaceLister helps list and get KinesisTriggers.
type KinesisTriggerNamespaceLister interface {
	// List lists all KinesisTriggers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.KinesisTrigger, err error)
	// Get retrieves the KinesisTrigger from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.KinesisTrigger, error)
	KinesisTriggerNamespaceListerExpansion
}

// kinesisTriggerNamespaceLister implements the KinesisTriggerNamespaceLister
// interface.
type kinesisTriggerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KinesisTriggers in the indexer for a given namespace.
func (s kinesisTriggerNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.KinesisTrigger, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.KinesisTrigger))
	})
	return ret, err
}

// Get retrieves the KinesisTrigger from the indexer for a given namespace and name.
func (s kinesisTriggerNamespaceLister) Get(name string) (*v1beta1.KinesisTrigger, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("kinesistrigger"), name)
	}
	return obj.(*v1beta1.KinesisTrigger), nil
}
