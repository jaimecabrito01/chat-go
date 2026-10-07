package user

import "errors"

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}
func (s *Service) Create(username string) (*User, error) {
	if username == "" {
		return nil, errors.New("username is empty")
	}
	existing, err := s.repository.FindByUsername(username)
	if err != nil || existing != nil {
		return nil, errors.New("username already exists")
	}

	user := &User{
		Username: username,
	}
	if err := s.repository.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}
