package controllers

import (
	"github.com/astaxie/beego"
	"beego_short_url/models"
	"encoding/json"
	"database/sql"
	"crypto/md5"
	"github.com/jmoiron/sqlx"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sqlx.DB
)

func InitDb()(err error){
	Db, err = sqlx.Open("mysql",beego.AppConfig.String("Db::dsn"))
	if err != nil{
		beego.Error("connect to mysql failed:",err)
		return
	}
	return
}

type ShortUrl struct {
	Id int64 `db:"id"`
	ShortUrl string `db:"short_url"`
	OriginUrl string `db:"origin_url"`
	HashCode string `db:"hash_code"`
}

type ShortUrlController struct {
	beego.Controller
}


func (c *ShortUrlController) Jump() {
	shortUrl := c.GetString("shorturl")
	if len(shortUrl) == 0{
		return
	}
	var req models.Short2LongRequest
	var resp *models.Short2LongResponse = &models.Short2LongResponse{}

	defer func(){
		if err := recover();err != nil{
			beego.Error("panic err:",err)
			//resp.Code = 500
			//resp.Message = "server busy"
			//c.Data["json"] = resp
			//c.ServeJSON()
			return
		}
	}()
	req.ShortUrl = shortUrl
	resp,err := Short2Long(&req)
	if err != nil{
		beego.Error("short2Long failed error:",err)
		return
	}

	beego.Info("origin url:%s short url:%s",resp.OriginUrl,shortUrl)
	c.Redirect(resp.OriginUrl,301)
}

func (c *ShortUrlController) ShortUrlList() {
	limit,err := c.GetInt("limit")
	if err != nil{
		beego.Warn("not have limit params use default 10")
		limit = 10
	}
	data,err := GetLastShortUrl(limit)
	if err != nil{
		beego.Error("from db get url list error:",err)

	}

	for i,v:= range data{
		v.ShortUrl = fmt.Sprintf("/jump/?shorturl=%s",v.ShortUrl)
		data[i] = v
	}

	c.Data["url_list"] = data
	c.TplName = "index.tpl"
}

func(c *ShortUrlController) Long2Short(){
	var req models.Long2ShortRequest
	var resp *models.Long2ShortResponse = &models.Long2ShortResponse{}

	defer func(){
		if err := recover();err != nil{
			beego.Error("panic err:",err)
			resp.Code = 500
			resp.Message = "server busy"
			c.Data["json"] = resp
			c.ServeJSON()
			return
		}
	}()


	err := json.Unmarshal(c.Ctx.Input.RequestBody,&req)
	if err != nil{
		beego.Error("unmarshal failed,err:",err)
		resp.Code = 1001
		resp.Message = "json unmarshal failed"
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}
	resp,err = Long2Short(&req)
	if err != nil{
		beego.Error("long2short failed,err:",err)
		resp.Code = 1002
		resp.Message = "long2short failed"
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}
	c.Data["json"] = resp
	c.ServeJSON()



}

func(c *ShortUrlController) Short2Long(){
	var req models.Short2LongRequest
	var resp *models.Short2LongResponse = &models.Short2LongResponse{}

	defer func(){
		if err := recover();err != nil{
			beego.Error("panic err:",err)
			resp.Code = 500
			resp.Message = "server busy"
			c.Data["json"] = resp
			c.ServeJSON()
			return
		}
	}()


	err := json.Unmarshal(c.Ctx.Input.RequestBody,&req)
	if err != nil{
		beego.Error("unmarshal failed,err:",err)
		resp.Code = 1001
		resp.Message = "json unmarshal failed"
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}
	resp,err = Short2Long(&req)
	if err != nil{
		beego.Error("Short2Long failed,err:",err)
		resp.Code = 1002
		resp.Message = "long2short failed"
		c.Data["json"] = resp
		c.ServeJSON()
		return
	}
	c.Data["json"] = resp
	c.ServeJSON()
}



func Long2Short(req *models.Long2ShortRequest) (response *models.Long2ShortResponse, err error) {
	response = &models.Long2ShortResponse{}
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
		response.ShortUrl = shortUrlt
		return
	}
	if err != nil{
		return
	}
	response.ShortUrl = short.ShortUrl
	return
}

func generateShortUrl(req *models.Long2ShortRequest,hashcode string)(shortUrl string,err error){
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


func Short2Long(req *models.Short2LongRequest) (response *models.Short2LongResponse, err error) {
	response = &models.Short2LongResponse{}
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


func GetLastShortUrl(limit int)(result []*models.ShortUrl,err error){
	err = Db.Select(&result,"select short_url from short_url ORDER BY id DESC  limit ? ",limit)
	return
}