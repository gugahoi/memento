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
	"github.com/gugahoi/memento/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type MementoV1alpha1Interface interface {
	RESTClient() rest.Interface
	RegistriesGetter
}

// MementoV1alpha1Client is used to interact with features provided by the memento.gugahoi.com group.
type MementoV1alpha1Client struct {
	restClient rest.Interface
}

func (c *MementoV1alpha1Client) Registries(namespace string) RegistryInterface {
	return newRegistries(c, namespace)
}

// NewForConfig creates a new MementoV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*MementoV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &MementoV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new MementoV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *MementoV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MementoV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *MementoV1alpha1Client {
	return &MementoV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *MementoV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
