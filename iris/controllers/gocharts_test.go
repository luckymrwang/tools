package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// 实际值
//var actualData = []string{
//	"2024-07-24 12:00:30,0",
//	"2024-07-24 12:01:00,64",
//	"2024-07-24 12:01:30,200",
//	"2024-07-24 12:02:00,340",
//	"2024-07-24 12:02:30,470",
//	"2024-07-24 12:03:00,607",
//	"2024-07-24 12:03:30,685",
//}

var actualData = []string{
	"2024-09-05 18:24:00,18.4",
	"2024-09-05 18:24:30,51.4",
	"2024-09-05 18:25:00,84.3",
	"2024-09-05 18:25:30,114",
	"2024-09-05 18:26:00,143",
	"2024-09-05 18:26:30,173",
	"2024-09-05 18:27:00,202",
	"2024-09-05 18:27:30,232",
	"2024-09-05 18:28:00,261",
	"2024-09-05 18:28:30,290",
	"2024-09-05 18:29:00,320",
	"2024-09-05 18:29:30,349",
	"2024-09-05 18:30:00,378",
	"2024-09-05 18:30:30,408",
	"2024-09-05 18:31:00,437",
	"2024-09-05 18:31:30,466",
	"2024-09-05 18:32:00,488",
}

var actualData3 = []string{
	"2024-09-04 17:08:30,16.8",
	"2024-09-04 17:09:00,49",
	"2024-09-04 17:09:30,82.4",
	"2024-09-04 17:10:00,112",
	"2024-09-04 17:10:30,141",
	"2024-09-04 17:11:00,171",
	"2024-09-04 17:11:30,200",
	"2024-09-04 17:12:00,230",
	"2024-09-04 17:12:30,259",
	"2024-09-04 17:13:00,288",
	"2024-09-04 17:13:30,318",
	"2024-09-04 17:14:00,347",
	"2024-09-04 17:14:30,376",
	"2024-09-04 17:15:00,406",
	"2024-09-04 17:15:30,435",
	"2024-09-04 17:16:00,464",
	"2024-09-04 17:16:30,489",
}

// 实例数
var scaleData = []string{
	"2024-09-05 18:24:00,1",
	"2024-09-05 18:24:30,1",
	"2024-09-05 18:25:00,1",
	"2024-09-05 18:25:30,2",
	"2024-09-05 18:26:00,2",
	"2024-09-05 18:26:30,2",
	"2024-09-05 18:27:00,2",
	"2024-09-05 18:27:30,2",
	"2024-09-05 18:28:00,3",
	"2024-09-05 18:28:30,3",
	"2024-09-05 18:29:00,3",
	"2024-09-05 18:29:30,3",
	"2024-09-05 18:30:00,3",
	"2024-09-05 18:30:30,4",
	"2024-09-05 18:31:00,4",
	"2024-09-05 18:31:30,4",
	"2024-09-05 18:32:00,4",
}

// 理论值
var expectData2 = []string{
	"2024-09-04 17:08:00,0",
	"2024-09-04 17:08:30,29",
	"2024-09-04 17:09:00,58",
	"2024-09-04 17:09:30,90",
	"2024-09-04 17:10:00,120",
	"2024-09-04 17:10:30,150",
	"2024-09-04 17:11:00,180",
	"2024-09-04 17:11:30,210",
	"2024-09-04 17:12:00,240",
	"2024-09-04 17:12:30,270",
	"2024-09-04 17:13:00,300",
	"2024-09-04 17:13:30,330",
	"2024-09-04 17:14:00,360",
	"2024-09-04 17:14:30,390",
	"2024-09-04 17:15:00,420",
	"2024-09-04 17:15:30,450",
	"2024-09-04 17:16:00,480",
}

var expectData = []string{
	"2024-09-05 18:24:00,30",
	"2024-09-05 18:24:30,60",
	"2024-09-05 18:25:00,90",
	"2024-09-05 18:25:30,120",
	"2024-09-05 18:26:00,150",
	"2024-09-05 18:26:30,180",
	"2024-09-05 18:27:00,210",
	"2024-09-05 18:27:30,240",
	"2024-09-05 18:28:00,270",
	"2024-09-05 18:28:30,300",
	"2024-09-05 18:29:00,330",
	"2024-09-05 18:29:30,360",
	"2024-09-05 18:30:00,390",
	"2024-09-05 18:30:30,420",
	"2024-09-05 18:31:00,450",
	"2024-09-05 18:31:30,480",
	"2024-09-05 18:32:00,510",
}

// 解析时间戳字符串为 time.Time 对象
func parseTimestampsAndCalculateTPS(timestamps []string) ([]string, []float64) {
	var tps []float64
	var labels []string
	for _, val := range timestamps {
		data := strings.Split(val, ",")
		t, _ := time.Parse("2006-01-02 15:04:05", data[0])
		labels = append(labels, t.Format("15:04:05"))
		tp, _ := strconv.ParseFloat(data[1], 64)
		tps = append(tps, tp)
	}
	return labels, tps
}

func parseTimestampsAndCalculateScale(timestamps []string) ([]string, []float64) {
	var scale []float64
	var labels []string
	for _, val := range timestamps {
		data := strings.Split(val, ",")
		t, _ := time.Parse("2006-01-02 15:04:05", data[0])
		labels = append(labels, t.Format("15:04:05"))
		tp, _ := strconv.ParseFloat(data[1], 64)
		scale = append(scale, tp)
	}
	return labels, scale
}

func TestFftEstimator(t *testing.T) {
	// 创建折线图
	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Transactions Per Second (TPS)",
	}))

	// 实际值
	labels, tps := parseTimestampsAndCalculateTPS(actualData)
	// 期望值
	_, expectTps := parseTimestampsAndCalculateTPS(expectData)

	// 设置数据
	line.SetXAxis(labels).
		AddSeries("实际值", generateLineItems(tps),
			charts.WithLineStyleOpts(opts.LineStyle{Color: "green"})).
		AddSeries("理论值", generateLineItems(expectTps),
			charts.WithLineStyleOpts(opts.LineStyle{Color: "red", Type: "dashed"}))

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		if err := components.NewPage().AddCharts(line, lineScale()).Render(w); err != nil {
			// nothing to do
			t.Error(err)
		}
	})
	fmt.Println("Open your browser and access 'http://localhost:7001'")
	http.ListenAndServe(":7001", nil)
}

// 生成折线图数据项
func generateLineItems(data []float64) []opts.LineData {
	items := make([]opts.LineData, 0)
	for _, v := range data {
		items = append(items, opts.LineData{Value: v, Symbol: "none"})
	}
	return items
}

func lineScale() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Pod Counts"}),
	)

	// 实例数
	labels, scale := parseTimestampsAndCalculateScale(scaleData)

	line.SetXAxis(labels).
		AddSeries("实例数", generateLineItems(scale)).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{
				Step: true,
			}),
		)
	return line
}
