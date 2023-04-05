package enum

import (
	"errors"
	"reflect"
	"sync"
	"unsafe"
)

type enumList struct {
	ByName    map[string]*Instance
	ByUintptr map[uintptr]*Instance
}

type Instance struct {
	name    string
	uIntPtr uintptr
}

func (i *Instance) String() string {
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
	mapByPtr := make(map[uintptr]*Instance)
	mapByString := make(map[string]*Instance)
	if instance.Kind() == reflect.Ptr {
		instance = instance.Elem()
	}
	for i := 0; i < instance.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("enum")
		if tag != "" {
			itemToCache := &Instance{
				name: tag,
			}
			instance.FieldByName(field.Name).Set(reflect.ValueOf(itemToCache))
			itemToCache.uIntPtr = uintptr(unsafe.Pointer(itemToCache))
			mapByPtr[itemToCache.uIntPtr] = itemToCache
			mapByString[tag] = itemToCache
		}
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

func ValueOf(val interface{}, value string) (*Instance, error) {
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
		return v, nil
	} else {
		return nil, errors.New("no enum value found")
	}
}
