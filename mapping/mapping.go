package mapping

import (
	"errors"
	"reflect"
	"strings"
	"time"
)

func MapTo(src interface{}, dst interface{}) error {
	srcValue := reflect.ValueOf(src)
	dstValue := reflect.ValueOf(dst)
	//如果src是指针
	if srcValue.Type().Kind() == reflect.Ptr {
		srcValue = srcValue.Elem() // 取具体内容
	}
	if dstValue.Kind() != reflect.Ptr || dstValue.IsNil() {
		return errors.New("dst is not a pointer or is nil")
	}
	dstValue = dstValue.Elem()
	item := reflect.New(dstValue.Type())
	err := setValue(srcValue, item)
	if err != nil {
		return err
	}
	if dstValue.IsValid() && dstValue.CanSet() {
		dstValue.Set(item.Elem())
	}
	return nil
}

func setValue(srcValue reflect.Value, dstValue reflect.Value) error {
	if dstValue.Kind() != reflect.Ptr || dstValue.IsNil() {
		return errors.New("dst is not a pointer or is nil")
	}
	dstType, dstValue := dstValue.Type().Elem(), dstValue.Elem()
	switch srcValue.Kind() {
	case reflect.Struct:
		if dstValue.Kind() != reflect.Struct {
			return errors.New("dst type should be a struct pointer")
		}
		for i := 0; i < dstValue.NumField(); i++ {
			fieldInfo := dstType.Field(i)
			ignore := fieldInfo.Tag.Get("ignore")
			if ignore == "true" { //映射忽略
				continue
			}
			value := findValueByName(srcValue, fieldInfo) //根据tag和字段名查找值
			if !value.IsValid() {
				continue
			}
			if value.Type().String() == "time.Time" { //处理time.Time时间类型
				if dstType.Field(i).Type.String() == "string" { //需要将time.Time转换为字符串
					timeFormat := fieldInfo.Tag.Get("timeFormat")
					if len(timeFormat) <= 0 {
						timeFormat = "2006-01-02 15:04:05" //默认时间格式
					}
					timeValue := value.Interface().(time.Time)
					//fmt.Println(fieldName + ":" + timeValue.Format(timeFormat))
					if dstValue.Field(i).IsValid() && dstValue.Field(i).CanSet() {
						dstValue.Field(i).Set(reflect.ValueOf(timeValue.Format(timeFormat)))
					}
				} else { //不需要转换 直接赋值
					if dstValue.Field(i).IsValid() && dstValue.Field(i).CanSet() && dstValue.Kind() == srcValue.Kind() {
						dstValue.Field(i).Set(value)
					}
				}
			} else {
				if dstValue.Field(i).IsValid() && dstValue.Field(i).CanSet() {
					item := reflect.New(dstValue.Field(i).Type())
					setValue(value, item)
					dstValue.Field(i).Set(item.Elem())
				}
			}
		}
	case reflect.Slice:
		if dstType.Kind() != reflect.Slice {
			//fmt.Println(dstType.Kind())
			return errors.New("dst type should be a slice")
		}
		for i := 0; i < srcValue.Len(); i++ {
			//fmt.Println(srcValue.Index(i))
			item := reflect.New(dstValue.Type().Elem())
			setValue(srcValue.Index(i), item)
			if dstValue.IsValid() && dstValue.CanSet() {
				dstValue.Set(reflect.Append(dstValue, item.Elem()))
			}
		}
	case reflect.Array:
		if dstType.Kind() != reflect.Slice && dstType.Kind() != reflect.Array {
			//fmt.Println(dstType.Kind())
			return errors.New("dst type should be a slice or a array")
		}
		if dstType.Kind() == reflect.Array {
			if dstValue.Len() < srcValue.Len() {
				return errors.New("dst array length should grater then src")
			}
			for i := 0; i < srcValue.Len(); i++ {
				//fmt.Println(srcValue.Index(i))
				item := reflect.New(dstValue.Type().Elem())
				setValue(srcValue.Index(i), item)
				if dstValue.Index(i).IsValid() && dstValue.Index(i).CanSet() {
					dstValue.Index(i).Set(item.Elem())
				}
			}
		}
		if dstType.Kind() == reflect.Slice {
			for i := 0; i < srcValue.Len(); i++ {
				//fmt.Println(srcValue.Index(i))
				item := reflect.New(dstValue.Type().Elem())
				setValue(srcValue.Index(i), item)
				if dstValue.IsValid() && dstValue.CanSet() {
					dstValue.Set(reflect.Append(dstValue, item.Elem()))
				}
			}
		}
	case reflect.Map:
		if dstType.Kind() != reflect.Map { //源数据为切片，要求目标也为map
			return errors.New("dst type should be a map")
		}
		dstValue.Set(reflect.MakeMap(dstValue.Type()))
		for _, key := range srcValue.MapKeys() {
			//fmt.Println(srcValue.MapIndex(key))
			item := reflect.New(dstValue.Type().Elem())
			setValue(srcValue.MapIndex(key), item)
			if dstValue.IsValid() && dstValue.CanSet() {
				dstValue.SetMapIndex(key, item.Elem())
			}
		}
	default:
		if dstValue.IsValid() && dstValue.CanSet() && dstValue.Kind() == srcValue.Kind() {
			dstValue.Set(srcValue)
		}
	}
	return nil
}

func findValueByName(srcValue reflect.Value, fieldInfo reflect.StructField) reflect.Value {
	fieldName := fieldInfo.Tag.Get("mappingField") //优先根据mappingField设置查找
	if len(fieldName) > 0 {
		value := srcValue.FieldByNameFunc(func(s string) bool {
			return strings.ToUpper(s) == strings.ToUpper(fieldName) //不区分大小写
		})
		return value
	}

	fieldName = fieldInfo.Name
	value := srcValue.FieldByNameFunc(func(s string) bool {
		return strings.ToUpper(s) == strings.ToUpper(fieldName) //不区分大小写
	})
	return value
}
