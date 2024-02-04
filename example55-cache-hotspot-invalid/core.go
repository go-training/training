package main

import "time"

type Middleware interface {
	GetArticle(req int, id int) *Article
	GetArticleDo(req int, id int) *Article
	GetArticleDoChan(req int, id int, t time.Duration) *Article
}
