/*
Copyright 2017 Gustavo Hoirisch.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1alpha1

import (
	v1alpha1 "github.com/gugahoi/memento/pkg/apis/registry/v1alpha1"
	scheme "github.com/gugahoi/memento/pkg/client/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RegistriesGetter has a method to return a RegistryInterface.
// A group's client should implement this interface.
type RegistriesGetter interface {
	Registries(namespace string) RegistryInterface
}

// RegistryInterface has methods to work with Registry resources.
type RegistryInterface interface {
	Create(*v1alpha1.Registry) (*v1alpha1.Registry, error)
	Update(*v1alpha1.Registry) (*v1alpha1.Registry, error)
	UpdateStatus(*v1alpha1.Registry) (*v1alpha1.Registry, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Registry, error)
	List(opts v1.ListOptions) (*v1alpha1.RegistryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Registry, err error)
	RegistryExpansion
}

// registries implements RegistryInterface
type registries struct {
	client rest.Interface
	ns     string
}

// newRegistries returns a Registries
func newRegistries(c *MementoV1alpha1Client, namespace string) *registries {
	return &registries{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the registry, and returns the corresponding registry object, and an error if there is any.
func (c *registries) Get(name string, options v1.GetOptions) (result *v1alpha1.Registry, err error) {
	result = &v1alpha1.Registry{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("registries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Registries that match those selectors.
func (c *registries) List(opts v1.ListOptions) (result *v1alpha1.RegistryList, err error) {
	result = &v1alpha1.RegistryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("registries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested registries.
func (c *registries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("registries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a registry and creates it.  Returns the server's representation of the registry, and an error, if there is any.
func (c *registries) Create(registry *v1alpha1.Registry) (result *v1alpha1.Registry, err error) {
	result = &v1alpha1.Registry{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("registries").
		Body(registry).
		Do().
		Into(result)
	return
}

// Update takes the representation of a registry and updates it. Returns the server's representation of the registry, and an error, if there is any.
func (c *registries) Update(registry *v1alpha1.Registry) (result *v1alpha1.Registry, err error) {
	result = &v1alpha1.Registry{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("registries").
		Name(registry.Name).
		Body(registry).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *registries) UpdateStatus(registry *v1alpha1.Registry) (result *v1alpha1.Registry, err error) {
	result = &v1alpha1.Registry{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("registries").
		Name(registry.Name).
		SubResource("status").
		Body(registry).
		Do().
		Into(result)
	return
}

// Delete takes name of the registry and deletes it. Returns an error if one occurs.
func (c *registries) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("registries").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *registries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("registries").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched registry.
func (c *registries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Registry, err error) {
	result = &v1alpha1.Registry{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("registries").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
