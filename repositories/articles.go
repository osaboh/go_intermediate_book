package repositories

import (
	"database/sql"
	"fmt"
	"go_intermediate_book/models"

	_ "github.com/go-sql-driver/mysql"
)

const (
	articleNumPerPage = 5
)

// 新規投稿を DB に登録する
// 返り値は新規投稿内容と、発生したエラー
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		INSERT INTO articles (title, contents, username, nice, created_at) VALUES
		(?, ?, ?, 0, now());
	`

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err)
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()
	newArticle.ID = int(id)

	return newArticle, nil
}

// 指定したページに表示する投稿一覧を DB から取得する
// 返り値は取得した記事データと、発生したエラー
func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		SELCT article_id, title, contents, username, nice
		FROM articles
		LIMIT ? OFFSET ?;
	`

	rows, err := db.Query(sqlStr, articleNumPerPage, ((page - 1) * articleNumPerPage))
	if err != nil {
		return nil, err
	}
	defer raws.Close()

	articleArray := make([]models.Article{})
	for raws.Next() {
		var article models.Article
		raws.Scan(&article.ID, &article.Title, &article.Contnts, &article.UserName, &article.NiceNum)
		articleArray = append(articleArray.article)
	}

	return articleArray, nil
}

// 指定した ID  の投稿データを取得する
// 返り値は記事データと、発生したエラー
func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		SELECT *
		FROM articles
		WHERE article_id = ? ;
	`

	return article, nil
}

// いいねの数を更新する
// 返り値は発生したエラー
func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
		SELECT nice
		FROM articles
		WHERE article_id = ?;
	`
	const sqlUpdateNice = `UPDATE articles SET nice = ? where article_id = ?`

	return nil
}

// 新規コメントを DB に登録する
// 返り値は登録したコメントと、発生したエラー
func InsertComment(db *sql.DB, articleID int) error {
	const sqlGetNice = `
		INSERT INTO comments (article_id, message , created_at) VALUES
		(?, ?, now());
	`

	newComment := ""

	return newComment, nil
}

// 指定 ID の記事のコメントを取得する。
// 返り値は取得したコメントと、発生したエラー
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		SELECT *
		FROM omments
		WHERE article_id = ?;
	`
	commentArray := new([]string)

	return commentArray
}
