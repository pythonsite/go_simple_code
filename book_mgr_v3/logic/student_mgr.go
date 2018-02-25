// 学生管理的处理逻辑
package logic

import (
	"sync"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type StudentMgr struct {
	// 学生id对应student map
	StudentMap map[int]*Student
	Lock sync.Mutex
}

// 学生管理的构造函数
func NewStudentMgr()(*StudentMgr){
	s :=  &StudentMgr{
		StudentMap:make(map[int]*Student,32),
	}
	s.load()
	return s
}

// 添加学生
func (s *StudentMgr) AddStudent(stu *Student)(err error){
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.StudentMap[stu.Id] = stu
	s.Save()
	return
}

// 根据学生id查询学生的信息
func (s *StudentMgr) GetStudentByid(id int)(stu *Student,err error){
	s.Lock.Lock()
	defer s.Lock.Unlock()

	stu,ok:=s.StudentMap[id]
	if !ok{
		err = fmt.Errorf("student id %d is not exists",id)
		return
	}
	return
}

// 获取学生的借的书籍信息
func (s *StudentMgr) GetStudentBorrowBooks(id int)(bookList []*Book,err error){
	stu,err := s.GetStudentByid(id)
	if err != nil{
		return
	}
	bookList = stu.GetBookList()
	return
}


// 数据持续化到文件
func(s *StudentMgr) Save(){
	data,err := json.Marshal(s)
	if err != nil{
		fmt.Printf("save failed,err:%v\n",err)
		return
	}
	err = ioutil.WriteFile(StudentMgrSavaPaht,data,0666)
	if err != nil{
		fmt.Printf("write file failed,err:%v",err)
		return
	}
	return
}

func(s *StudentMgr) load(){
	data,err := ioutil.ReadFile(StudentMgrSavaPaht)
	if err != nil{
		fmt.Printf("load failed,err:%v",err)
		return
	}
	err = json.Unmarshal(data,s)
	if err != nil{
		fmt.Printf("unmarshal failed,err:%v",err)
		return
	}
	fmt.Printf("load data from disk success\n")


}