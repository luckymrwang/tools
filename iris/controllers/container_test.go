package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"reflect"
	"sync"
	"testing"
	"time"
	"tools/iris/common"
)

type Dog struct {
	Age int `json:",omitempty"`
}

func TestHandleExecShell(t *testing.T) {
	age := 1
	d := Dog{
		Age: age,
	}

	b, _ := json.Marshal(d)
	fmt.Println(string(b))
	var dd Dog
	err := json.Unmarshal(b, &dd)
	if err != nil {
		fmt.Println("dd", err)
	}

	ddd := new(Dog)
	err = json.Unmarshal(b, ddd)
	if err != nil {
		fmt.Println("ddd", err)
	}
}

func TestGet(t *testing.T) {
	var f float64 = 1.23
	fmt.Printf("%0.f\n", f)
	a := "I'm string"
	fmt.Println(a)
	var ptr *int
	if ptr == nil {
		fmt.Println("ptr is nil")
	} else {
		fmt.Println("redirect to...")
		goto to
	}
to:
	fmt.Println("goto...")

	fmt.Printf("ptr 的值为 : %x\n", ptr)
}

func TestResponseRecorder(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))

	// Output:
	// 200
	// text/html; charset=utf-8
	// <html><body>Hello World!</body></html>
}

func TestPipe(t *testing.T) {
	c1 := exec.Command("ls")
	c2 := exec.Command("wc", "-l")

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var b2 bytes.Buffer
	c2.Stdout = &b2

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
	io.Copy(os.Stdout, &b2)
}

func TestEnv(t *testing.T) {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
}

func TestJson(t *testing.T) {
	type FruitBasket struct {
		Name    string
		Fruit   []string
		Id      int64 `json:"ref"` // 声明对应的json key
		Created time.Time
	}

	jsonData := []byte(`
    {
        "Name": "Standard",
        "Fruit": [
             "Apple",
            "Banana",
            "Orange"
        ],
        "ref": 999,
        "Created": "2018-04-09T23:00:00Z"
    }`)

	var basket FruitBasket
	err := json.Unmarshal(jsonData, &basket)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(basket.Name, basket.Fruit, basket.Id)

	data, err := json.Marshal(jsonData)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("jsonData:", string(jsonData))
	fmt.Println("jsonMarshal:", string(data))
}

type controllerManger struct {
	internalStop <-chan struct{}

	internalStopper chan<- struct{}
}

func TestChannel(t *testing.T) {
	stop := make(chan struct{})

	c := &controllerManger{
		internalStop:    stop,
		internalStopper: stop,
	}

	c.Start()

	time.Sleep(10 * time.Second)
	fmt.Println("10s...")
}

func (cm *controllerManger) Start() {
	defer close(cm.internalStopper)

	go start(cm.internalStop)
	fmt.Println("start.....")

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("it's time 5s.")
		return
	}
}

func start(stop <-chan struct{}) {
	<-stop
	fmt.Println("start is over....")
}

func TestMask(t *testing.T) {
	// This mask corresponds to a /31 subnet for IPv4.
	bm := net.CIDRMask(26, 32)
	fmt.Println(bm.String())
	ones, bits := bm.Size()
	fmt.Println(ones, bits)

	// This mask corresponds to a /64 subnet for IPv6.
	fmt.Println(net.CIDRMask(64, 128))

	fmt.Printf("ip: %s\n", net.ParseIP("192.166.1.4").String())
	ip, cidr, err := net.ParseCIDR("192.166.1.4")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("ip: %s, cidr: %s\n", ip.String(), cidr.String())
}

func TestStopCh(t *testing.T) {
	stop := make(chan struct{})
	worker(stop)
	time.Sleep(3 * time.Second)
	close(stop)
	fmt.Println("close stop.")
}

func worker(stopCh <-chan struct{}) {
	t := time.NewTicker(1 * time.Second)
	go func() {
		defer fmt.Println("worker exit")
		// Using stop channel explicit exit
		for {
			select {
			case <-stopCh:
				fmt.Println("Recv stop signal")
				return
			case <-t.C:
				fmt.Println("Working .")
			}
		}
	}()
	return
}

func TestSlice(t *testing.T) {
	subnets := make([]string, 0)
	subnets = append(subnets, getSlice()...)
	fmt.Println(subnets)
	return
}

func getSlice() []string {
	return nil
}

