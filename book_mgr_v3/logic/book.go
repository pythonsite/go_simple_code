// 关于书籍的处理逻辑

package logic

import (
	"errors"
	"sync"
)

// 定义一个书的结构体，包含名字，数量，作者和出版日期
type Book struct {
	BookId string
	Name string
	Num int
	Author string
	PublishDate int64
	BorrowCount int
	Lock sync.Mutex
}

// 这里是一个构造函数
func NewBook(bookId,name string,num int,author string,publishDate int64)(book *Book){
	book = &Book{
		BookId:bookId,
		Name:name,
		Num:num,
		Author:author,
		PublishDate:publishDate,
	}
	return
}

// 借书
func(b *Book) Borrow()(err error){
	b.Lock.Lock()
	defer b.Lock.Unlock()
	if b.Num <= 0{
		err = errors.New("book is not enough")
		return
	}
	b.Num -= 1
	b.BorrowCount += 1
	return
}

// 还书
func (b*Book) Back()(err error){
	b.Lock.Lock()
	defer b.Lock.Lock()
	b.Num +=1
	return
}

