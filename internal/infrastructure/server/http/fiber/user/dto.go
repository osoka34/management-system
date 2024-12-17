package user

import "service/internal/app/command/user"

type RegisterRequest struct {
	Login    string `json:"login"    required:"true"`
	Password string `json:"password" required:"true"`
}


type RegisterResponse struct {
    BearerToken string `json:"bearer_token"`
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

func (r RegisterRequest) ToCmd() *user.CreateUserCmd {
	return &user.CreateUserCmd{
		Login:    r.Login,
		Password: r.Password,
	}
}
