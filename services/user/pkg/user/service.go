package user

type UserModel struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService interface {
	CreateUser(user *UserModel) (string, error)
	GetUser(id string) (*UserModel, error)
	UpdateUser(user *UserModel) (string, error)
	DeleteUser(id string) error
}

type userService struct{}

func NewUserService() UserService { return &userService{} }

func (s *userService) CreateUser(user *UserModel) (string, error) {
	return user.ID, nil
}

func (s *userService) GetUser(id string) (*UserModel, error) {

	return nil, nil
}

func (s *userService) UpdateUser(user *UserModel) (string, error) {
	return user.ID, nil
}

func (s *userService) DeleteUser(id string) error {
	return nil
}