func TestEqueue(t *testing.T) {
	a := []string{"a", "b"}
	b := []string{"b", "a"}
	fmt.Println(reflect.DeepEqual(a, b))
	return
}

func TestWg(t *testing.T) {
	var wg sync.WaitGroup
	slics := make([]string, 10)
	slics = append(slics, "bbbb")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go cha(&wg, slics)
	}
	wg.Wait()
	fmt.Println(slics)
}

func cha(wg *sync.WaitGroup, aa []string) {
	defer wg.Done()

	aa[1] = "kkkkkk"
	return
}

func TestSlice2(t *testing.T) {
	slice := make([]int, 2, 3)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	fmt.Printf("slice: %v, addr: %p \n", slice, slice)
	changeSlice(slice)
	fmt.Printf("slice: %v, addr: %p \n", slice, slice)
}

func changeSlice(s []int) {
	s[1] = 111
	s = append(s, 3)
	//s = append(s, 4)
	fmt.Printf("func s: %v, addr: %p \n", s, s)
}

func TestStruct1(t *testing.T) {
	ddd := new(Dog)
	fmt.Println(ddd)
	if ddd == nil {
		fmt.Println("ddd is nil")
	}

	var ggg *Dog
	fmt.Println(ggg)
	if ggg == nil {
		fmt.Println("ggg is nil")
	}
	data := make(map[string]interface{})
	fmt.Println(data)
	mm1(nil)
}

func mm1(datamap map[string]interface{}) {
	fmt.Println(datamap)
	if len(datamap) > 0 {
		fmt.Println("not 0")
	} else {
		fmt.Println("0")
	}
	da1 := make(map[string]string)
	da1["a"] = "a"
	da1["b"] = "b"

	da2 := make(map[string]string)
	da2["a"] = "a"
	da2["c"] = "c"

	for key, val := range da1 {
		if da2[key] == val {
			fmt.Println("equal: ", key, val)
		}
	}
	fmt.Println(da1["cccccc"], "oo")
}

func TestAsyncCall(t *testing.T) {
	for i := 1; i <= 3; i++ {
		// 为每个协程创建一个带有超时的上下文
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(i)*2*time.Second)
		defer cancel() // 确保取消上下文以释放资源

		go downloadData(ctx, i)
	}

	// 等待一段时间以便所有协程可以执行
	time.Sleep(10 * time.Second)
}

func downloadData(ctx context.Context, id int) {
	dealLine, dead := ctx.Deadline()
	fmt.Printf("协程 %d,deadLine:%v,dead:%v\n", id, dealLine, dead)
	select {
	case <-time.After(5 * time.Second): // 模拟下载任务耗时
		fmt.Printf("协程 %d: 数据下载完成\n", id)
	case <-ctx.Done():
		fmt.Printf("协程 %d: 下载超时: %v\n", id, ctx.Err())
	}
}

func TestClosure(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
		time.Sleep(2 * time.Second)
	}
}

type A struct {
	B string
	C int
	D bool
}

func TestA(t *testing.T) {
	a := &A{}
	fmt.Println(a)

	var b *A
	fmt.Println(b)

	c := new(A)
	fmt.Println(c)

	d := new([]int)
	fmt.Println(d)

	e := make([]int, 0)
	fmt.Println(e)

	var f A
	fmt.Println(f)
}

type S1 struct {
	AA string
	BB map[string]string
}

func TestS1(t *testing.T) {
	s := S1{}
	fmt.Println(s)

	s.AA = "ceshi\n"
	s.BB = map[string]string{"kkk": "cedsn\n"}
	fmt.Println(s)
	str := common.JSONEncode(s)
	fmt.Println(str)
}

type Fence struct {
	runtimeId string
	gpuCnt    int
	status    string
}

func TestFence(t *testing.T) {
	cnts := []Fence{
		{runtimeId: "a", gpuCnt: 1, status: "Ready"},
		{runtimeId: "b", gpuCnt: 1, status: "Ready"},
		{runtimeId: "c", gpuCnt: 1, status: "xx"},
		{runtimeId: "a", gpuCnt: 3, status: "Ready"},
		{runtimeId: "b", gpuCnt: 2, status: "xx"},
		{runtimeId: "a", gpuCnt: 1, status: "Ready"},
		{runtimeId: "c", gpuCnt: 2, status: "Rexxady"},
	}
	runtimeGPUNum := make(map[string]int)
	for _, cnt := range cnts {
		runtimeGPUNum[cnt.runtimeId] += cnt.gpuCnt
	}
}
