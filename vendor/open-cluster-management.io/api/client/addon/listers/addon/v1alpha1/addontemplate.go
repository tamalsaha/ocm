// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "open-cluster-management.io/api/addon/v1alpha1"
)

// AddOnTemplateLister helps list AddOnTemplates.
// All objects returned here must be treated as read-only.
type AddOnTemplateLister interface {
	// List lists all AddOnTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AddOnTemplate, err error)
	// Get retrieves the AddOnTemplate from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AddOnTemplate, error)
	AddOnTemplateListerExpansion
}

// addOnTemplateLister implements the AddOnTemplateLister interface.
type addOnTemplateLister struct {
	indexer cache.Indexer
}

// NewAddOnTemplateLister returns a new AddOnTemplateLister.
func NewAddOnTemplateLister(indexer cache.Indexer) AddOnTemplateLister {
	return &addOnTemplateLister{indexer: indexer}
}

// List lists all AddOnTemplates in the indexer.
func (s *addOnTemplateLister) List(selector labels.Selector) (ret []*v1alpha1.AddOnTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AddOnTemplate))
	})
	return ret, err
}

// Get retrieves the AddOnTemplate from the index for a given name.
func (s *addOnTemplateLister) Get(name string) (*v1alpha1.AddOnTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("addontemplate"), name)
	}
	return obj.(*v1alpha1.AddOnTemplate), nil
}