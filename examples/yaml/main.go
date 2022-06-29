package main

import (
	"fmt"
	"github.com/screw-coding/yaml"
)

type Student struct {
	Name    string    `yaml:"name"`
	Age     int       `yaml:"age"`
	Teacher Teacher   `yaml:"teacher"`
	Friends []Student `yaml:"friends"`
}

type Teacher struct {
	Name   string   `yaml:"name"`
	Course []string `yaml:"course"`
}

var studentYaml = `
name: lajidai
age: 30
teacher:
  name: wanglaoshi
  course: 
    - Math
    - English
friends:
  - name: xiaowang
    age: 20
  - name: xiaohong
    age: 28
`

func main() {
	student := Student{}
	_ = yaml.UnmarshalFromString(studentYaml, &student)
	fmt.Println(student)
}
