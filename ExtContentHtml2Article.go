package c_code

import "github.com/sundy-li/html2article"

func ExtContentHtml2Article(content string) (article *html2article.Article, err error) {
	//urlStr := "https://studyseo.net/q/5eaba3c1d9d706e065243539"
	//ext, err := html2article.NewFromUrl(urlStr)
	ext, err := html2article.NewFromHtml(content)
	if err != nil {
		return
	}
	article, err = ext.ToArticle()
	if err != nil {
		return
	}
	//println("article title is =>", article.Title)
	//println("article publishtime is =>", article.Publishtime) //using UTC timezone
	//println("article content is =>", article.Content)

	//parse the article to be readability
	//article.Readable(urlStr)
	//println("read=>", article.ReadContent)
	return
}
