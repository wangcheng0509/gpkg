package log

import (
	"errors"

	"github.com/wangcheng0509/gpkg/mysqlconn"
	"github.com/wangcheng0509/gpkg/try"
)

var databaseSetting = &mysqlconn.Database{}

func saveDBLog(log LogModel) error {
	var err interface{}
	try.Try(func() {
		defer func() {
			if err = recover(); err != nil {
				try.Throw(1, err.(string))
			}
		}()
		db.Create(&log)
	}).Catch(1, func(e try.Exception) {
		err = errors.New(e.Msg)
	}).Finally(func() {})

	return nil
}

func CloseDB() {
	defer db.Close()
}
