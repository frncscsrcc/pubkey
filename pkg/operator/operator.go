package operator

import (
	"fmt"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
	pubkeyV1alpha1 "pubkey/pkg/clients/clientset/versioned"
	pubkeyInformers "pubkey/pkg/clients/informers/externalversions"
	"time"
)

type Operator struct {
	namespace             string
	coreClientSet         *kubernetes.Clientset
	coreInformerFactory   informers.SharedInformerFactory
	pubkeyClientSet       *pubkeyV1alpha1.Clientset
	pubkeyInformerFactory pubkeyInformers.SharedInformerFactory

	//	podToObserve map[string]bool

	workqueue workqueue.DelayingInterface

	initialized bool
	done        chan struct{}
	log         Log
	config      Config
}

func New(appConfig Config) (*Operator, error) {
	o := &Operator{}

	config, err := clientcmd.BuildConfigFromFlags("", appConfig.KubeConfig)
	if err != nil {
		panic(err)
	}

	// Keep a reference of the kubernetes core (pods, ...) client set
	coreClientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return o, err
	}

	// Initialize a shared informer factory for the core objects (eg: pods)
	coreInformerFactory := informers.NewSharedInformerFactory(coreClientSet, time.Second*30)

	// Keep a reference of the pubkey specific CRD client set
	pubkeyClientSet, err := pubkeyV1alpha1.NewForConfig(config)
	if err != nil {
		return o, err
	}

	// Initialize a shared informer factory for the pubkey objects (Pubkey)
	pubkeyInformerFactory := pubkeyInformers.NewSharedInformerFactory(pubkeyClientSet, time.Second*30)

	o.namespace = appConfig.Namespace
	o.coreClientSet = coreClientSet
	o.coreInformerFactory = coreInformerFactory
	o.pubkeyClientSet = pubkeyClientSet
	o.pubkeyInformerFactory = pubkeyInformerFactory
	o.workqueue = workqueue.NewDelayingQueue()
	//o.podToObserve = make(map[string]bool)
	o.done = make(chan struct{})
	o.log = NewLog(appConfig.DebugLevel)
	o.config = appConfig

	return o, nil
}

func (o *Operator) Init() {
	// Initialize the pubkey informers
	pubkeysInformer := o.pubkeyInformerFactory.Sirocco().V1alpha1().Pubkeys()
	pubkeysInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    o.addedPubkeyHandler,
		UpdateFunc: o.updatedPubkeyHandler,
		DeleteFunc: o.deletedPubkeyHandler,
	})

	// Activate the informer and wait for the cache
	o.coreInformerFactory.Start(o.done)
	o.coreInformerFactory.WaitForCacheSync(o.done)
	o.pubkeyInformerFactory.Start(o.done)
	o.pubkeyInformerFactory.WaitForCacheSync(o.done)

	// Register the pubkey already present in the cluster
	if err := o.registerPubkeys(); err != nil {
		o.log.Error.Println("Problem in pubkeys initialization. Skip it.")
	}

	// Mark the fact the object is initialized
	o.initialized = true

	o.log.Info.Println("Pubkey operator is initialized")

	// Initialize webserver (in a separate thread)
	go o.initializeWebServer()
}

func (o *Operator) Run() {
	o.log.Info.Println("Pubkey operator is running")

	// Be sure the done channel triggers a shutdown
	// This function is execute in a separate thread
	go func(o *Operator) {
		// Wait done signal
		<-o.done
		o.workqueue.ShutDown()
	}(o)

	for true {
		generic, shutdown := o.workqueue.Get()
		if shutdown {
			break
		}

		var ok bool
		var qi queueItem

		// Cast the item to be a queueItem structure
		qi, ok = generic.(queueItem)
		if !ok {
			o.log.Info.Println("Invalid working queue item. Just ignoring it.")
			continue
		}

		o.log.Trace.Println("Received queue item " + qi.operation + " for " + qi.item)

		// Callback function to call when the queue item is processed
		// itemDone := func() {
		// 	o.workqueue.Done(qi)
		// }

		switch qi.getOperation() {

		default:

		}
	}
}

func Show(i interface{}) {
	fmt.Printf("%+v\n", i)
}
