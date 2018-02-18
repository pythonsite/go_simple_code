// 学生的处理逻辑
package logic

import (
	"sync"
	"fmt"
)

type Student struct {
	Id int
	Name string
	Grade int
	Identify string
	Sex int
	BookMap map[string]*Book
	Lock sync.Mutex  //互斥锁
}

// 学生的构造函数
func NewStudent(id int,name string,grade int,identify string,sex int)(stu *Student){
	stu = &Student{
		Id:id,
		Name:name,
		Grade:grade,
		Identify:identify,
		Sex:sex,
		BookMap:make(map[string]*Book,32),
	}
	return
}

// 学生添加书籍
func (s *Student) AddBook(b *Book){
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.BookMap[b.BookId]=b
	return
}

// 学生还书
func (s *Student) BackBook(bookId string)(err error){
	s.Lock.Lock()
	defer s.Lock.Unlock()
	_,ok := s.BookMap[bookId]
	if !ok{
		// 格式化输出错误
		err = fmt.Errorf("student id:%d not exist book book_id:%s",s.Id,bookId)
		return
	}
	delete(s.BookMap,bookId)
	return
}

// 获取学生借的书籍
func (s *Student) GetBookList()(bookList[]*Book){
	s.Lock.Lock()
	defer s.Lock.Unlock()
	for _,v := range s.BookMap{
		bookList = append(bookList,v)
	}
	return
}