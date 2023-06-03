package interfaces

type Authusecase interface {
	//Register (ctx context.Context,user domain.User)error
	HashPassword(password string)string

}
