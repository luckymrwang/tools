package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")

	// any approach to require this configuration into your program.
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))

	fmt.Println(viper.Get("clothing.jacket")) // this would be "steve"

	env := viper.New()
	env.AutomaticEnv()
	fmt.Println(env.Get("HIVE_HOME"))

	// file1 := viper.New()
	// file1.SetConfigName("app1.yaml")
	// file1.SetConfigType("yaml")
	// file1.AddConfigPath("/Users/sino/Downloads")
	// if err := file1.ReadInConfig(); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("file1:", file1.AllSettings())

	file1 := viper.New()
	content, err := ioutil.ReadFile("/Users/sino/Downloads/app1.yaml")
	if err != nil {
		panic(err)
	}
	file1.SetConfigType("yaml")
	if err := file1.ReadConfig(bytes.NewBuffer(content)); err != nil {
		panic(err)
	}
	fmt.Println("file1:", file1.AllSettings())
	for key, _ := range file1.GetStringMap("redis") {
		fmt.Println("key:", key)
	}
	// file2 := viper.New()
	// file2.SetConfigName("app2")
	// file2.SetConfigType("yaml")
	// file2.AddConfigPath("/Users/sino/Downloads")
	// if err := file2.ReadInConfig(); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("file2:", file2.AllSettings())

	s := "http://${incloudConfigURL:100.2.29.121:32001}"
	arg := s[strings.Index(s, "${")+2 : strings.Index(s, "}")]
	fmt.Println(arg)
	defaultVal := arg[strings.Index(arg, ":")+1:]
	fmt.Println(defaultVal, strings.Index(defaultVal, "100.2.29.12:"))
	arg = arg[:strings.Index(arg, ":")]
	fmt.Println(arg)
}
