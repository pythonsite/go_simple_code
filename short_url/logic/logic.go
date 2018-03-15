package logic

import(
	"go_dev/11/short_url/model"
	"github.com/jmoiron/sqlx"
	"fmt"
	"crypto/md5"
	"database/sql"
)

var (
	Db *sqlx.DB
)

type ShortUrl struct {
	Id int64 `db:"id"`
	ShortUrl string `db:"short_url"`
	OriginUrl string `db:"origin_url"`
	HashCode string `db:"hash_code"`
}

func InitDb(dsn string)(err error) {
	// 数据库初始化
	Db, err = sqlx.Open("mysql",dsn)
	if err != nil{
		fmt.Println("connect to mysql failed:",err)
		return
	}
	return
}

func Long2Short(req *model.Long2ShortRequest) (response *model.Long2ShortResponse, err error) {
	response = &model.Long2ShortResponse{}
	urlMd5 := fmt.Sprintf("%x",md5.Sum([]byte(req.OriginUrl)))
	var short ShortUrl
	err = Db.Get(&short,"select id,short_url,origin_url,hash_code from short_url where hash_code=?",urlMd5)
	if err == sql.ErrNoRows{
		err = nil
		// 数据库中没有记录，重新生成一个新的短url
		shortUrl,errRet := generateShortUrl(req,urlMd5)
		if errRet != nil{
			err = errRet
			return
		}
		response.ShortUrl = shortUrl
		return
	}
	if err != nil{
		return
	}
	response.ShortUrl = short.ShortUrl
	return
}

func generateShortUrl(req *model.Long2ShortRequest,hashcode string)(shortUrl string,err error){
	result,err := Db.Exec("insert INTO short_url(origin_url,hash_code)VALUES (?,?)",req.OriginUrl,hashcode)
	if err != nil{
		return
	}
	// 0-9a-zA-Z 六十二进制
	insertId,_:= result.LastInsertId()
	shortUrl = transTo62(insertId)
	_,err = Db.Exec("update short_url set short_url=? where id=?",shortUrl,insertId)
	if err != nil{
		fmt.Println(err)
		return
	}
	return
}

// 将十进制转换为62进制   0-9a-zA-Z 六十二进制
func transTo62(id int64)string{
	// 1 -- > 1
	// 10-- > a
	// 61-- > Z
	charset := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var shortUrl []byte
	for{
		var result byte
		number := id % 62
		result = charset[number]
		var tmp []byte
		tmp = append(tmp,result)
		shortUrl = append(tmp,shortUrl...)
		id = id / 62
		if id == 0{
			break
		}
	}
	fmt.Println(string(shortUrl))
	return string(shortUrl)
}


func Short2Long(req *model.Short2LongRequest) (response *model.Short2LongResponse, err error) {
	response = &model.Short2LongResponse{}
	var short ShortUrl
	err = Db.Get(&short,"select id,short_url,origin_url,hash_code from short_url where short_url=?",req.ShortUrl)
	if err == sql.ErrNoRows{
		response.Code = 404
		return
	}
	if err != nil{
		response.Code = 500
		return
	}
	response.OriginUrl = short.OriginUrl
	return
}