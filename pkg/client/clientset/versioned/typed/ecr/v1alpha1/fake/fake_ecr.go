/*
Copyright 2017 Gustavo Hoirisch.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package fake

import (
	v1alpha1 "github.com/gugahoi/memento/pkg/apis/ecr/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeECRs implements ECRInterface
type FakeECRs struct {
	Fake *FakeMementoV1alpha1
	ns   string
}

var ecrsResource = schema.GroupVersionResource{Group: "memento.gugahoi.com", Version: "v1alpha1", Resource: "ecrs"}

var ecrsKind = schema.GroupVersionKind{Group: "memento.gugahoi.com", Version: "v1alpha1", Kind: "ECR"}

// Get takes name of the eCR, and returns the corresponding eCR object, and an error if there is any.
func (c *FakeECRs) Get(name string, options v1.GetOptions) (result *v1alpha1.ECR, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ecrsResource, c.ns, name), &v1alpha1.ECR{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ECR), err
}

// List takes label and field selectors, and returns the list of ECRs that match those selectors.
func (c *FakeECRs) List(opts v1.ListOptions) (result *v1alpha1.ECRList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ecrsResource, ecrsKind, c.ns, opts), &v1alpha1.ECRList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ECRList{}
	for _, item := range obj.(*v1alpha1.ECRList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eCRs.
func (c *FakeECRs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ecrsResource, c.ns, opts))

}

// Create takes the representation of a eCR and creates it.  Returns the server's representation of the eCR, and an error, if there is any.
func (c *FakeECRs) Create(eCR *v1alpha1.ECR) (result *v1alpha1.ECR, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ecrsResource, c.ns, eCR), &v1alpha1.ECR{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ECR), err
}

// Update takes the representation of a eCR and updates it. Returns the server's representation of the eCR, and an error, if there is any.
func (c *FakeECRs) Update(eCR *v1alpha1.ECR) (result *v1alpha1.ECR, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ecrsResource, c.ns, eCR), &v1alpha1.ECR{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ECR), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeECRs) UpdateStatus(eCR *v1alpha1.ECR) (*v1alpha1.ECR, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ecrsResource, "status", c.ns, eCR), &v1alpha1.ECR{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ECR), err
}

// Delete takes name of the eCR and deletes it. Returns an error if one occurs.
func (c *FakeECRs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ecrsResource, c.ns, name), &v1alpha1.ECR{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeECRs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ecrsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ECRList{})
	return err
}

// Patch applies the patch and returns the patched eCR.
func (c *FakeECRs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ECR, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ecrsResource, c.ns, name, data, subresources...), &v1alpha1.ECR{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ECR), err
}
