package client

import "reflect"

type Client interface{
	Exec(methodName string, params ... interface{}) ([]reflect.Value, error)
}
