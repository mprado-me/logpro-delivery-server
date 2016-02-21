package user

type Login struct {
	Email    string
	Password string
	Kind     string
}

func NewLogin(email, password, kind string) Login {
	login := Login{
		Email:    email,
		Password: password,
		Kind:     kind,
	}
	return login
}
