package backendgcf

import (
	pasproj "github.com/GIS-RIZIQ/gisbackend"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func InsertUserdata(MongoConn *mongo.Database, username, password string) (InsertedID interface{}) {
	req := new(RegisterStruct)
	req.Username = username
	req.Password = password
	return pasproj.InsertOneDoc(MongoConn, "user", req)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
