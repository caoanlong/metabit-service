package services

import (
	"metabit-service/api/models"
	"metabit-service/api/repositories"
)

type TokenService interface {
	FindAll() []models.Token
	FindList() []models.Token
	FindById(tokenID uint64) models.Token
	Insert(token models.Token) models.Token
	Update(token models.Token) models.Token
	Del(tokenID uint64)
}

type tokenService struct {
	tokenRepository repositories.TokenRepository
}

func NewTokenService() TokenService {
	return &tokenService{
		tokenRepository: repositories.NewTokenRepository(),
	}
}

func (service *tokenService) FindAll() []models.Token {
	return service.tokenRepository.FindAll()
}

func (service *tokenService) FindList() []models.Token {
	//TODO implement me
	panic("implement me")
}

func (service *tokenService) FindById(tokenID uint64) models.Token {
	return service.tokenRepository.FindById(tokenID)
}

func (service *tokenService) Insert(token models.Token) models.Token {
	return service.tokenRepository.Insert(token)
}

func (service *tokenService) Update(token models.Token) models.Token {
	//TODO implement me
	panic("implement me")
}

func (service *tokenService) Del(tokenID uint64) {
	//TODO implement me
	panic("implement me")
}
