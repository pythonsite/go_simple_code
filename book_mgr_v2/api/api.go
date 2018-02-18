// api的主要处理逻辑
package main

import (
	"net/http"
	"go_dev/07/book_mgr_v2/logic"
	"strconv"
	"encoding/json"
	"fmt"
)


var (
	bookMgr *logic.BookMgr
	studentMgr *logic.StudentMgr
)

func init(){
	bookMgr = logic.NewBookMgr()
	studentMgr = logic.NewStudentMgr()
}

func responseError(w http.ResponseWriter,code int){
	m := make(map[string]interface{},16)
	m["code"] = code
	m["message"] = getMessage(code)
	data,err := json.Marshal(m)
	if err != nil{
		w.Write([]byte("{\"code\":500,\"message\":\"server busy\"}"))
		return
	}
	w.Write(data)
}

func responseSuccess(w http.ResponseWriter,code int,data interface{}){
	m := make(map[string]interface{},16)
	m["code"] = code
	m["message"] = getMessage(code)
	m["data"] = data


	dataByte,err := json.Marshal(m)
	if err != nil{
		w.Write([]byte("{\"code\":500,\"message\":\"server busy\"}"))
		return
	}
	w.Write(dataByte)
}

func addBook(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	bookId := r.FormValue("book_id")
	name := r.FormValue("name")
	numStr := r.FormValue("num")
	author := r.FormValue("author")
	publishDateStr := r.FormValue("publish")

	num ,err := strconv.Atoi(numStr)
	if err != nil{
		responseError(w,ErrInvalidParameter)
		return
	}

	publishDate ,err := strconv.Atoi(publishDateStr)
	if err != nil{
		responseError(w,ErrInvalidParameter)
		return
	}

	if len(name) == 0 || len(author) == 0 || len(bookId) == 0{
		responseError(w,ErrInvalidParameter)
		return
	}

	book := logic.NewBook(bookId,name,num,author,int64(publishDate))
	err = bookMgr.AddBook(book)
	if err != nil{
		responseError(w,ErrServerBusy)
		return
	}
	responseSuccess(w,ErrSuccess,nil)
}

func searchBookName(w http.ResponseWriter,r *http.Request){
	r.ParseForm()

	name := r.FormValue("name")



	if len(name) == 0 {
		responseError(w,ErrInvalidParameter)
		return
	}

	bookList := bookMgr.SearchByBookName(name)

	responseSuccess(w,ErrSuccess,bookList)
}

func searchBookAuthor(w http.ResponseWriter,r *http.Request){
	r.ParseForm()

	author := r.FormValue("author")



	if len(author) == 0 {
		responseError(w,ErrInvalidParameter)
		return
	}

	bookList := bookMgr.SearchByAuthor(author)

	responseSuccess(w,ErrSuccess,bookList)
}

func addStudent(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	idStr := r.FormValue("id")
	name := r.FormValue("name")
	gradeStr := r.FormValue("grade")
	identify := r.FormValue("identify")
	sexStr := r.FormValue("sex")

	id,err := strconv.Atoi(idStr)
	if err != nil{
		responseError(w,ErrInvalidParameter)
		return
	}
	grade,err := strconv.Atoi(gradeStr)
	if err != nil{
		responseError(w,ErrInvalidParameter)
		return
	}

	sex,err := strconv.Atoi(sexStr)
	if err != nil{
		responseError(w,ErrInvalidParameter)
		return
	}

	if len(name) == 0 || (sex != 0 && sex != 1)|| len(identify) == 0{
		responseError(w,ErrInvalidParameter)
		return
	}

	stu := logic.NewStudent(id,name,grade,identify,sex)
	err = studentMgr.AddStudent(stu)
	if err != nil{
		responseError(w,ErrServerBusy)
		return
	}
	responseSuccess(w,ErrSuccess,nil)
}

func borrow(w http.ResponseWriter,r *http.Request){
	r.ParseForm()

	sidStr := r.FormValue("sid")
	bid := r.FormValue("bid")
	sid ,err := strconv.Atoi(sidStr)
	if err != nil{
		responseError(w,ErrInvalidParameter)
		return
	}

	if len(bid) == 0{
		responseError(w,ErrInvalidParameter)
		return
	}

	student,err := studentMgr.GetStudentByid(sid)
	if err != nil{
		responseError(w,ErrInvalidParameter)
		return
	}
	err = bookMgr.Borrow(student,bid)
	if err != nil{
		responseError(w,ErrInvalidParameter)
		return
	}

	responseSuccess(w,ErrSuccess,nil)
}

func bookList(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	sidStr := r.FormValue("sid")
	sid ,err := strconv.Atoi(sidStr)
	if err != nil{
		responseError(w,ErrInvalidParameter)
		return
	}
	bookList,err := studentMgr.GetStudentBorrowBooks(sid)
	if err != nil{
		responseError(w,ErrInvalidParameter)
		return
	}

	responseSuccess(w,ErrSuccess,bookList)
}

func getTop10(w http.ResponseWriter,r *http.Request){

	bookList := bookMgr.GetTop10()

	responseSuccess(w,ErrSuccess,bookList)
}

func main(){
	http.HandleFunc("/book/add",addBook)
	http.HandleFunc("/book/searchName",searchBookName)
	http.HandleFunc("/book/searchAuthor",searchBookAuthor)
	http.HandleFunc("/student/add",addStudent)
	http.HandleFunc("/student/borrow",borrow)
	http.HandleFunc("/student/bookList",bookList)
	http.HandleFunc("/book/",getTop10)

	err := http.ListenAndServe(":8080",nil)
	if err != nil{
		fmt.Println("server start failed,err is",err)
	}
}