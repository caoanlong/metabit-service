package repositories

import (
	"gorm.io/gorm"
	"metabit-service/core/config"
	"metabit-service/core/models"
)

type TokenRepository interface {
	FindAll() []models.Token
	FindList(chainType string, network string) []models.Token
	FindById(tokenID uint64) models.Token
	Insert(token models.Token) models.Token
	Update(token models.Token) models.Token
	Del(token models.Token)
}

type tokenConnection struct {
	connection *gorm.DB
}

func NewTokenRepository() TokenRepository {
	return &tokenConnection{
		connection: config.G_DB,
	}
}

func (db *tokenConnection) FindAll() []models.Token {
	var tokens []models.Token
	db.connection.Find(&tokens)
	return tokens
}

func (db *tokenConnection) FindList(chainType string, network string) []models.Token {
	var tokens []models.Token
	db.connection.Find(&tokens)
	return tokens
}

func (db *tokenConnection) FindById(tokenID uint64) models.Token {
	var token models.Token
	db.connection.Find(&token, tokenID)
	return token
}

func (db *tokenConnection) Insert(token models.Token) models.Token {
	db.connection.Save(&token)
	return token
}

func (db *tokenConnection) Update(token models.Token) models.Token {
	db.connection.Save(&token)
	return token
}

func (db *tokenConnection) Del(token models.Token) {
	db.connection.Delete(token)
}
