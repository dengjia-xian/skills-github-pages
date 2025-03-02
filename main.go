package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name  string
	Grade int
}
type studentMgr struct {
	students map[string]Student
}

func (s *studentMgr) ShowStudent() {
	if len(s.students) == 0 {
		fmt.Println("暂无学生")
		return
	}
	for _, v := range s.students {
		fmt.Println(v.Name, "成绩:", v.Grade)
	}
}
func (s *studentMgr) AddStudent() {
	var stu string
	fmt.Printf("请输入学生姓名：")
	fmt.Scanln(&stu)
	_, ok := s.students[stu]
	if ok {
		fmt.Println("学生已存在")
		return
	} else {
		s.students[stu] = Student{
			Name: stu,
		}
		fmt.Println("学生已添加")
		return
	}
}
func (s *studentMgr) EditStudentName() {
	var name_ago, name_now string
	fmt.Printf("请输入学生以前的姓名：")
	fmt.Scanln(&name_ago)
	fmt.Printf("请输入学生此时的姓名:")
	fmt.Scanln(&name_now)
	_, ok := s.students[name_ago]
	if ok {
		stu := Student{
			name_now,
			s.students[name_ago].Grade,
		}
		delete(s.students, name_ago)
		s.students[name_now] = stu
		fmt.Printf("%s名字已修改成%s\n", name_ago, name_now)
	} else {
		fmt.Println("此人不存在")
	}
}
func (s *studentMgr) EditStudentGrade() {
	var (
		name      string
		grade_now int
	)
	fmt.Printf("请输入学生的姓名：")
	fmt.Scanln(&name)
	fmt.Printf("请输入学生此时的成绩:")
	fmt.Scanln(&grade_now)
	_, ok := s.students[name]
	if ok {
		s.students[name] = Student{
			name,
			grade_now,
		}
		fmt.Printf("%s的成绩已修改为%d\n", name, grade_now)
	} else {
		fmt.Println("此人不存在")
	}
}
func (s *studentMgr) DelStudent() {
	var name string
	fmt.Printf("请输入你不想要的学生的姓名:")
	fmt.Scanln(&name)
	_, ok := s.students[name]
	if ok {
		delete(s.students, name)
		fmt.Printf("%s已踢出\n", name)
	} else {
		fmt.Println("此人不存在")
	}
}
func ShowMenu() {
	fmt.Println(`
1.展示所有学生信息
2.添加学生
3.修改学生姓名
4.修改学生成绩
5.将学生踢出班级
6.退出
7.保存数据
8.加载数据`)
}
func (s *studentMgr) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(s.students, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
func (s *studentMgr) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &s.students)
}
func main() {
	var p *studentMgr = &studentMgr{
		students: make(map[string]Student),
	}
	for {
	Q:
		ShowMenu()
		var n int
		fmt.Printf("请输入操作序号:")
		fmt.Scanln(&n)
		switch n {
		case 1:
			(*p).ShowStudent()
		case 2:
			(*p).AddStudent()
		case 3:
			(*p).EditStudentName()
		case 4:
			(*p).EditStudentGrade()
		case 5:
			(*p).DelStudent()
		case 6:
			goto L
		case 7:
			if err := p.SaveToFile("students.json"); err != nil {
				fmt.Println("保存数据失败:", err)
			} else {
				fmt.Println("保存数据成功")
			}
		case 8:
			if err := p.LoadFromFile("students.json"); err != nil {
				fmt.Println("加载数据失败:", err)
			} else {
				fmt.Println("加载数据成功")
			}
		default:
			{
				fmt.Println("请重新输入")
				goto Q
			}
		}
	}
L:
}
