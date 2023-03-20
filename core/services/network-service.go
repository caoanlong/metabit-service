package services

import (
	"metabit-service/core/models"
	"metabit-service/core/repositories"
)

type NetworkService interface {
	FindAll() (error, *[]models.Network)
	FindList() (error, *[]models.Network)
	FindById(networkID uint64) (error, *models.Network)
	Insert(network *models.Network) (error, *models.Network)
	Update(network *models.Network) (error, *models.Network)
	Del(networkID uint64) error
}

type networkService struct {
	networkRepository repositories.NetworkRepository
}

func NewNetworkService() NetworkService {
	return &networkService{
		networkRepository: repositories.NewNetworkRepository(),
	}
}

func (service *networkService) FindAll() (error, *[]models.Network) {
	return service.networkRepository.FindAll()
}

func (service *networkService) FindList() (error, *[]models.Network) {
	//TODO implement me
	panic("implement me")
}

func (service *networkService) FindById(tokenID uint64) (error, *models.Network) {
	return service.networkRepository.FindById(tokenID)
}

func (service *networkService) Insert(network *models.Network) (error, *models.Network) {
	return service.networkRepository.Insert(network)
}

func (service *networkService) Update(network *models.Network) (error, *models.Network) {
	//TODO implement me
	panic("implement me")
}

func (service *networkService) Del(networkID uint64) error {
	//TODO implement me
	panic("implement me")
}
