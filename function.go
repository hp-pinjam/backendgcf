package backendgcf

import (
	"github.com/whatsauth/watoken"
	"golang.org/x/crypto/bcrypt"
)

// func InsertUserdata(MongoConn *mongo.Database, username, password string) (InsertedID interface{}) {
// 	req := new(RegisterStruct)
// 	req.Username = username
// 	req.Password = password
// 	return pasproj.InsertOneDoc(MongoConn, "user", req)
// }

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsExist(Tokenstr, PublicKey string) bool {
	id := watoken.DecodeGetId(PublicKey, Tokenstr)
	if id == "" {
		return false
	}
	return true
}
