package app

import "github.com/nekohor/mangosteen/internal/app/dbx"

func InitDB() (func(), error) {
	dbCleanFuncMap, err := dbx.InitGormDB()
	if err != nil {
		return nil, err
	}
	return func() {
		for _, cleanFunc := range dbCleanFuncMap {
			cleanFunc()
		}
	}, err
}
