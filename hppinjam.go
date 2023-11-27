package backendgcf

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/aiteung/atdb"
	"github.com/petapedia/peda"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/mongo"
)

// func GCFHandler(MONGOCONNSTRINGENV, dbname, collectionname string) string {
// 	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
// 	datagedung := GetAllBangunanLineString(mconn, collectionname)
// 	return GCFReturnStruct(datagedung)
// }

func GCFPostHandler(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	var Response Credential
	Response.Status = false
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	var datauser User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
	} else {
		if IsPasswordValid(mconn, collectionname, datauser) {
			Response.Status = true
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(PASETOPRIVATEKEYENV))
			if err != nil {
				Response.Message = "Gagal Encode Token : " + err.Error()
			} else {
				Response.Message = "Selamat Datang"
				Response.Token = tokenstring
			}
		} else {
			Response.Message = "Password Salah"
		}
	}

	return GCFReturnStruct(Response)
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func InsertUser(db *mongo.Database, collection string, userdata User) string {
	hash, _ := HashPassword(userdata.Password)
	userdata.Password = hash
	atdb.InsertOneDoc(db, collection, userdata)
	return "Ini username : " + userdata.Username + "ini password : " + userdata.Password
}

func Register(Mongoenv, dbname string, r *http.Request) string {
	resp := new(pasproj.Credential)
	userdata := new(RegisterStruct)
	resp.Status = false
	conn := GetConnectionMongo(Mongoenv, dbname)
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		resp.Status = true
		hash, err := pasproj.HashPass(userdata.Password)
		if err != nil {
			resp.Message = "Gagal Hash Password" + err.Error()
		}
		InsertUserdata(conn, userdata.Username, hash)
		resp.Message = "Berhasil Input data"
	}
	response := pasproj.ReturnStringStruct(resp)
	return response
}

func Login(Privatekey, MongoEnv, dbname, Colname string, r *http.Request) string {
	var resp pasproj.Credential
	mconn := pasproj.MongoCreateConnection(MongoEnv, dbname)
	var datauser peda.User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else {
		if peda.IsPasswordValid(mconn, Colname, datauser) {
			tokenstring, err := watoken.Encode(datauser.Username, os.Getenv(Privatekey))
			if err != nil {
				resp.Message = "Gagal Encode Token : " + err.Error()
			} else {
				resp.Status = true
				resp.Message = "Selamat Datang"
				resp.Token = tokenstring
			}
		} else {
			resp.Message = "Password Salah"
		}
	}
	return pasproj.ReturnStringStruct(resp)
}
