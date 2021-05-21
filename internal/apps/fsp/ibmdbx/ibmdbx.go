package ibmdbx

import (
	//"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/ibmdb/go_ibm_db"
	"log"
	"github.com/nekohor/mangosteen/pkg/errors"
	"sync"
)

var (
	once sync.Once
	dbInstance map[string]*sqlx.DB
)

func GetConnStringMap() map[string]string {
	m := make(map[string]string)
	m["ap"] = "DATABASE=mgfsp;HOSTNAME=170.0.35.111;PORT=50000;PROTOCOL=TCPIP;UID=ap;PWD=p@ssw0rd;"
	m["pw"] = "DATABASE=mg;HOSTNAME=170.0.35.140;PORT=50000;PROTOCOL=TCPIP;UID=pw;PWD=pwd@PW;"
	m["pd"] = "DATABASE=mg;HOSTNAME=170.0.35.140;PORT=50000;PROTOCOL=TCPIP;UID=pd;PWD=pwd@PD;"
	m["qu"] = "DATABASE=mg;HOSTNAME=170.0.35.140;PORT=50000;PROTOCOL=TCPIP;UID=qu;PWD=pwd@QU;"
	return m
}

func DB(userId string) *sqlx.DB {
	db, err := GetInstance(userId)
	if err != nil {
		log.Println(err)
	}
	return db
}

func GetInstance(userId string) (*sqlx.DB, error) {
	once.Do(func() {
		var err error
		dbInstance = make(map[string]*sqlx.DB)
		for userId, connString := range GetConnStringMap() {
			dbInstance[userId], err = sqlx.Open("go_ibm_db", connString)
			if err != nil {
				log.Println(userId)
				log.Println(err)
			}
		}
	})
	for userIdTag, _ := range dbInstance {
		if userId == userIdTag {
			return dbInstance[userId], nil
		}
	}
	return nil, errors.New("Wrong userId!")
}

