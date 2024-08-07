package service

import (
	"context"
	"errors"

	"rpc_gen/kitex_gen/user"
	"user/biz/dal/mysql"
	"user/biz/model"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.p
	if req.Password != req.ConfirmPassword {
		err = errors.New("password must be the same as ConfirmPassword")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(hashedPassword),
	}
	if err = model.Create(mysql.DB, s.ctx, newUser); err != nil {
		return
	}

	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
