package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Model Struct
type MovieCopy struct {
	Id           int `orm:"pk"`
	SourceId     string
	Name         string `orm:"size(100)"`
	EnglishName  string
	Year         int
	Region       string
	Length       string
	Type         string
	Rating       float32
	RatingNumber int
	Sitcom       int
	Episode      int
	CreateTime   time.Time
	UpdateTime   time.Time
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(106.75.27.238:3307)/sample?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(MovieCopy))
}

func main() {
	o := orm.NewOrm()

	movie := MovieCopy{
		SourceId:     "1303578",
		Name:         "米兰奇迹",
		EnglishName:  "Miracolo a Milano",
		Year:         1951,
		Region:       "意大利",
		Length:       "100分钟|France:92分钟|90分钟",
		Type:         "剧情|喜剧|奇幻",
		Rating:       7.8,
		RatingNumber: 701,
		Sitcom:       0,
		Episode:      0,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}

	for i := 1; i < 10000; i++ {
		movies := make([]MovieCopy, 0)
		for j := 0; j < 100; j++ {
			movies = append(movies, movie)
		}

		_, err := o.InsertMulti(200, movies)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	//	// insert
	//	id, err := o.Insert(&movie)
	//	fmt.Printf("ID: %d, ERR: %v\n", id, err)
}
