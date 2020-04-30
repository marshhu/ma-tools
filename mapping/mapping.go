package mapping

import (
	"errors"
	"fmt"
	"reflect"
)

func MapTo(src interface{}, dst interface{}) error {
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)
	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	//如果src是指针
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem() // 取具体内容
	}

	if dstValue.Kind() != reflect.Ptr || dstValue.IsNil() {
		return errors.New("dst is not a pointer or is nil")
	}

	switch srcType.Kind() {
	case reflect.Struct: //结构体
		if dstType.Elem().Kind() != reflect.Struct { //源数据为结构体，目标结构则要求为结构体指针
			return errors.New("dst type should be a struct pointer")
		}
		if dstValue.IsZero() {
			ev := dstValue.Elem()
			if ev.CanSet() {
				ev.Set(reflect.New(ev.Type().Elem()))
			}
		}
		dstType, dstValue := dstType.Elem(), dstValue.Elem()
		if !dstValue.IsValid() {
			return errors.New("dst value Is invalid")
		}
		for i := 0; i < dstValue.NumField(); i++ {
			fieldInfo := dstType.Field(i)
			fieldName := fieldInfo.Name
			value := srcValue.FieldByName(fieldName)
			if !value.IsValid() || value.Type() != fieldInfo.Type {
				tag := fieldInfo.Tag // a reflect.StructTag
				tagName := tag.Get("json")
				value = srcValue.FieldByName(tagName)
				if !value.IsValid() || value.Type() != fieldInfo.Type {
					continue
				}
			}
			if dstValue.Field(i).IsValid() && dstValue.Field(i).CanSet() {
				dstValue.Field(i).Set(value)
			}
		}
	case reflect.Slice, reflect.Array: //切片
		if dstType.Kind() != reflect.Slice || dstType.Kind() != reflect.Array {
			return errors.New("dst type should be a slice")
		}
		for i := 0; i <= srcValue.Len(); i++ {
			fmt.Println(srcValue.Index(i))
			item := reflect.New(dstValue.Type().Elem()).Elem()
			copyValue(srcValue.Index(i), item)
			dstValue.Set(reflect.Append(dstValue, item))
		}
	case reflect.Map: //map
		if dstType.Kind() != reflect.Map { //源数据为切片，要求目标也为map
			return errors.New("dst type should be a map")
		}
		for _, key := range srcValue.MapKeys() {
			fmt.Println(srcValue.MapIndex(key))
			item := reflect.New(dstValue.Type().Elem()).Elem()
			copyValue(srcValue.MapIndex(key), item)
			dstValue.Set(reflect.Append(dstValue, item))
		}
	default:
		panic(fmt.Sprintf("%v cannot mapping", srcType.Kind()))

	}
	return nil
}

func copyValue(srcValue reflect.Value, dstValue reflect.Value) error {

	return nil
}
