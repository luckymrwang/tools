package common

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"reflect"

	"github.com/henrylee2cn/mahonia"
	"gopkg.in/mgo.v2/bson"
)

func IsInt(s interface{}) bool {
	_, e := strconv.ParseInt(ToString(s), 10, 64)
	return e == nil
}

func IsFloat(s interface{}) bool {
	_, e := strconv.ParseFloat(ToString(s), 64)
	return e == nil
}

func IsArray(v interface{}) bool {
	if IsEmpty(v) {
		return false
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Array, reflect.Slice:
		return true
	default:
		return false
	}
}

func IsMapArray(v interface{}) bool {
	a, b := v.([]interface{})
	if b {
		for _, m := range a {
			switch m.(type) {
			case map[string]interface{}:
				return true
			default:
				return false
			}
		}
	}
	return false
}

func IsJson(b []byte) bool {
	var j json.RawMessage
	return json.Unmarshal(b, &j) == nil
}

func IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}
	switch v.(type) {
	case P:
		return len(v.(P)) == 0
	case []interface{}:
		return len(v.([]interface{})) == 0
	case []P:
		return len(v.([]P)) == 0
	case *[]P:
		return len(*v.(*[]P)) == 0
	}
	return ToString(v) == ""
}

func ToInt(s interface{}, defaultValue ...int) int {
	i, e := strconv.Atoi(ToString(s))
	if e != nil && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return i
}

func ToInt64(s interface{}, defaultValue ...int64) int64 {
	switch s.(type) {
	case int64:
		return s.(int64)
	case int:
		return int64(s.(int))
	case float64:
		return int64(s.(float64))
	}
	i64, e := strconv.ParseInt(ToString(s), 10, 64)
	if e != nil && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return i64
}

func ToFloat(s interface{}, defaultValue ...float64) float64 {
	f64, e := strconv.ParseFloat(ToString(s), 64)
	if e != nil && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return f64
}

func ToBool(s interface{}) (bool, error) {
	switch ToString(s) {
	case "1", "t", "T", "true", "TRUE", "True":
		return true, nil
	case "0", "f", "F", "false", "FALSE", "False":
		return false, nil
	}
	return false, fmt.Errorf("Parse Fial:%v", ToString(s))
}

func ToString(v interface{}, def ...string) string {
	if v != nil {
		switch v.(type) {
		case bson.ObjectId:
			return v.(bson.ObjectId).Hex()
		case []byte:
			return string(v.([]byte))
		case *P, P:
			var p P
			switch v.(type) {
			case *P:
				if v.(*P) != nil {
					p = *v.(*P)
				}
			case P:
				p = v.(P)
			}
			var keys []string
			for k := range p {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			r := "P{"
			for _, k := range keys {
				r = JoinStr(r, k, ":", p[k], " ")
			}
			r = JoinStr(r, "}")
			return r
		case map[string]interface{}, []P, []interface{}:
			return JSONEncode(v)
		case int64:
			return strconv.FormatInt(v.(int64), 10)
		case []string:
			s := ""
			for _, j := range v.([]string) {
				s = JoinStr(s, ",", j)
			}
			if len(s) > 0 {
				s = s[1:]
			}
			return s
		default:
			return fmt.Sprintf("%v", v)
		}
	}
	if len(def) > 0 {
		return def[0]
	} else {
		return ""
	}
}

func ToP(v interface{}) P {
	if v != nil {
		switch v.(type) {
		case P:
			return v.(P)
		case *P:
			return *v.(*P)
		case string:
			return JSONDecode(v.(string))
		case []byte:
			return JSONDecode(string(v.([]byte)))
		case map[string]interface{}:
			return v.(map[string]interface{})
		default:
			return P{}
		}
	}
	return P{}
}

func ToStrings(v interface{}) []string {
	strs := []string{}
	if v != nil {
		switch v.(type) {
		case []interface{}:
			for _, i := range v.([]interface{}) {
				strs = append(strs, ToString(i))
			}
		case []string:
			for _, i := range v.([]string) {
				strs = append(strs, i)
			}
		case string, interface{}:
			strs = append(strs, ToString(v))
		}
	}
	return strs
}

func GbkToUtf8(s []byte) []byte {
	utf8 := mahonia.NewDecoder("gbk").ConvertString(string(s))
	return []byte(utf8)
}

//去重且遍历不会改变顺序
func StringSort(slc []string) []string {
	result := []string{}
	m := make(map[string]struct{})
	for _, e := range slc {
		if _, ok := m[e]; !ok {
			result = append(result, e)
			m[e] = struct{}{}
		}
	}
	return result
}
