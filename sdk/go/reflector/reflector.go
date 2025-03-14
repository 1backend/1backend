package reflector

import (
	"encoding/json"
	"reflect"
)

// Creates a deep copy of a slice of map or structs
// into a struct of 'instances'.
func DeepCopySliceIntoType(result []any, instance any) ([]any, error) {
	bs, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	elemType := reflect.TypeOf(instance)

	sliceType := reflect.SliceOf(elemType)
	slice := reflect.MakeSlice(sliceType, 0, 0)

	slicePtr := reflect.New(sliceType)
	slicePtr.Elem().Set(slice)

	sliceI := slicePtr.Interface()

	err = json.Unmarshal(bs, sliceI)
	if err != nil {
		return nil, err
	}

	v := reflect.ValueOf(sliceI)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Slice {
		panic("not a slice")
	}

	ret := []any{}
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i).Interface()
		ret = append(ret, elem)
	}

	return ret, nil
}

func DeepCopyIntoType(obj any, instance any) (any, error) {
	cop := createNewElement(instance)

	bs, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bs, &cop)
	if err != nil {
		return nil, err
	}

	return reflect.ValueOf(cop).Elem().Interface(), nil
}

func createNewElement(instance interface{}) interface{} {
	instanceType := reflect.TypeOf(instance)
	newElement := reflect.New(instanceType)

	return newElement.Interface()
}

func DeepCopyIntoMap(obj any) (any, error) {
	bs, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var v interface{}
	err = json.Unmarshal(bs, &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
