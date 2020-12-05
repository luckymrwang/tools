package common

import (
	"database/sql"
	"reflect"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type P map[string]interface{}

func (p *P) Copy() P {
	pn := make(P)
	for k, v := range *p {
		pn[k] = v
	}
	return pn
}

func (p P) CopyFrom(from P) {
	for k, v := range from {
		if IsEmpty(p[k]) {
			p[k] = v
		}
	}
}

func (p *P) ToInt(s ...string) {
	for _, k := range s {
		v := ToString((*p)[k])
		if !IsEmpty(v) {
			(*p)[k] = ToInt(v)
		}
	}
}

func (p *P) Like(s ...string) {
	for _, k := range s {
		v := ToString((*p)[k])
		if !IsEmpty(v) {
			(*p)[k] = &bson.RegEx{Pattern: v, Options: "i"}
		}
	}
}

func (p *P) ToP(s ...string) (r P) {
	for _, k := range s {
		v := ToString((*p)[k])
		r = JSONDecode(v)
		(*p)[k] = r
	}
	return
}

func (p *P) Get(k string, def interface{}) interface{} {
	r := (*p)[k]
	if r == nil {
		r = def
	}
	return r
}

func SetKv(p P, k string, v []string) {
	if len(v) == 1 {
		if len(v[0]) > 0 {
			p[k] = v[0]
		}
	} else {
		p[k] = v
	}
}

func ModelToP(o interface{}) P {
	info := P{}
	if o == nil {
		return info
	} else {
		s := reflect.ValueOf(o).Elem()
		for i := 0; i < s.NumField(); i++ {
			f := s.Type().Field(i)
			key := f.Tag.Get("json")
			if key == "" || key == "-" {
				continue
			}
			value := s.Field(i).Interface()
			if key == "update_time" {
				//value = ToBeiJingTime(value.(time.Time))
				value = value.(time.Time).Format("2006-01-02 15:04:05")
			}
			info[key] = value
		}
		return info
	}
}

func ModelToArrayP(o []*interface{}) []P {
	var array = []P{}
	for _, v := range o {
		array = append(array, ModelToP(v))
	}
	return array
}

func SQLRowToP(rows *sql.Rows) (p []P, err error) {
	columns, err := rows.Columns()
	if err != nil {
		Error("db mysql parse error:", err.Error())
		return nil, err
	}
	result := []P{}
	for rows.Next() {
		column := make([]interface{}, len(columns))
		args := make([]interface{}, 0)
		for k := range column {
			args = append(args, &column[k])
		}
		err = rows.Scan(args...)
		entry := P{}
		for i, col := range columns {
			var v interface{}
			val := column[i]
			switch val.(type) {
			default:
				if val == nil {
					v = ""
				} else {
					v = val
				}
			case []byte:
				v = string(val.([]byte))
			case string:
				v = val
			case time.Time:
				v = val.(time.Time).Format("2006-01-02 15:04:05")
			}

			// 注意一律小写处理
			entry[strings.ToLower(col)] = v
		}
		result = append(result, entry)
	}
	return result, nil
}
