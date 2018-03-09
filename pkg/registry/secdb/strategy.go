package secdb

import (
	"fmt"
	"github.com/sanjid133/ksd/apis/sho"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
)

func NewStrategy(typer runtime.ObjectTyper) secdbStrategy {
	fmt.Println("NewStrategy")
	return secdbStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, the presence of Initializers if any
// and error in case the given runtime.Object is not a Flunder
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	fmt.Println("GetAttrs")
	apiserver, ok := obj.(*sho.Ksd)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not a SecDb")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), apiserver.Initializers != nil, nil
}

func MatchSecDb(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *sho.Ksd) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type secdbStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (secdbStrategy) NamespaceScoped() bool {
	return true
}
func (secdbStrategy) PrepareForCreate(ctx genericapirequest.Context, obj runtime.Object) {
	fmt.Println("PrepareForCreate")
}

func (secdbStrategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
	fmt.Println("PrepareForUpdate")
}

func (secdbStrategy) Validate(ctx genericapirequest.Context, obj runtime.Object) field.ErrorList {
	fmt.Println("Validate")
	return field.ErrorList{}
}

func (secdbStrategy) AllowCreateOnUpdate() bool {
	fmt.Println("AllowCreateOnUpdate")
	return false
}

func (secdbStrategy) AllowUnconditionalUpdate() bool {
	fmt.Println("AllowUnconditionalUpdate")
	return false
}

func (secdbStrategy) Canonicalize(obj runtime.Object) {
	fmt.Println("Canonicalize")
}

func (secdbStrategy) ValidateUpdate(ctx genericapirequest.Context, obj, old runtime.Object) field.ErrorList {
	fmt.Println("Canonicalize")
	return field.ErrorList{}
}
