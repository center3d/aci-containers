/*
Copyright 2019 The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CustomResourceDefinitionLister helps list CustomResourceDefinitions.
type CustomResourceDefinitionLister interface {
	// List lists all CustomResourceDefinitions in the indexer.
	List(selector labels.Selector) (ret []*apiextensions.CustomResourceDefinition, err error)
	// Get retrieves the CustomResourceDefinition from the index for a given name.
	Get(name string) (*apiextensions.CustomResourceDefinition, error)
	CustomResourceDefinitionListerExpansion
}

// customResourceDefinitionLister implements the CustomResourceDefinitionLister interface.
type customResourceDefinitionLister struct {
	indexer cache.Indexer
}

// NewCustomResourceDefinitionLister returns a new CustomResourceDefinitionLister.
func NewCustomResourceDefinitionLister(indexer cache.Indexer) CustomResourceDefinitionLister {
	return &customResourceDefinitionLister{indexer: indexer}
}

// List lists all CustomResourceDefinitions in the indexer.
func (s *customResourceDefinitionLister) List(selector labels.Selector) (ret []*apiextensions.CustomResourceDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*apiextensions.CustomResourceDefinition))
	})
	return ret, err
}

// Get retrieves the CustomResourceDefinition from the index for a given name.
func (s *customResourceDefinitionLister) Get(name string) (*apiextensions.CustomResourceDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(apiextensions.Resource("customresourcedefinition"), name)
	}
	return obj.(*apiextensions.CustomResourceDefinition), nil
}
