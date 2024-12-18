package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go_intermediate_book/handlers"
	"go_intermediate_book/models"
	"log"
	"net/http"
)

func insert() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,
		dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// データを挿入する処理
	article := models.Article{
		Title:    "insert test",
		Contents: "Can I insert data correctly?",
		UserName: "osaboh",
	}

	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) values
		(?, ?, ?, 0, now())
	`

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err)
		return

	}

	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

	return
}

func scan() {

	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,
		dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	articleID := 1
	const sqlStr = `
		SELECT *
		FROM articles
		WHERE article_id = ?;
	`
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return

	}

	// DB から取得したデータを変数 article に読み出す
	var article models.Article
	var createdTime sql.NullTime
	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName,
		&article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	fmt.Printf("%+v\n", article)

}

func transaction() {

	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser,
		dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 現在のいいねを取得する
	article_id := 1
	const sqlGetNice = `
		SELECT nice FROM articles WHERE article_id = ?;
	`
	row := tx.QueryRow(sqlGetNice, article_id)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	// 変数 nicenum に現在のいいねを読み込む
	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	// いいねをインクリメント

	const sqlUpdateNice = `
		UPDATE articles SET nice = ? WHERE article_id = ?
	`
	row = tx.QueryRow(sqlUpdateNice, nicenum+1, article_id)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	// コミットして確定
	tx.Commit()
}

func main() {

	insert()
	scan()
	transaction()

	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("Server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
