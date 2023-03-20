package repositories

import (
	"errors"
	"gorm.io/gorm"
	"metabit-service/core/config"
	"metabit-service/core/models"
)

type NetworkRepository interface {
	FindAll() (error, *[]models.Network)
	FindList(chainType uint, network string) (error, *[]models.Network)
	FindById(networkID uint64) (error, *models.Network)
	Insert(network *models.Network) (error, *models.Network)
	Update(network *models.Network) (error, *models.Network)
	Del(network *models.Network) error
}

type networkContext struct {
	db func() *gorm.DB
}

func NewNetworkRepository() NetworkRepository {
	return &networkContext{
		db: config.GetDB,
	}
}

func (ctx *networkContext) FindAll() (error, *[]models.Network) {
	var networks []models.Network
	tx := ctx.db().Find(&networks)
	if tx.Error != nil {
		return tx.Error, nil
	}
	return nil, &networks
}

func (ctx *networkContext) FindList(chainType uint, network string) (error, *[]models.Network) {
	var networks []models.Network
	tx := ctx.db().Find(&networks)
	if tx.Error != nil {
		return tx.Error, nil
	}
	return nil, &networks
}

func (ctx *networkContext) FindById(networkID uint64) (error, *models.Network) {
	var network models.Network
	tx := ctx.db().Find(&network, networkID)
	if tx.Error != nil {
		return tx.Error, nil
	}
	return nil, &network
}

func (ctx *networkContext) Insert(network *models.Network) (error, *models.Network) {
	tx := ctx.db().Create(network)
	if tx.Error != nil {
		return tx.Error, nil
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed"), nil
	}
	return nil, network
}

func (ctx *networkContext) Update(network *models.Network) (error, *models.Network) {
	tx := ctx.db().Updates(network)
	if tx.Error != nil {
		return tx.Error, nil
	}
	if tx.RowsAffected == 0 {
		return errors.New("update failed"), nil
	}
	return nil, network
}

func (ctx *networkContext) Del(network *models.Network) error {
	tx := ctx.db().Delete(network)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("del failed")
	}
	return nil
}
