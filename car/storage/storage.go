package storage

import "rent-car/models"

type IStorage interface {
	CloseDB()
	Car() ICarStorage
	Customer() ICustomerStorage
	// Order() IOrderStorage
}

type ICarStorage interface {
	Create(models.Car) (string, error)
	GetByID(id string) (models.Car, error)
	GetAll(request models.GetAllCarsRequest) (models.GetAllCarsResponse, error)
	Update(models.Car) (string, error)
	Delete(string) error
	GetAllOrdersCars(req models.GetAllCarsRequest) ([]models.GetOrder, error)
	GetByIDCustomerCarr(id string) (models.GetOrder, error)
}

type ICustomerStorage interface {
	Create(models.Customer) (string, error)
	GetByID(id string) (models.Customer, error)
	// GetAll(request models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error)
	Update(models.Customer) (string, error)
	Delete(string) error
	GetCustomerCars(req models.GetAllCarsRequest) ([]models.GetOrder, error)
	GetByIDCustomerCar(id string) (models.GetOrder, error) 
}

// type IOrderStorage interface {
// 	CreateOrder(models.CreateOrder) (string, error)
// 	UpdateOrder(models.GetOrder) (string, error)
// 	GetOne(id string) (models.GetOrder, error)
// 	GetAll(req models.GetAllOrdersRequest) (models.GetAllOrdersResponse,error) 
// 	DeleteOrder(string) error
// }