package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type Bookdetail struct {
	BookName  string `gorm:"column:book_name;type:varchar(255) CHARACTER SET utf8mb4"`
	Author    string `gorm:"column:author;type:varchar(255) CHARACTER SET utf8mb4"`
	Publicer  string `gorm:"column:publicer;type:varchar(255) CHARACTER SET utf8mb4"`
	Bookpages int    `gorm:"column:book_pages;type:varchar(255) CHARACTER SET utf8mb4"`
	Price     string `gorm:"column:price;type:varchar(255) CHARACTER SET utf8mb4"`
	Score     string `gorm:"column:score;type:varchar(255) CHARACTER SET utf8mb4"`
	Into      string `gorm:"column:into;type:varchar(255) CHARACTER SET utf8mb4"`
}

func (b Bookdetail) String() string {
	return "书籍名字:" + b.BookName + " 作者 :" + b.Author + " 出版社" + b.Publicer + " 书籍页数：" + strconv.Itoa(b.Bookpages) + " 价格：" + b.Price + " 得分" + b.Score + " \n简介:" + b.Into
}

func (n Bookdetail) TableName() string {
	return "book_detail"
}

type MysqlDB struct {
	dbEngine *gorm.DB
}

func OpenDB() (*MysqlDB, error) {
	db := &MysqlDB{}
	dsn := "zjx:123456@tcp(127.0.0.1:3326)/book?charset=utf8"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.dbEngine = d
	return db, nil
}

func (p *MysqlDB) InitTable(dst interface{}) error {
	if !p.dbEngine.Migrator().HasTable(dst) {
		return p.dbEngine.Migrator().CreateTable(dst)
	}

	return nil
}

func (p *MysqlDB) Insert(value interface{}) error {
	tx := p.dbEngine.Create(value)
	return tx.Error
}
