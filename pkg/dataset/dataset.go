package dataset

import "errors"

type I7e interface {
	Load(config string) (interface{}, error)
}

type Type string

type EmptyData []interface{}

type DataDesc map[string]interface{}

var Reg = make(map[Type]I7e, 0)

func Load(typ Type, config string) (interface{}, error) {
	if instance, exist := Reg[typ]; exist {
		return instance.Load(config)
	}
	return nil, errors.New("dataset type not found")
}



