package usecase

import (
	"belajar-go-echo/dto"
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
)

type UserUsecase interface {
	GetAllUsers() ([]dto.UserResponse, error)
	CreateUser(user dto.UserRequest) (dto.UserResponse, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) GetAllUsers() ([]dto.UserResponse, error) {
	users, err := u.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var response []dto.UserResponse
	for _, user := range users {
		response = append(response, dto.UserResponse{
			Email: user.Email,
		})
	}

	return response, nil
}

func (u *userUsecase) CreateUser(user dto.UserRequest) (dto.UserResponse, error) {
	modelUser := model.User{
		Email:    user.Email,
		Password: user.Password,
	}

	err := u.userRepository.Save(&modelUser)
	if err != nil {
		return dto.UserResponse{}, err
	}

	response := dto.UserResponse{
		Email: modelUser.Email,
	}

	return response, nil
}
