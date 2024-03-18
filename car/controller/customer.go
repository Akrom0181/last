package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rent-car/models"
	"rent-car/pkg/check"

	"github.com/google/uuid"
)

func (c Controller) Customer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateCustomer(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		_, ok := values["id"]
		if !ok {
			c.GetCustomerCars(w, r)
		} else {
			c.GetByIDCustomeCarr(w, r)
		}
	case http.MethodPut:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.UpdateCustomer(w, r)
		}

	case http.MethodDelete:
		values := r.URL.Query()
		_, ok := values["id"]
		if ok {
			c.DeleteCustomer(w, r)
		}

	default:
		handleResponse(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

func (c Controller) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	customer := models.Customer{}

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		errStr := fmt.Sprintf("error while decoding request body,err:%v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}

	if err := check.ValidateGmailAddress(customer.Gmail); err != nil {
		fmt.Println("error while validating email address")
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := check.ValidatePhoneNumber(customer.Phone); err != nil {
		fmt.Println("error while validating phoneNumber: ", customer.Phone)
		handleResponse(w, http.StatusBadRequest, err)
		return
	}

	id, err := c.Store.Customer().Create(customer)
	if err != nil {
		fmt.Println("error while creating customer,err:", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, id)
}
func (c Controller) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	customer := models.Customer{}
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		errStr := fmt.Sprintf("error while decoding request body,err:%v\n", err)
		fmt.Println(errStr)
		handleResponse(w, http.StatusBadRequest, errStr)
		return
	}
	customer.Id = r.URL.Query().Get("id")
	err := uuid.Validate(customer.Id)
	if err != nil {
		fmt.Println("error while validating,err", err)
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := c.Store.Customer().Update(customer)
	if err != nil {
		fmt.Println("error while updating customer,err", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, id)
}

// func (c Controller) GetAllCustomer(w http.ResponseWriter, r *http.Request) {
// 	var (
// 		values  = r.URL.Query()
// 		search  string
// 		request = models.GetAllCustomersRequest{}
// 	)
// 	if _, ok := values["search"]; ok {
// 		search = values["search"][0]
// 	}
// 	request.Search = search

// 	page, err := ParsePageQueryParam(r)
// 	if err != nil {
// 		fmt.Println("error while parsing page, err: ", err)
// 		handleResponse(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	limit, err := ParseLimitQueryParam(r)
// 	if err != nil {
// 		fmt.Println("error while parsing limit, err: ", err)
// 		handleResponse(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	fmt.Println("page: ", page)
// 	fmt.Println("limit: ", limit)

// 	request.Page = page
// 	request.Limit = limit

// 	customers, err := c.Store.Customer().GetAll(request)
// 	if err != nil {
// 		fmt.Println("error while getting customers,err:", err)
// 		handleResponse(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	handleResponse(w, http.StatusOK, customers)
// }
// func (c Controller) GetByIDCustomer(w http.ResponseWriter, r *http.Request) {
// 	values := r.URL.Query()
// 	id := values["id"][0]

// 	customer, err := c.Store.Customer().GetByID(id)
// 	if err != nil {
// 		fmt.Println("error while getting customer by id")
// 		handleResponse(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	handleResponse(w, http.StatusOK, customer)
// }

func (c Controller) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println("id", id)

	err := uuid.Validate(id)
	if err != nil {
		fmt.Println("error while validating id,err:", err.Error())
		handleResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	err = c.Store.Customer().Delete(id)
	if err != nil {
		fmt.Println("error while deleting customer, err:", err)
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, id)
}

func (c Controller) GetCustomerCars(w http.ResponseWriter,r *http.Request)  {
	var (
		values  = r.URL.Query()
		search  string
		request = models.GetAllCarsRequest{}
	)
	if _, ok := values["search"]; ok {
		search = values["search"][0]
	}

	request.Search = search

	page, err := ParsePageQueryParam(r)
	if err != nil {
		fmt.Println("error while parsing page, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(r)
	if err != nil {
		fmt.Println("error while parsing limit, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit

	customers, err := c.Store.Customer().GetCustomerCars(request)
	if err != nil {
		fmt.Println("error while getting customers, err: ", err)
		handleResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(w, http.StatusOK, customers)
}

func (c Controller) GetByIDCustomeCarr(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values["id"][0]

	customer, err := c.Store.Customer().GetByIDCustomerCar(id)
	if err != nil {
		fmt.Println("error while getting customerCar by id")
		handleResponse(w, http.StatusInternalServerError, err)
		return
	}
	handleResponse(w, http.StatusOK, customer)


}