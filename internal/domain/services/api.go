package api

type UserService interface {
	GetUser(ctx context.Context, id int) (model.User, error)
	CreateNewUser(ctx context.Context, model.User) (error)
	UpdateUser(ctx context.Context, model.User) (error)
	DeleteUser(ctx context.Context, id int) (error)
}