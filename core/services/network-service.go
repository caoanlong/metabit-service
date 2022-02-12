package services

import (
	"metabit-service/core/models"
	"metabit-service/core/repositories"
)

type NetworkService interface {
	FindAll() []models.Network
	FindList() []models.Network
	FindById(networkID uint64) models.Network
	Insert(network models.Network) models.Network
	Update(network models.Network) models.Network
	Del(networkID uint64)
}

type networkService struct {
	networkRepository repositories.NetworkRepository
}

func NewNetworkService() NetworkService {
	return &networkService{
		networkRepository: repositories.NewNetworkRepository(),
	}
}

func (service *networkService) FindAll() []models.Network {
	return service.networkRepository.FindAll()
}

func (service *networkService) FindList() []models.Network {
	//TODO implement me
	panic("implement me")
}

func (service *networkService) FindById(tokenID uint64) models.Network {
	return service.networkRepository.FindById(tokenID)
}

func (service *networkService) Insert(network models.Network) models.Network {
	return service.networkRepository.Insert(network)
}

func (service *networkService) Update(network models.Network) models.Network {
	//TODO implement me
	panic("implement me")
}

func (service *networkService) Del(networkID uint64) {
	//TODO implement me
	panic("implement me")
}
