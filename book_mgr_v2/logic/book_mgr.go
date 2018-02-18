// 书籍管理的处理逻辑

package logic

import (
	"sync"
	"fmt"
	"sort"
)

type BookMgr struct {
	BookList []*Book
	// 存储bookid 到借书学生的列表
	BookStudentMap map[string][]*Student
	// 书籍名字到书籍的索引
	BookNameMap map[string][]*Book
	// 书籍作者到书籍的索引
	BookAuthorMap map[string][]*Book
	Lock sync.Mutex

}

// 这里是一个构造函数
func NewBookMgr()(bookmgr*BookMgr){
	bookmgr = &BookMgr{
		BookStudentMap:make(map[string][]*Student,32),
		BookNameMap:make(map[string][]*Book),
		BookAuthorMap:make(map[string][]*Book,32),
	}
	return
}

// 添加书籍
func (b *BookMgr) AddBook(book *Book)(err error){
	b.Lock.Lock()
	defer b.Lock.Unlock()
	// 添加到book列表中
	b.BookList = append(b.BookList,book)
	// 更新书籍名字到同一个书籍名字对应的book列表
	bookList,ok := b.BookNameMap[book.Name]
	if !ok{
		var tmp []*Book
		tmp = append(tmp,book)
		bookList = tmp
	}else{
		bookList = append(bookList,book)
	}
	b.BookNameMap[book.Name] = bookList

	// 更新数据作者到同一个作者对应的book列表
	bookList,ok = b.BookAuthorMap[book.Author]
	if !ok{
		var tmp []*Book
		tmp = append(tmp,book)
		bookList = tmp
	}else{
		bookList = append(bookList,book)
	}
	b.BookAuthorMap[book.Author] = bookList


	return
}

// 通过书籍名字查找
func (b *BookMgr) SearchByBookName(bookName string)(bookList []*Book){
	b.Lock.Lock()
	b.Lock.Unlock()
	bookList = b.BookNameMap[bookName]
	return
}

// 通过书籍作者查找
func (b *BookMgr) SearchByAuthor(Author string)(bookList []*Book){
	b.Lock.Lock()
	b.Lock.Unlock()
	bookList = b.BookAuthorMap[Author]
	return
}

// 通过书籍出版日期查找，这里的参数是一个范围
func (b *BookMgr) SearchByPushlish(min int64,max int64)(bookList []*Book){
	b.Lock.Lock()
	b.Lock.Unlock()
	for _,v:= range b.BookList{
		if v.PublishDate >= min && v.PublishDate <= max{
			bookList = append(bookList,v)
		}
	}
	return
}

// 用于学生借书
func (b *BookMgr) Borrow(student *Student,bookId string) (err error){
	b.Lock.Lock()
	defer b.Lock.Unlock()
	var book *Book
	for _,v := range b.BookList{
		if v.BookId == bookId{
			book = v
			break
		}
	}
	if book == nil{
		err = fmt.Errorf("book id [%d] is not exist",bookId)
		return
	}
	err = book.Borrow()
	if err != nil{
		return
	}
	student.AddBook(book)
	return
}

// 这里实现了Len() Less() Swap()后就实现sort中的Interface接口，就可以调用sort.Sort方法
func (b *BookMgr) Len() int{
	return len(b.BookList)
}

// 如果是>是从大到小排序，如果是< 则是从小到大排序
func (b *BookMgr) Less(i,j int) bool{
	return b.BookList[i].BorrowCount > b.BookList[j].BorrowCount
}

func (b *BookMgr) Swap(i,j int){
	b.BookList[i],b.BookList[j]= b.BookList[j],b.BookList[i]
}

// 获取top10
func (b *BookMgr) GetTop10()(bookList []*Book){
	sort.Sort(b)
	for i :=0;i<10;i++{
		if i >= len(b.BookList){
			break
		}
		bookList = append(bookList,b.BookList[i])

	}
	return
}

