package controller

import (
	"fmt"
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	mementoClient "github.com/gugahoi/memento/pkg/client/clientset/versioned"
	mementoScheme "github.com/gugahoi/memento/pkg/client/clientset/versioned/scheme"
	mementoInformerFactory "github.com/gugahoi/memento/pkg/client/informers/externalversions"
	mementoLister "github.com/gugahoi/memento/pkg/client/listers/memento/v1alpha1"
)

// Controller is the controller for Memento
type Controller struct {
	client *kubernetes.Clientset
	lister mementoLister.RegistryLister
	synced cache.InformerSynced
	queue  workqueue.RateLimitingInterface
}

// New creates a new controller
func New(
	client *kubernetes.Clientset,
	registryClient mementoClient.Interface,
) (*Controller, error) {

	// dbInformerFactory acts like a cache for db resources like above
	registryInformer := mementoInformerFactory.NewSharedInformerFactory(registryClient, 10*time.Minute)
	// start go routines with our informers
	go registryInformer.Start(nil)

	informer := registryInformer.Memento().V1alpha1().Registries()
	lister := informer.Lister()

	// register CRD into api
	mementoScheme.AddToScheme(scheme.Scheme)

	c := &Controller{
		client: client,
		lister: lister,
		synced: informer.Informer().HasSynced,
		queue:  workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "MementoRegistry"),
	}

	log.Info("setting up event handlers")
	informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: c.enqueue,
		UpdateFunc: func(old, new interface{}) {
			if !reflect.DeepEqual(old, new) {
				c.enqueue(new)
			}
		},
		DeleteFunc: c.enqueue,
	})

	return c, nil
}

// Run starts the controller
func (c *Controller) Run(threadiness int, stopChan <-chan struct{}) error {
	// do not allow panics to crash the controller
	defer runtime.HandleCrash()

	// shutdown the queue when done
	defer c.queue.ShutDown()

	log.Info("running Memento Controller")

	log.Info("waiting for cache to sync")
	if !cache.WaitForCacheSync(stopChan, c.synced) {
		return fmt.Errorf("timeout waiting for sync")
	}
	log.Info("caches synced successfully")

	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopChan)
	}

	// block until we are told to exit
	<-stopChan
	return nil
}

func (c *Controller) runWorker() {
	// process the next item in queue until it is empty
	for c.processNextWorkItem() {
	}
}

func (c *Controller) processNextWorkItem() bool {
	// get next item from work queue
	key, quit := c.queue.Get()
	if quit {
		return false
	}

	// indicate to queue when work is finished on a specific item
	defer c.queue.Done(key)

	//  Sync is to push changes for a postgresdb resource
	err := c.pgmgr.Sync(key.(string))
	if err == nil {
	// processed successfully, lets forget item in queue and return success
	log.Infof("processing item %s", key)
	c.queue.Forget(key)
	return true
	// }

	// There was an error processing the item, log and requeue
	runtime.HandleError(err)

	// Add item back in with a rate limited backoff
	c.queue.AddRateLimited(key)

	return true
}

func (c *Controller) enqueue(obj interface{}) {
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(fmt.Errorf("error obtaining key for enqueued object: %v", err))
	}
	log.Infof("enqueueing: %s", key)
	c.queue.Add(key)
}
