package operator

import (
	"k8s.io/apimachinery/pkg/labels"
)


func (o *Operator) registerPubkeys() error {
	pubkeysInformer := o.pubkeyInformerFactory.Sirocco().V1alpha1().Pubkeys()
	pubkeys, err := pubkeysInformer.Lister().Pubkeys(o.namespace).List(labels.Everything())
	if err != nil {
		return err
	}
	
	for _, pubkey := range pubkeys {
		Show(pubkey)
	}

	return nil
}

func (o *Operator) addedPubkeyHandler(new interface{}) {

}

func (o *Operator) updatedPubkeyHandler(old, new interface{}) {

}

func (o *Operator) deletedPubkeyHandler(old interface{}) {

}