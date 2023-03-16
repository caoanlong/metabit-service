package repositories

import (
	"gorm.io/gorm"
	"metabit-service/core/config"
	"metabit-service/core/models"
)

type NetworkRepository interface {
	FindAll() []models.Network
	FindList(shortName string, networkType string) []models.Network
	FindById(networkID uint64) models.Network
	Insert(network models.Network) models.Network
	Update(network models.Network) models.Network
	Del(network models.Network)
}

type networkConnection struct {
	connection *gorm.DB
}

func NewNetworkRepository() NetworkRepository {
	return &networkConnection{
		connection: config.GDb,
	}
}

func (db *networkConnection) FindAll() []models.Network {
	var networks []models.Network
	db.connection.Find(&networks)
	return networks
}

func (db *networkConnection) FindList(shortName string, networkType string) []models.Network {
	var networks []models.Network
	db.connection.Find(&networks)
	return networks
}

func (db *networkConnection) FindById(networkID uint64) models.Network {
	var networks models.Network
	db.connection.Find(&networks, networkID)
	return networks
}

func (db *networkConnection) Insert(network models.Network) models.Network {
	db.connection.Save(&network)
	return network
}

func (db *networkConnection) Update(network models.Network) models.Network {
	db.connection.Save(&network)
	return network
}

func (db *networkConnection) Del(network models.Network) {
	db.connection.Delete(network)
}
