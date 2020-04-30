package log

import (
	"time"

	"github.com/wangcheng0509/gpkg/mysqlconn"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type LogModel struct {
	Id          int       `gorm:"primary_key;AUTO_INCREMENT;column:Id"`
	Application string    `gorm:"column:Application"`
	ClassName   string    `gorm:"column:ClassName"`
	Message     string    `gorm:"column:Message"`
	StackTrace  string    `gorm:"column:StackTrace"`
	Level       int       `gorm:"column:Level"`
	CreatedDate time.Time `gorm:"column:CreatedDate"`
}

func (LogModel) TableName() string {
	return "Applicationlog"
}

type Loglevel int

var (
	db      *gorm.DB
	logtype int
	_dblog  = 1
)

const (
	LogDEBUG Loglevel = iota
	LogINFO
	LogERROR
	LogWARN
	LogFATAL
)

func InitDBLog(databaseSetting mysqlconn.Database) {
	db = mysqlconn.GetDB(databaseSetting)
	logtype = _dblog
}

func Info(log LogModel) error {
	var err error
	switch logtype {
	case _dblog:
		err = saveDBLog(log)
		break
	}
	return err
}
