package utils

import "reflect"

// SizeOfPublicStruct returns the size in number of bytes, of a struct's public fields by using reflection.
// Thanks to https://stackoverflow.com/questions/51431933/how-to-get-size-of-struct-containing-data-structures-in-go
func SizeOfPublicStruct(v interface{}) int {
	size := int(reflect.TypeOf(v).Size())
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(v)
		for i := 0; i < s.Len(); i++ {
			size += SizeOfPublicStruct(s.Index(i).Interface())
		}
	case reflect.Map:
		s := reflect.ValueOf(v)
		keys := s.MapKeys()
		size += int(float64(len(keys)) * 10.79) // approximation from https://golang.org/src/runtime/hashmap.go
		for i := range keys {
			size += SizeOfPublicStruct(keys[i].Interface()) + SizeOfPublicStruct(s.MapIndex(keys[i]).Interface())
		}
	case reflect.String:
		size += reflect.ValueOf(v).Len()
	case reflect.Struct:
		s := reflect.ValueOf(v)
		for i := 0; i < s.NumField(); i++ {
			if s.Field(i).CanInterface() {
				size += SizeOfPublicStruct(s.Field(i).Interface())
			}
		}
	}
	return size
}
