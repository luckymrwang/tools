// slice
package main

import (
	//	"encoding/json"
	"fmt"
)

type AppParam struct {
	BigAppId  int32
	Table     string
	TradeType string
	StartTime int32
	EndTime   int32
	Zids      string
}

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	//	fmt.Println("Hello World!")
	//	s := []string{"A", "B", "C", "D", "E"}
	//	x := s[2:]
	//	x[1] = "M"
	//	y := s[:4]
	//		type Ret struct {
	//			key string
	//			//			count int64
	//			//			userId string
	//		}
	//		ret := Ret{}
	//		ret.key = fmt.Sprintf("%s#%d", x, cap(x))

	//		fmt.Println(ret.key, "hello\n")
	//	z := append(y, "K")
	//	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n%v\n", x, cap(x), len(y), y, cap(y), z, s)

	//	s = s[4:4]
	//	z := s[:]
	//	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", s, cap(s), len(s), x, y, z)
	//	var s Serverslice
	//	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	//	json.Unmarshal([]byte(str), &s)

	//	app := AppParam{BigAppId: 1009, Table: "kuaiyong", TradeType: "", StartTime: 123456, EndTime: 123456, Zids: "1,2,3"}
	//	sqlStr := fmt.Sprintf("select userid, zoneid, cost, curType, trade_type from %s where createtime between ? and ? and zoneid in(%s)", app.Table, app.Zids)
	//	if len(app.TradeType) > 0 {
	//		sqlAppend := fmt.Sprintf(" and trade_type in (%s)", app.TradeType)
	//		sqlStr = sqlStr + sqlAppend
	//	}
	//	fmt.Println(sqlStr)

	//	s1 := "abcd"
	//	b1 := []byte(s1)
	//	fmt.Println(b1) // [97 98 99 100]

	//	s2 := "中文"
	//	b2 := []byte(s2)
	//	fmt.Println(b2) // [228 184 173 230 150 135], unicode，每个中文字符会由三个byte组成

	//	r := []rune(s2)
	//	fmt.Println(r) // [20013 25991], 每个字一个数值

	//	m := Message{"Alice", "Hello", 1294706395881547000}
	//	b, err := json.Marshal(m)
	//	if err != nil {
	//		fmt.Println(err)
	//	}

	//	fmt.Println(string(b))
	//	var mes Message
	//	json.Unmarshal([]byte(b), &mes)
	//	fmt.Println(mes)
	//	fmt.Println(mes.Name)

	mes := make(map[int32]*Message, 1)
	//	mes1 := make(map[int32]AppParam, 3)
	fmt.Println(mes)
	mes[1] = &Message{Name: "hello", Body: "bo", Time: 2}
	fmt.Println(mes)
	fmt.Println(mes[1].(type))
	//	fmt.Println(mes1)

}
