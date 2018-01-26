package models

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	//"encoding/binary"
	_ "github.com/lib/pq"
	"io"
	"log"
	"math"
	"math/big"
	//"strconv"
)

type User struct {
	Id       int
	UserName string
	Password string
	Salt     string
}

func NewUser(username string, password string) *User {

	user := User{
		UserName: username,
	}

	user.SetPassword(password)

	return &user
}

func (user *User) SetPassword(password string) {
	var intSalt *big.Int

	intSalt, _ = rand.Int(rand.Reader, big.NewInt(math.MaxInt32))

	hasher := sha1.New()
	io.WriteString(hasher, password)
	io.WriteString(hasher, intSalt.String())

	user.Password = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	user.Salt = intSalt.String()

}

func AuthenticateUser(username string, password string) *User {

	connString := "dbname=postgres sslmode=disable user=postgres"

	db, err := sql.Open("postgres", connString)

	if err != nil {
		log.Println(err.Error())

	}

	rows, err := db.Query("SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	for rows.Next() {
		var uName, pass string
		var id string
		err = rows.Scan(&uName, &pass, &id)

		if err != nil {
			log.Println(err.Error())
			return nil
		}

		if pass != password {
			log.Println("Invalid Password")
			return nil

		}

		user := NewUser(username, password)

		return user

	}

	log.Println("User %s does not exist", username)

	return nil
}
