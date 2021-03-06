package service

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/bukalapak/packen/response"
	"github.com/luqmanarifin/minisso/model"
)

const (
	TOKEN_LIFETIME = 1 * time.Hour
	COOKIE_NAME    = "minisso"
)

// Decode body to a struct
func Decode(r *http.Request, object interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	return decoder.Decode(&object)
}

// handle response both success or failed
func HandleResponse(w http.ResponseWriter, response interface{}, err string, status int) {
	w.Header().Set("Content-Type", "application/json")
	switch status {
	case 200:
		WriteSuccess(w, response, status)
		break
	case 201:
		WriteSuccess(w, response, status)
		break
	default:
		WriteError(w, err, status)
		break
	}

}

// build error response and write it
func WriteError(w http.ResponseWriter, err string, status int) {
	errCust := response.CustomError{
		Message:  err,
		HTTPCode: status,
	}
	errs := []error{errCust}
	res := response.BuildError(errs)
	response.Write(w, res, status)
}

// build success response and write it
func WriteSuccess(w http.ResponseWriter, data interface{}, status int) {
	res := response.BuildSuccess(data, response.MetaInfo{HTTPStatus: status})
	response.Write(w, res, status)
}

func ExtractCredential(r *http.Request) (credential model.Credential, token string, err error) {
	if err = Decode(r, &credential); err != nil {
		log.Printf("decoding user error")
		return
	}
	cookie, err := r.Cookie(COOKIE_NAME)
	if err == nil {
		token = cookie.Value
	}
	err = nil
	return
}

func GenerateString(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func GenerateToken(userId int64) model.Token {
	return model.Token{
		Token:     GenerateString(20),
		UserId:    userId,
		ExpiredAt: time.Now().Add(TOKEN_LIFETIME),
	}
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers",
		"X-CSRF-Token, Authorization, Origin, X-Requested-With, Content-Type, Accept, Cache-Control, Content-Length, Accept-Encoding")
}
