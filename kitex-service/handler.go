package main

import (
	"context"
	"fmt"

	usermodel "github.com/jun-chiang/go-framework-practice/kitex-service/kitex_gen/usermodel"
)

var users []*usermodel.User

func init() {
	users = make([]*usermodel.User, 3)
	for i := 0; i < 3; i++ {
		users[i] = &usermodel.User{
			ID:        int64(i + 1),
			LoginName: fmt.Sprintf("userLoginName_%d", i+1),
			Password:  "password",
			NickName:  fmt.Sprintf("user_%d", i+1),
		}
	}
}

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// SignIn implements the UserServiceImpl interface.
func (s *UserServiceImpl) SignIn(ctx context.Context, user *usermodel.User) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// SignUp implements the UserServiceImpl interface.
func (s *UserServiceImpl) SignUp(ctx context.Context, user *usermodel.User) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// GetUserList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserList(ctx context.Context) (resp []*usermodel.User, err error) {
	// TODO: Your code here...
	return
}

// GetUserById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, id int64) (resp *usermodel.User, err error) {
	fmt.Println("接收到远程RPC调用,参数是:", id)
	return users[id], nil
}
