package repositories

import (
	"errors"
	"gorm.io/gorm"
	"metabit-service/core/config"
	"metabit-service/core/models"
)

type TokenRepository interface {
	FindAll() (error, *[]models.Token)
	FindList(chainType uint, network string) (error, *[]models.Token)
	FindById(tokenID uint64) (error, *models.Token)
	Insert(token *models.Token) (error, *models.Token)
	Update(token *models.Token) (error, *models.Token)
	Del(token *models.Token) error
}

type tokenContext struct {
	db func() *gorm.DB
}

func NewTokenRepository() TokenRepository {
	return &tokenContext{
		db: config.GetDB,
	}
}

func (ctx *tokenContext) FindAll() (error, *[]models.Token) {
	var tokens []models.Token
	tx := ctx.db().Find(&tokens)
	if tx.Error != nil {
		return tx.Error, nil
	}
	return nil, &tokens
}

func (ctx *tokenContext) FindList(chainType uint, network string) (error, *[]models.Token) {
	var tokens []models.Token
	tx := ctx.db().Find(&tokens)
	if tx.Error != nil {
		return tx.Error, nil
	}
	return nil, &tokens
}

func (ctx *tokenContext) FindById(tokenID uint64) (error, *models.Token) {
	var token models.Token
	tx := ctx.db().Find(&token, tokenID)
	if tx.Error != nil {
		return tx.Error, nil
	}
	return nil, &token
}

func (ctx *tokenContext) Insert(token *models.Token) (error, *models.Token) {
	tx := ctx.db().Create(token)
	if tx.Error != nil {
		return tx.Error, nil
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed"), nil
	}
	return nil, token
}

func (ctx *tokenContext) Update(token *models.Token) (error, *models.Token) {
	tx := ctx.db().Updates(token)
	if tx.Error != nil {
		return tx.Error, nil
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed"), nil
	}
	return nil, token
}

func (ctx *tokenContext) Del(token *models.Token) error {
	tx := ctx.db().Delete(token)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("del failed")
	}
	return nil
}
