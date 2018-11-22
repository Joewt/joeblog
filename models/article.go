package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

func QueryByKeyArt(key string)(string,bool){
	Art := Article{Key: "key"}
	err := DB.Read(&Art,"Key")
	if err == orm.ErrNoRows {
		return "",false
	}
	return Art.Key, true
}

func InsertArticle(key string, title string, content string,uid uint, summary string)(error){

	art := Article{
		Key: key,
		User_id: uid,
		Title: title,
		Summary: summary,
		Content: content,
	}

	Art_id, err := DB.Insert(&art)
	if err != nil {
		return err
	}

	logs.Info("插入了一篇文章", Art_id)

	return nil
}


func QueryArtByPage(title string,page, limit int64)( []*Article,int64, error) {
	var art []*Article
	article := new(Article)
	qs := DB.QueryTable(article)
	num, err := qs.Filter("title__icontains", title).Limit(limit, (page-1)*limit).OrderBy("-Id").All(&art)
	if err != nil {
		return nil,0,err
	}
	return art,num, nil
}

func QueryArtCount(title string)(int64, error){
	cnt, err := DB.QueryTable("article").Filter("title__icontains", title).Count()
	if err != nil {
		logs.Info(err)
	}

	return cnt, nil
}