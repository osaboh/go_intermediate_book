# 記事データを格納するためのテーブル
CREATE TABLE IF NOT EXISTS articles (
    article_id integer unsigned AUTO_INCREMENT PRIMARY KEY,
    title varchar(100) NOT NULL,
    contents text NOT NULL,
    username varchar(100) NOT NULL,
    nice integer NOT NULL,
    created_at DATETIME
);

# コメントデータを格納するためのテーブル
CREATE TABLE IF NOT EXISTS comments (
    comment_id integer unsigned AUTO_INCREMENT PRIMARY KEY,
    article_id integer unsigned NOT NULL,
    message text NOT NULL,
    created_at DATETIME,
    FOREIGN KEY (article_id) REFERENCES articles (article_id)
);
