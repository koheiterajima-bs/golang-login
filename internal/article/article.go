package article

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/koheiterajima-bs/golang-crud/internal/utility"
)

type article struct {
	Id    int
	Title string
	Body  string
}

// テンプレートの初期化
var tmpl *template.Template

func init() {
	funcMap := template.FuncMap{
		// nl2br：HTML内で改行を<br />タグに変換するカスタム関数
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br />", -1))
		},
	}

	// ParseGlob関数で、web/template/*に格納されたテンプレートファイルを全て読み込み、tmplに格納する
	tmpl, _ = template.New("article").Funcs(funcMap).ParseGlob("web/template/*")

	// GORMを使ってデータベース接続を設定し、AutoMigrateでArticleテーブルをデータベースに作成または更新する
	utility.Db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(article{})
}

func Index(w http.ResponseWriter, r *http.Request) {
	var allArticles []article
	// GORMを使って、articleテーブルから全ての記事を取得する
	utility.Db.Find(&allArticles)
	// index.htmlテンプレートを使って、取得した記事データを表示する
	if err := tmpl.ExecuteTemplate(w, "index.html", allArticles); err != nil {
		log.Fatal(err)
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	// URLからidを取得
	id := r.URL.Query().Get("id")
	var article article
	// articleテーブルから指定されたIDの記事を1件取得
	utility.Db.First(&article, id)
	// show.htmlテンプレートを使って、取得した記事を表示する
	tmpl.ExecuteTemplate(w, "show.html", article)
}

func Create(w http.ResponseWriter, r *http.Request) {
	// リクエストがGETメソッドの場合、create.htmlを表示
	if r.Method == "GET" {
		tmpl.ExecuteTemplate(w, "create.html", nil)
		// リクエストがPOSTメソッドの場合、フォームから送信されたデータを受け取る
	} else if r.Method == "POST" {
		title := r.FormValue("title")
		body := r.FormValue("body")
		article := article{Title: title, Body: body}
		utility.Db.Create(&article)
		http.Redirect(w, r, "/", 301)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	// リクエストがGETの場合、編集画面の表示
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		var article article
		utility.Db.First(&article, id)
		tmpl.ExecuteTemplate(w, "edit.html", article)
		// リクエストがPOSTの場合、記事の更新
	} else if r.Method == "POST" {
		title := r.FormValue("title")
		body := r.FormValue("body")
		id := r.FormValue("id")

		var article article
		utility.Db.First(&article, id)
		article.Title = title
		article.Body = body
		utility.Db.Save(&article)
		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	// リクエストがGETの場合、削除確認画面の表示
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		var article article
		utility.Db.First(&article, id)
		tmpl.ExecuteTemplate(w, "delete.html", article)
	} else if r.Method == "POST" {
		id := r.FormValue("id")
		utility.Db.Delete(&article{}, id)
		http.Redirect(w, r, "/", 301)
	}
}
