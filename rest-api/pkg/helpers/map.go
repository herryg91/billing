package helpers

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
)

func StructToMap(in interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	tag := "json"

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return map[string]interface{}{}
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)
		if tagv := fi.Tag.Get(tag); tagv != "" {
			// set key of map to value in struct field
			out[tagv] = v.Field(i).Interface()
		}
	}
	return out
}

func MapToStruct(in map[string]interface{}, data interface{}) {
	cfg := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   &data,
		TagName:  "json",
	}
	decoder, _ := mapstructure.NewDecoder(cfg)
	decoder.Decode(in)
}
