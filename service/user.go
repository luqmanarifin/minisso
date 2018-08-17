package service

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/luqmanarifin/minisso/database"
	"github.com/luqmanarifin/minisso/model"
)

type UserService struct {
	mysql database.Mysql
	redis database.Redis
}

func NewUserService(mysqlOpt database.MysqlOption, redisOpt database.RedisOption) UserService {
	mysql, err := database.NewXorm(mysqlOpt)
	if err != nil {
		log.Fatal("cant connect to mysql")
	}
	redis, err := database.NewRedis(redisOpt)
	return UserService{
		mysql: mysql,
		redis: redis,
	}
}

func (u *UserService) Cookie(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cookie, err := r.Cookie("kentang")
	if err != nil {
		log.Printf("ambil cookie error")
	} else {
		log.Printf("cookie %s: %s", cookie.Name, cookie.Value)
	}
	setCookie := http.Cookie{Name: "luqman", Value: "ganteng"}
	http.SetCookie(w, &setCookie)
	HandleResponse(w, nil, "ok", 200)
}

func (u *UserService) Signup(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	EnableCors(&w)
	credential, tokenString, err := ExtractCredential(r)
	user := credential.User
	user.Role = "user"

	log.Printf("%s %s (%s) want to signup.\n", user.FirstName, user.LastName, user.Email)

	if err != nil {
		HandleResponse(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	// check if app is authorized
	if !u.isAppAuthorized(credential.Application) {
		HandleResponse(w, nil, "App is not authorized.", 401)
		return
	}

	// if there is cookie and valid
	if u.redis.IsTokenValid(tokenString) {
		u.handleTokenValid(w, tokenString)
		return
	}

	// if email already available, 401 reject
	if u.mysql.IsEmailExist(user.Email) {
		HandleResponse(w, nil, "Email already exists.", 401)
		return
	}

	// 200 created and give token
	u.mysql.CreateUser(user)
	u.handleLoggedIn(w, r, u.mysql.FindUserByEmail(user.Email))
}

func (u *UserService) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	EnableCors(&w)
	credential, tokenString, err := ExtractCredential(r)
	user := credential.User

	log.Printf("%s:%s want to login.\n", user.Email, user.Password)

	if err != nil {
		HandleResponse(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	// check if app is authorized
	if !u.isAppAuthorized(credential.Application) {
		HandleResponse(w, nil, "App is not authorized.", 401)
		return
	}

	// if there is cookie and valid, 200 give current token
	if u.redis.IsTokenValid(tokenString) {
		u.handleTokenValid(w, tokenString)
		return
	}

	// if user pass invalid, 401 wrong password
	if correctUser := u.mysql.FindUserByEmail(user.Email); correctUser.Password != user.Password {
		u.handleWrongPassword(w, credential)
		return
	}

	// 200 ok give token
	user = u.mysql.FindUserByEmail(user.Email)
	u.handleLoggedIn(w, r, user)
}

func (u *UserService) Validate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	EnableCors(&w)
	credential, tokenString, err := ExtractCredential(r)

	log.Printf("Token %s want to be validated.\n", tokenString)

	if err != nil {
		HandleResponse(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	// check if app is authorized
	if !u.isAppAuthorized(credential.Application) {
		HandleResponse(w, nil, "App is not authorized.", 401)
		return
	}

	// if there is cookie and valid, 200 ok give user info
	if u.redis.IsTokenValid(tokenString) {
		u.handleTokenValid(w, tokenString)
		return
	}

	// 401 not logged-in
	u.handleTokenInvalid(w)
}

func (u *UserService) handleLoggedIn(w http.ResponseWriter, r *http.Request, user model.User) {
	user = u.mysql.TouchUserLogin(r, user)
	token := u.mysql.CreateToken(GenerateToken(user.Id))
	u.mysql.CreateLogin(model.Login{
		UserId: user.Id,
		Token:  token.Token,
		Status: "success",
	})
	u.redis.AddToken(token.Token, user.Id, TOKEN_LIFETIME)
	http.SetCookie(w, &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   token.Token,
		Expires: time.Now().Add(TOKEN_LIFETIME),
	})
	HandleResponse(w, user, "ok", 200)
}

func (u *UserService) handleTokenValid(w http.ResponseWriter, tokenString string) {
	userId := u.redis.GetUserId(tokenString)
	user := u.mysql.FindUserById(userId)
	http.SetCookie(w, &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   tokenString,
		Expires: time.Now().Add(TOKEN_LIFETIME),
		MaxAge:  int(TOKEN_LIFETIME),
	})
	HandleResponse(w, user, "ok", 200)
}

func (u *UserService) handleTokenInvalid(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   COOKIE_NAME,
		MaxAge: -1,
	})
	HandleResponse(w, nil, "Token invalid", 401)
}

func (u *UserService) handleWrongPassword(w http.ResponseWriter, credential model.Credential) {
	u.mysql.CreateLogin(model.Login{
		UserId:        credential.User.Id,
		ApplicationId: credential.Application.Id,
		Status:        "failed wrong password",
	})
	http.SetCookie(w, &http.Cookie{
		Name:   COOKIE_NAME,
		MaxAge: -1,
	})
	HandleResponse(w, nil, "Wrong email/password", 401)
}

func (u *UserService) isAppAuthorized(app model.Application) bool {
	realApp := u.mysql.FindApplicationByClientId(app.ClientId)
	return (realApp.ClientSecret == app.ClientSecret)
}
