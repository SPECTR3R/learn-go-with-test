package walk

import (
	"log"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	log.Println("walk:", val)
	field := val.Field(0)
	fn(field.String())
}
