package services

import (
	"metabit-service/core/models"
	"metabit-service/core/repositories"
)

type TokenService interface {
	FindAll() (error, *[]models.Token)
	FindList() (error, *[]models.Token)
	FindById(tokenID uint64) (error, *models.Token)
	Insert(token *models.Token) (error, *models.Token)
	Update(token *models.Token) (error, *models.Token)
	Del(tokenID uint64) error
}

type tokenService struct {
	tokenRepository repositories.TokenRepository
}

func NewTokenService() TokenService {
	return &tokenService{
		tokenRepository: repositories.NewTokenRepository(),
	}
}

func (service *tokenService) FindAll() (error, *[]models.Token) {
	return service.tokenRepository.FindAll()
}

func (service *tokenService) FindList() (error, *[]models.Token) {
	//TODO implement me
	panic("implement me")
}

func (service *tokenService) FindById(tokenID uint64) (error, *models.Token) {
	return service.tokenRepository.FindById(tokenID)
}

func (service *tokenService) Insert(token *models.Token) (error, *models.Token) {
	return service.tokenRepository.Insert(token)
}

func (service *tokenService) Update(token *models.Token) (error, *models.Token) {
	//TODO implement me
	panic("implement me")
}

func (service *tokenService) Del(tokenID uint64) error {
	//TODO implement me
	panic("implement me")
}
