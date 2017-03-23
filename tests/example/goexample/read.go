package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/astaxie/beego/httplib"
)

func main() {
	file, err := os.Open("/Users/sino/Downloads/19.csv")
	/* 数据格式以tab键分割
	de494ae732bb552caaxxxxxxxx	2017-03-19 00:03:11	1489852991000	10.2.1	iPhone	1D280B2F-F253-4DED-9693-97AC818F7D47	02:00:00:00:00:00	1489852482000	VJNV3y	乐思-IOS_4	183.17.xx.xx
	de494ae732bb552caaxxxxxxxx	2017-03-19 00:03:15	1489852995000	10.2.1	iPhone	B267942D-695B-4B30-8983-5D97CBD76C2C	02:00:00:00:00:00	1489850039000	VJNV3y	乐思-IOS_4	140.224.xx.xx
	*/
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := map[string]string{
		"RvieUb": "微信mp",
		"UzIBJv": "UC头条",
		"iYJzqu": "UC头条",
		"ZfauUz": "爱奇艺信息流",
		"eI3Ebe": "爱奇艺信息流",
		"Azaema": "爱奇艺信息流",
		"UJN3Q3": "爱奇艺信息流",
		"a6B7nm": "百度信息流",
		"bAjIzi": "今日头条",
		"nABfAb": "今日头条",
		"iQVbaa": "智汇推-快报三小",
		"QfYnue": "智汇推-快报三小",
		"bQVzq2": "UC头条",
		"VJNV3y": "UC头条",
		"UrmqM3": "好耶-智汇推腾讯新闻三小",
		"bURbya": "百度SEM",
		"mEjeUr": "MOBUPPS",
		"aYR363": "好耶-智汇推腾讯新闻三小",
		"jaeMZn": "陌陌",
	}
	for scanner.Scan() {
		slic := strings.Split(scanner.Text(), "\t")
		if slic[7] == "" || slic[8] == "" {
			continue
		}

		adnetname, ok := m[slic[8]]
		if !ok {
			fmt.Println(slic[8], "not ok")
			continue
		}

		//		for k, v := range slic {
		//			fmt.Println(k, v)
		//		}

		url := fmt.Sprintf("xxxx/reyun_callback?appid=%s&idfa=%s&channel=%s&spreadurl=%s&activetime=%s", "2005001003", slic[5], adnetname, slic[8], slic[7])
		httpGet(url)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func httpGet(url string) {
	req := httplib.Get(url)
	str, err := req.String()
	if err != nil {
		fmt.Println(url)
	}
	fmt.Println(str)
}
