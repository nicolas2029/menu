package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbUser struct {
	TypeDB   string `json:"type_db"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     string `json:"port"`
	NameDB   string `json:"name_db"`
}

var (
	db   *gorm.DB
	once sync.Once
)

// Drivers
const (
	Postgres string = "POSTGRES"
)

// New create the connection with db
func New(file string) {
	once.Do(func() {
		u := loadFileDB(file)
		switch u.TypeDB {
		case Postgres:
			newPostgresDB(&u)
		}
	})
}

func loadFileDB(file string) dbUser {
	var err error
	m, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("no se pudo cargar las credenciales de DB: %v", err)
	}
	u := dbUser{}
	err = json.Unmarshal(m, &u)
	if err != nil {
		log.Fatalf("fallo en unmarshal DB: %v", err)
	}
	return u
}

// newPostgresDB
func newPostgresDB(u *dbUser) {
	var err error
	dsn := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", u.User, u.Password, u.Port, u.NameDB)
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("no se pudo abrir la base de datos: %v", err)
	}

	fmt.Println("conectado a postgres")
}

// DB return a unique instance of db
func DB() *gorm.DB {
	return db
}
