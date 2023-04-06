package yaenum

import (
	"errors"
	"reflect"
	"sync"
	"unsafe"
)

type enumList struct {
	ByName    map[string]interface{}
	ByUintptr map[uintptr]interface{}
}

type Instance[T comparable] struct {
	name    string
	uIntPtr uintptr
}

func (i *Instance[T]) String() string {
	return i.name
}

var lock sync.RWMutex
var enumCache = make(map[string]*enumList)

func Init[T comparable](val *T) *T {
	instance := reflect.ValueOf(val)
	t := instance.Type()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	name := t.Name()

	lock.Lock()
	defer lock.Unlock()
	mapByPtr := make(map[uintptr]interface{})
	mapByString := make(map[string]interface{})
	if instance.Kind() == reflect.Ptr {
		instance = instance.Elem()
	}
	var instancePtrType = reflect.TypeOf(&Instance[T]{})
	var instanceType = reflect.TypeOf(Instance[T]{})
	for i := 0; i < instance.NumField(); i++ {
		field := t.Field(i)
		if field.Type != instancePtrType {
			if field.Type == instanceType {
				panic("Please define Instance to ptr type")
			}
			continue
		}
		tag := field.Tag.Get("enum")
		if len(tag) == 0 {
			tag = field.Name
		}
		itemToCache := &Instance[T]{
			name: tag,
		}
		instance.FieldByName(field.Name).Set(reflect.ValueOf(itemToCache))
		itemToCache.uIntPtr = uintptr(unsafe.Pointer(itemToCache))
		mapByPtr[itemToCache.uIntPtr] = itemToCache
		mapByString[tag] = itemToCache
	}
	if len(mapByPtr) == 0 {
		panic("no enum tag found")
	}

	list := &enumList{
		ByName:    mapByString,
		ByUintptr: mapByPtr,
	}
	enumCache[name] = list

	return val
}

func ValueOf[T comparable](val *T, value string) (*Instance[T], error) {
	instance := reflect.ValueOf(val)
	t := instance.Type()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	name := t.Name()

	lock.RLock()
	defer lock.RUnlock()
	list, ok := enumCache[name]
	if !ok {
		return nil, errors.New("enum not initiated")
	}
	if v, ok := list.ByName[value]; ok {
		return v.(*Instance[T]), nil
	} else {
		return nil, errors.New("no enum value found")
	}
}
