package main

import (
	"context"
	usermodel "github.com/jun-chiang/go-framework-practice/kitex-client/kitex_gen/usermodel"
)

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
	// TODO: Your code here...
	return
}
