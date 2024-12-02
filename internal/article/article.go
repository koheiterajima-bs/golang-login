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
}

func Index(w http.ResponseWriter, r *http.Request) {
	// データベースから全ての列を取得するコマンド
	selected, err := utility.Db.Query("SELECT * FROM article")
	if err != nil {
		panic(err.Error())
	}
	// 記事データ格納用
	data := []article{}
	// データベースから取得した各行を処理
	// selected.Next()はデータベースから返された各行にアクセスするための繰り返し処理
	for selected.Next() {
		article := article{}
		// 現在の行のデータをarticleという構造体に読み込ませる
		err = selected.Scan(&article.Id, &article.Title, &article.Body)
		if err != nil {
			panic(err.Error())
		}
		data = append(data, article)
	}
	// データベースとの接続を閉じる
	selected.Close()

	// テンプレートを使ってHTMLを生成し、クライアントに返す
	if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		log.Fatal(err)
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	// URLからidを取得
	id := r.URL.Query().Get("id")
	// データベースから指定されたidの記事を取得
	selected, err := utility.Db.Query("SELECT * FROM article WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	// データベースの結果を処理
	article := article{}
	for selected.Next() {
		err = selected.Scan(&article.Id, &article.Title, &article.Body)
		if err != nil {
			panic(err.Error())
		}
	}
	// データベースとの接続を閉じる
	selected.Close()

	// テンプレートを使ってHTMLを生成し、クライアントに返す
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
		// データベースに新しい記事を挿入
		insert, err := utility.Db.Prepare("INSERT INTO article(title, body) VALUES(?, ?)")
		if err != nil {
			panic(err.Error())
		}
		insert.Exec(title, body)
		// 記事作成後、トップページへリダイレクト
		http.Redirect(w, r, "/", 301)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	// リクエストがGETの場合、編集画面の表示
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		// データベースから該当記事の取得
		selected, err := utility.Db.Query("SELECT * FROM article WHERE id = ?", id)
		if err != nil {
			panic(err.Error())
		}
		// 結果をスキャンして構造体に格納
		article := article{}
		for selected.Next() {
			err = selected.Scan(&article.Id, &article.Title, &article.Body)
			if err != nil {
				panic(err.Error())
			}
		}
		// データベースとの接続を閉じる
		selected.Close()
		// 編集画面を生成して表示
		tmpl.ExecuteTemplate(w, "edit.html", article)
		// リクエストがPOSTの場合、記事の更新
	} else if r.Method == "POST" {
		title := r.FormValue("title")
		body := r.FormValue("body")
		id := r.FormValue("id")
		// データベースの更新クエリを準備・実行
		insert, err := utility.Db.Prepare("UPDATE article SET title=?, body=? WHERE id = ?")
		if err != nil {
			panic(err.Error())
		}
		insert.Exec(title, body, id)
		// 記事更新後、トップページへリダイレクト
		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	// リクエストがGETの場合、削除確認画面の表示
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		// データベースから該当記事を取得
		selected, err := utility.Db.Query("SELECT * FROM article WHERE id = ?", id)
		if err != nil {
			panic(err.Error())
		}
		// 結果をスキャンして構造体に格納
		article := article{}
		for selected.Next() {
			err = selected.Scan(&article.Id, &article.Title, &article.Body)
			if err != nil {
				panic(err.Error())
			}
		}
		// データベースとの接続を閉じる
		selected.Close()
		tmpl.ExecuteTemplate(w, "delete.html", article)
	} else if r.Method == "POST" {
		id := r.FormValue("id")
		insert, err := utility.Db.Prepare("DELETE FROM article WHERE id = ?")
		if err != nil {
			panic(err.Error())
		}
		insert.Exec(id)
		// 記事削除後、トップページへリダイレクト
		http.Redirect(w, r, "/", 301)
	}
}
