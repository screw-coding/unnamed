package yaml

import (
	"log"
	"testing"
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

var teacher = Teacher{
	Name:   "wanglaoshi",
	Course: []string{"Math", "English"},
}

var friends = []Student{
	{
		Name: "xiaowang",
		Age:  20,
	},
	{
		Name: "xiaohong",
		Age:  28,
	},
}

var target = Student{
	Name:    "lajidai",
	Age:     30,
	Teacher: teacher,
	Friends: friends,
}

func TestUnmarshalFromString(t *testing.T) {

	t.Run("UnmarshalFromString success", func(t *testing.T) {
		student := Student{}
		err := UnmarshalFromString(studentYaml, &student)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		if student.Age != target.Age {
			log.Fatalf("age should be : %d", target.Age)
		}

		if student.Name != target.Name {
			log.Fatalf("name should be: %s", target.Name)
		}

		if len(student.Friends) != 2 {
			log.Fatalf("student should have %d friends", len(target.Friends))
		}

		if student.Teacher.Name != target.Teacher.Name {
			log.Fatalf("student's teacher should named %s", target.Teacher.Name)
		}
	})
}
