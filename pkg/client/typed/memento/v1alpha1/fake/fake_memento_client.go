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
	v1alpha1 "github.com/gugahoi/memento/pkg/client/typed/memento/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeMementoV1alpha1 struct {
	*testing.Fake
}

func (c *FakeMementoV1alpha1) ECRs(namespace string) v1alpha1.ECRInterface {
	return &FakeECRs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeMementoV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
