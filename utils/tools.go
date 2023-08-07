package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"reflect"
)

// IsBlank 判断空值
func IsBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String, reflect.Array:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Complex64, reflect.Complex128:
		return value.Complex() == 0+0i
	case reflect.Interface, reflect.Ptr, reflect.Chan, reflect.Map, reflect.Slice:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// ToInterfaceArr 对象转 interface切片
func ToInterfaceArr(arr interface{}) []interface{} {
	if reflect.TypeOf(arr).Kind() != reflect.Slice {
		return nil
	}

	arrValue := reflect.ValueOf(arr)
	retArr := make([]interface{}, arrValue.Len())
	for k := 0; k < arrValue.Len(); k++ {
		retArr[k] = arrValue.Index(k).Interface()
	}
	return retArr
}

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

// DistinctElement 字符串数组去重
func DistinctElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// IsExistItem 判断切片中是否包含某一元素
func IsExistItem(val interface{}, arr interface{}) bool {
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(arr)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

// JSONUnmarshal json反解析
func JSONUnmarshal(jsonStr string, out interface{}) error {
	switch t := out.(type) {
	case *string:
		*t = jsonStr
		return nil
	default:
		if err := json.Unmarshal([]byte(jsonStr), out); err != nil {
			return err
		}
	}
	return nil
}

// BaseEncode base加密
func BaseEncode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
