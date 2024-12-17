package user

import "service/internal/app/command/user"

type Register struct {
	Login    string `json:"login"    required:"true"`
	Password string `json:"password" required:"true"`
}

type Login struct {
	Login    string `json:"login"    required:"true"`
	Password string `json:"password" required:"true"`
}

func (l Login) ToCmd() *user.GetUserCmd {
	return &user.GetUserCmd{
		Login:    l.Login,
		Password: l.Password,
	}
}

func (r Register) ToCmd() *user.CreateUserCmd {
	return &user.CreateUserCmd{
		Login:    r.Login,
		Password: r.Password,
	}
}
