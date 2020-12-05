package common

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"html/template"
	"math/rand"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"
)

var IMPORT_EXPIRE = 3 * 60

func Error(a ...interface{}) {
	fmt.Println(a...)
}

func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}
	vi := reflect.ValueOf(v)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

//JSONDecode Json格式解析
func JSONDecode(str string) (p P) {
	_bytes := []byte(str)
	if len(_bytes) <= 0 {
		return p
	}
	err := json.Unmarshal(_bytes, &p)
	if err != nil {
		Error(fmt.Sprintf("JSONDecode Error:%v,Str:%v", err.Error(), str))
	}
	return
}

//JSONDecodeArrayCheck Json数组格式解析(返回解析错误)
func JSONDecodeArrayCheck(str string) (p []P, err error) {
	p = []P{}
	_bytes := []byte(str)
	if len(_bytes) <= 0 {
		return
	}
	d := json.NewDecoder(bytes.NewReader(_bytes))
	d.UseNumber()
	err = d.Decode(&p)
	if err != nil {
		Error(fmt.Sprintf("JSONDecodeArrayCheck Error:%v,Str:%v", err.Error(), str))
	}
	return
}

//JSONDecodeArray Json数组格式解析
func JSONDecodeArray(str string) (p []P) {
	p = []P{}
	_bytes := []byte(str)
	if len(_bytes) <= 0 {
		return
	}
	d := json.NewDecoder(bytes.NewReader(_bytes))
	d.UseNumber()
	err := d.Decode(&p)
	if err != nil {
		Error(fmt.Sprintf("JSONDecodeArray Error:%v,Str:%v", err.Error(), str))
	}
	return
}

//JSONEncodeCheck Json格式编码(返回转换错误)
func JSONEncodeCheck(value interface{}) (string, error) {
	_bytes, err := json.Marshal(value)
	if err != nil {
		Error("JSONEncodeCheck Error:", value, err)
		return "", err
	}
	return string(_bytes), nil
}

//JSONEncode Json格式编码
func JSONEncode(value interface{}) string {
	_bytes, err := json.Marshal(value)
	if err != nil {
		Error("JSONEncode Error:", value, err)
		return ""
	}
	return string(_bytes)
}

func InArray(s string, a []string) bool {
	for _, x := range a {
		if x == s {
			return true
		}
	}
	return false
}

func StartsWith(s string, a ...string) bool {
	for _, x := range a {
		if strings.HasPrefix(s, x) {
			return true
		}
	}
	return false
}

func EndsWith(s string, a ...string) bool {
	for _, x := range a {
		if strings.HasSuffix(s, x) {
			return true
		}
	}
	return false
}

func Rand(start int, end int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(end)
	if r < start {
		r = start + rand.Intn(end-start)
	}
	//time.Sleep(1 * time.Nanosecond)
	return r
}

func Replace(src string, find []string, r string) string {
	for _, v := range find {
		src = strings.Replace(src, v, r, -1)
	}
	return src
}

func ReplaceRegx(src string, regex []string, r string) string {
	for _, v := range regex {
		src = strings.Replace(src, v, r, -1)
		re := regexp.MustCompile(v)
		src = re.ReplaceAllString(src, r)
	}
	return src
}

func Count(src string, find []string) (c int) {
	for _, v := range find {
		c += strings.Count(src, v)
	}
	return
}

func Trim(str string) string {
	return strings.TrimSpace(str)
}

func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func JoinStr(val ...interface{}) (r string) {
	for _, v := range val {
		r += ToString(v)
	}
	return
}

func GetRandomString(number int) string {
	str := "23456789abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < number; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func RenderTpl(tpl string, data interface{}) string {
	var bb bytes.Buffer
	//t, err := template.ParseFiles(tpl)
	t, err := template.New(Md5(tpl)).Parse(tpl)
	if err != nil {
		Error(err)
	}
	t.Execute(&bb, data)
	return bb.String()
}

//Md5 MD5加密
func Md5(s ...interface{}) (r string) {
	return getHash("md5", s...)
}

func getHash(algorithm string, s ...interface{}) (r string) {
	r = hex.EncodeToString(hashBytes(algorithm, s...))
	return
}

func Exec(cmd string, exp ...int) (str string, e error) {
	osname := runtime.GOOS
	var r *exec.Cmd
	if osname == "windows" {
		r = exec.Command("cmd", "/c", cmd)
	} else {
		r = exec.Command("/bin/bash", "-c", cmd)
	}
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	r.Stdout = &stdout
	r.Stderr = &stderr
	r.Start()
	if len(exp) < 1 {
		exp = []int{IMPORT_EXPIRE}
	}
	done := make(chan error)
	go func() { done <- r.Wait() }()

	timeout := time.After(time.Duration(exp[0]) * time.Second)
	select {
	case <-timeout:
		r.Process.Kill()
		e = errors.New("Command timed out")
	case e = <-done:
		str = fmt.Sprintf("%v:%v", stdout.String(), stderr.String())
	}
	if e != nil {
		Error("Exec Failed:", e.Error(), str, cmd)
	} else {
		Error(fmt.Sprintf("Exec End:[%v]", stdout.String()), cmd)
	}

	return
}

func GetMD5Sign(p P) string {
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var str string
	for _, v := range keys {
		str += v + ToString(p[v])
	}
	return Md5(str)
}

func hashBytes(algorithm string, s ...interface{}) (r []byte) {
	var h hash.Hash
	switch algorithm {
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha2", "sha256":
		h = sha256.New()
	}
	for _, value := range s {
		switch value.(type) {
		case []byte:
			h.Write(value.([]byte))
		default:
			h.Write([]byte(ToString(value)))
		}
	}
	r = h.Sum(nil)
	return
}
