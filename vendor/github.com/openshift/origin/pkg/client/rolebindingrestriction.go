package client

import (
	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/watch"

	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
)

type RoleBindingRestrictionsNamespacer interface {
	RoleBindingRestrictions(namespace string) RoleBindingRestrictionInterface
}

type RoleBindingRestrictionInterface interface {
	List(opts kapi.ListOptions) (*authorizationapi.RoleBindingRestrictionList, error)
	Get(name string) (*authorizationapi.RoleBindingRestriction, error)
	Create(roleBindingRestriction *authorizationapi.RoleBindingRestriction) (*authorizationapi.RoleBindingRestriction, error)
	Update(roleBindingRestriction *authorizationapi.RoleBindingRestriction) (*authorizationapi.RoleBindingRestriction, error)
	Delete(name string) error
	Watch(opts kapi.ListOptions) (watch.Interface, error)
}

type roleBindingRestrictions struct {
	r  *Client
	ns string
}

// newRoleBindingRestrictions returns a roleBindingRestrictions
func newRoleBindingRestrictions(c *Client, namespace string) *roleBindingRestrictions {
	return &roleBindingRestrictions{
		r:  c,
		ns: namespace,
	}
}

func (c *roleBindingRestrictions) List(opts kapi.ListOptions) (result *authorizationapi.RoleBindingRestrictionList, err error) {
	result = &authorizationapi.RoleBindingRestrictionList{}
	err = c.r.Get().
		Namespace(c.ns).
		Resource("rolebindingrestrictions").
		VersionedParams(&opts, kapi.ParameterCodec).
		Do().
		Into(result)
	return
}

func (c *roleBindingRestrictions) Get(name string) (result *authorizationapi.RoleBindingRestriction, err error) {
	result = &authorizationapi.RoleBindingRestriction{}
	err = c.r.Get().
		Namespace(c.ns).
		Resource("rolebindingrestrictions").
		Name(name).
		Do().
		Into(result)
	return
}

func (c *roleBindingRestrictions) Create(roleBindingRestriction *authorizationapi.RoleBindingRestriction) (result *authorizationapi.RoleBindingRestriction, err error) {
	result = &authorizationapi.RoleBindingRestriction{}
	err = c.r.Post().
		Namespace(c.ns).
		Resource("rolebindingrestrictions").
		Body(roleBindingRestriction).
		Do().
		Into(result)
	return
}

func (c *roleBindingRestrictions) Update(roleBindingRestriction *authorizationapi.RoleBindingRestriction) (result *authorizationapi.RoleBindingRestriction, err error) {
	result = &authorizationapi.RoleBindingRestriction{}
	err = c.r.Put().
		Namespace(c.ns).
		Resource("rolebindingrestrictions").
		Name(roleBindingRestriction.Name).
		Body(roleBindingRestriction).
		Do().
		Into(result)
	return
}

func (c *roleBindingRestrictions) Delete(name string) (err error) {
	err = c.r.Delete().
		Namespace(c.ns).
		Resource("rolebindingrestrictions").
		Name(name).
		Do().
		Error()
	return
}

func (c *roleBindingRestrictions) Watch(opts kapi.ListOptions) (watch.Interface, error) {
	return c.r.Get().
		Prefix("watch").
		Namespace(c.ns).
		Resource("rolebindingrestrictions").
		VersionedParams(&opts, kapi.ParameterCodec).
		Watch()
}
