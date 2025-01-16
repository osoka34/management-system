package user

import (
	"service/internal/app/command/user"
	"service/internal/domain/entity"
	"time"
)

type RegisterRequest struct {
	Login    string `json:"login"    required:"true"`
	Password string `json:"password" required:"true"`
}

type RegisterResponse struct {
	BearerToken  string `json:"bearer_token"`
	RefreshToken string `json:"refresh_token"`
	UserId       string `json:"user_id"`
}

type Login struct {
	Login    string `json:"login"    required:"true"`
	Password string `json:"password" required:"true"`
}

type LoginResponse struct {
	BearerToken  string `json:"bearer_token"`
	RefreshToken string `json:"refresh_token"`
	UserId       string `json:"user_id"`
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

type GetAllUsersRequest struct {
}

func (r GetAllUsersRequest) ToCmd() *user.GetAllCommand {
	return &user.GetAllCommand{}
}

type User struct {
	Id        string     `json:"id"`
	Login     string     `json:"login"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func FromEntityUser(u *entity.User) *User {
	return &User{
		Id:        u.Id.String(),
		Login:     u.Login,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type GetAllUsersResponse struct {
	Users []*User
}

func NewGetAllUsersResponse(users []*entity.User) *GetAllUsersResponse {

	out := make([]*User, 0, len(users))

	for _, u := range users {
		out = append(out, FromEntityUser(u))
	}

	return &GetAllUsersResponse{Users: out}
}
