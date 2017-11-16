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
	v1alpha1 "github.com/gugahoi/memento/pkg/apis/ecr/v1alpha1"
	scheme "github.com/gugahoi/memento/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ECRsGetter has a method to return a ECRInterface.
// A group's client should implement this interface.
type ECRsGetter interface {
	ECRs(namespace string) ECRInterface
}

// ECRInterface has methods to work with ECR resources.
type ECRInterface interface {
	Create(*v1alpha1.ECR) (*v1alpha1.ECR, error)
	Update(*v1alpha1.ECR) (*v1alpha1.ECR, error)
	UpdateStatus(*v1alpha1.ECR) (*v1alpha1.ECR, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ECR, error)
	List(opts v1.ListOptions) (*v1alpha1.ECRList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ECR, err error)
	ECRExpansion
}

// eCRs implements ECRInterface
type eCRs struct {
	client rest.Interface
	ns     string
}

// newECRs returns a ECRs
func newECRs(c *MementoV1alpha1Client, namespace string) *eCRs {
	return &eCRs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eCR, and returns the corresponding eCR object, and an error if there is any.
func (c *eCRs) Get(name string, options v1.GetOptions) (result *v1alpha1.ECR, err error) {
	result = &v1alpha1.ECR{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ecrs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ECRs that match those selectors.
func (c *eCRs) List(opts v1.ListOptions) (result *v1alpha1.ECRList, err error) {
	result = &v1alpha1.ECRList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ecrs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eCRs.
func (c *eCRs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ecrs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a eCR and creates it.  Returns the server's representation of the eCR, and an error, if there is any.
func (c *eCRs) Create(eCR *v1alpha1.ECR) (result *v1alpha1.ECR, err error) {
	result = &v1alpha1.ECR{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ecrs").
		Body(eCR).
		Do().
		Into(result)
	return
}

// Update takes the representation of a eCR and updates it. Returns the server's representation of the eCR, and an error, if there is any.
func (c *eCRs) Update(eCR *v1alpha1.ECR) (result *v1alpha1.ECR, err error) {
	result = &v1alpha1.ECR{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ecrs").
		Name(eCR.Name).
		Body(eCR).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *eCRs) UpdateStatus(eCR *v1alpha1.ECR) (result *v1alpha1.ECR, err error) {
	result = &v1alpha1.ECR{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ecrs").
		Name(eCR.Name).
		SubResource("status").
		Body(eCR).
		Do().
		Into(result)
	return
}

// Delete takes name of the eCR and deletes it. Returns an error if one occurs.
func (c *eCRs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ecrs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eCRs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ecrs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched eCR.
func (c *eCRs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ECR, err error) {
	result = &v1alpha1.ECR{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ecrs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
