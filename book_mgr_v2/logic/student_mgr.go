// 学生管理的处理逻辑
package logic

import (
	"sync"
	"fmt"
)

type StudentMgr struct {
	// 学生id对应student map
	StudentMap map[int]*Student
	Lock sync.Mutex
}

// 学生管理的构造函数
func NewStudentMgr()(*StudentMgr){
	return &StudentMgr{
		StudentMap:make(map[int]*Student,32),
	}
}

// 添加学生
func (s *StudentMgr) AddStudent(stu *Student)(err error){
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.StudentMap[stu.Id] = stu
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