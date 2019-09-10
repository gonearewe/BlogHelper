package models

import (
	"path/filepath"
	"strings"
	"time"
)

type Article struct {
	title             string
	subtitle          string
	tags              []string
	createTime        string
	author            string
	backgroundPicture string
}

func NewArticle(
	title, subtitle, author, picture string,
	tags []string) *Article {

	time := time.Now().Format("2006-01-02")
	article := &Article{
		title:             title,
		subtitle:          subtitle,
		author:            author,
		createTime:        time,
		tags:              tags,
		backgroundPicture: picture,
	}
	return article
}
func (a *Article) FillTemplate(template string) string {
	r := strings.NewReplacer(
		"TITLE", a.title,
		"SUBTITLE", a.subtitle,
		"AUTHOR", a.author,
		"CREATETIME", a.createTime,
		"PICTURE", a.backgroundPicture,
	)
	tmp := r.Replace(template)
	index := strings.Index(tmp, "tags:")
	tagsString := "\n"
	for _, tag := range a.tags {
		tagsString = strings.Join([]string{tagsString, "     ", tag, "\n"}, "")
	}
	return tmp[:index+5] + tagsString + tmp[index+5:]

}

func (a *Article) GetFileAddr(filePath string) string {
	name := a.createTime + "-" + a.title + ".md"
	return filepath.Join(filePath, name)
}
