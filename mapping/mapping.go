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

	dstType, dstValue = dstType.Elem(), dstValue.Elem()
	switch srcType.Kind() {
	case reflect.Struct: //结构体
		item := reflect.New(dstValue.Type())
		setValue(srcValue, item)
		if dstValue.CanSet() {
			dstValue.Set(item.Elem())
		}
	case reflect.Slice: //切片
		if dstType.Kind() != reflect.Slice {
			//fmt.Println(dstType.Kind())
			return errors.New("dst type should be a slice")
		}
		for i := 0; i < srcValue.Len(); i++ {
			//fmt.Println(srcValue.Index(i))
			item := reflect.New(dstValue.Type().Elem())
			setValue(srcValue.Index(i), item)
			if dstValue.CanSet() {
				dstValue.Set(reflect.Append(dstValue, item.Elem()))
			}
		}
	case reflect.Array: //数组
		if dstType.Kind() != reflect.Slice && dstType.Kind() != reflect.Array {
			//fmt.Println(dstType.Kind())
			return errors.New("dst type should be a slice or a array")
		}
		if dstType.Kind() == reflect.Array {
			if dstValue.Len() < srcValue.Len() {
				return errors.New("dst array length should grater then src")
			}
			for i := 0; i < srcValue.Len(); i++ {
				fmt.Println(srcValue.Index(i))
				item := reflect.New(dstValue.Type().Elem())
				setValue(srcValue.Index(i), item)
				if dstValue.Index(i).CanSet() {
					dstValue.Index(i).Set(item.Elem())
				}
			}
		}
		if dstType.Kind() == reflect.Slice {
			for i := 0; i < srcValue.Len(); i++ {
				//fmt.Println(srcValue.Index(i))
				item := reflect.New(dstValue.Type().Elem())
				setValue(srcValue.Index(i), item)
				if dstValue.CanSet() {
					dstValue.Set(reflect.Append(dstValue, item.Elem()))
				}
			}
		}
	case reflect.Map: //map
		if dstType.Kind() != reflect.Map { //源数据为切片，要求目标也为map
			return errors.New("dst type should be a map")
		}
		for _, key := range srcValue.MapKeys() {
			//fmt.Println(srcValue.MapIndex(key))
			item := reflect.New(dstValue.Type())
			setValue(srcValue.MapIndex(key), item)
			dstValue.SetMapIndex(key, srcValue.MapIndex(key))
		}
	default:
		panic(fmt.Sprintf("%v cannot mapping", srcType.Kind()))
	}
	return nil
}

func setValue(srcValue reflect.Value, dstValue reflect.Value) error {
	if dstValue.Kind() != reflect.Ptr || dstValue.IsNil() {
		return errors.New("dst is not a pointer or is nil")
	}
	dstType, dstValue := dstValue.Type().Elem(), dstValue.Elem()
	if srcValue.Kind() == reflect.Struct {
		if dstValue.Kind() != reflect.Struct {
			return errors.New("dst type should be a struct pointer")
		}
		for i := 0; i < dstValue.NumField(); i++ {
			fieldInfo := dstType.Field(i)
			fieldName := fieldInfo.Name
			value := srcValue.FieldByName(fieldName)
			if !value.IsValid() || value.Kind() != fieldInfo.Type.Kind() {
				continue
			}
			if value.Kind() == reflect.Struct {
				item := reflect.New(dstValue.Field(i).Type())
				setValue(value, item)
				if dstValue.Field(i).CanSet() {
					dstValue.Field(i).Set(item.Elem())
				}
			} else {
				if dstValue.Field(i).IsValid() && dstValue.Field(i).CanSet() {
					dstValue.Field(i).Set(value)
				}
			}
		}
	} else {
		if dstValue.CanSet() {
			dstValue.Set(srcValue)
		}
	}
	return nil
}
