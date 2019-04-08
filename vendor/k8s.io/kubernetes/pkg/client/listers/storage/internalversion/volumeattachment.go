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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	storage "k8s.io/kubernetes/pkg/apis/storage"
)

// VolumeAttachmentLister helps list VolumeAttachments.
type VolumeAttachmentLister interface {
	// List lists all VolumeAttachments in the indexer.
	List(selector labels.Selector) (ret []*storage.VolumeAttachment, err error)
	// Get retrieves the VolumeAttachment from the index for a given name.
	Get(name string) (*storage.VolumeAttachment, error)
	VolumeAttachmentListerExpansion
}

// volumeAttachmentLister implements the VolumeAttachmentLister interface.
type volumeAttachmentLister struct {
	indexer cache.Indexer
}

// NewVolumeAttachmentLister returns a new VolumeAttachmentLister.
func NewVolumeAttachmentLister(indexer cache.Indexer) VolumeAttachmentLister {
	return &volumeAttachmentLister{indexer: indexer}
}

// List lists all VolumeAttachments in the indexer.
func (s *volumeAttachmentLister) List(selector labels.Selector) (ret []*storage.VolumeAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*storage.VolumeAttachment))
	})
	return ret, err
}

// Get retrieves the VolumeAttachment from the index for a given name.
func (s *volumeAttachmentLister) Get(name string) (*storage.VolumeAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(storage.Resource("volumeattachment"), name)
	}
	return obj.(*storage.VolumeAttachment), nil
}
