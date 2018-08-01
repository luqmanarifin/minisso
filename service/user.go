package service

import (
	"log"
	"net/http"

	"github.com/bukalapak/product-mysql-command/handler/util"
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
	user := model.User{}
	if err := util.Decode(r, user); err != nil {
		HandleResponse(w, nil, err.Error(), http.StatusInternalServerError)
		return
	}

	// if there is cookie and valid, 200 ok already loggedin

	// if email already available, 200 ok reject
	if u.mysql.IsEmailExist(user.Email) {
		HandleResponse(w, nil, "Email already exists.", 200)
	}

	// 201 created and give token
	u.mysql.CreateUser(user)

}

func (u *UserService) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// if there is cookie and valid, 200 ok already loggedin

	// if user pass invalid, 200 wrong password

	// 200 ok give token
}

func (u *UserService) Validate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// if there is cookie and valid, 200 ok give user info

	// 200 wrong password
}
