package handler

import (
	"reflect"
)

func bootHandler(modules []any) {
	for _, module := range modules {
		methods := reflect.TypeOf(module)

		for i := 0; i < methods.NumMethod(); i++ {
			method := methods.Method(i)
			handlers := reflect.ValueOf(module).MethodByName(method.Name).Call(nil)

			for _, handler := range handlers[1].Interface().([]any) {
				reflect.ValueOf(handler).MethodByName("Boot").Call([]reflect.Value{
					reflect.ValueOf(handlers[0].Interface()),
				})
			}
		}
	}
}
