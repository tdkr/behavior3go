package loader

import (
	"fmt"
	. "github.com/tdkr/behavior3go/core"
	"reflect"
)

//定义注册结构map
type RegisterStructMaps struct {
	maps map[string]reflect.Type
}

func NewRegisterStructMaps() *RegisterStructMaps {
	return &RegisterStructMaps{make(map[string]reflect.Type)}
}

//根据name初始化结构
//在这里根据结构的成员注解进行DI注入，这里没有实现，只是简单都初始化
func (rsm *RegisterStructMaps) New(name string) (interface{}, error) {
	fmt.Println("New ", name)
	var c interface{}
	var err error
	if v, ok := rsm.maps[name]; ok {
		c = reflect.New(v).Interface()
		fmt.Println("found ", name, "  ", reflect.TypeOf(c))
		return c, nil
	} else {
		err = fmt.Errorf("not found %s struct", name)
		fmt.Println("New no found", name, "  ", len(rsm.maps))
	}
	return nil, err
}

//查询是否存在
func (rsm *RegisterStructMaps) CheckElem(name string) bool {
	if _, ok := rsm.maps[name]; ok {
		return true
	}
	return false
}

//根据名字注册实例
func (rsm *RegisterStructMaps) Register(name string, c interface{}) {
	rsm.maps[name] = reflect.TypeOf(c).Elem()
}

func (rsm *RegisterStructMaps) CreateNode(name string) IBaseNode {
	if rsm.CheckElem(name) {
		if tnode, err := rsm.New(name); err == nil {
			node := tnode.(IBaseNode)
			return node
		} else {
			fmt.Println("new ", name, " err:", err)
		}
	}
	panic("RegisterStructMaps.CreateNode: Invalid node name:" + name)
	return nil
}
