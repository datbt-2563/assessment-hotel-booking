package http

import (
	"context"
	"hotel-booking/customer/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerUsecase domain.CustomerUsecase
}

const (
	customerEndpoint              = "/customer"
	searchCustomerByNameEndpoint  = "/customer-name"
	searchCustomerByPhoneEndpoint = "/customer-phone"
)

func (h *CustomerHandler) GetCustomer(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	re, _ := h.customerUsecase.GetCustomer(context.Background(), uint(id))
	ctx.JSON(200, re)
}

func (h *CustomerHandler) UpdateCustomer(ctx *gin.Context) {
	var customer domain.Customer
	_ = ctx.BindJSON(&customer)
	re, _ := h.customerUsecase.UpdateCustomer(context.Background(), &customer)
	ctx.JSON(200, re)
}

func (h *CustomerHandler) CreateCustomer(ctx *gin.Context) {
	var customer domain.Customer
	ctx.BindJSON(&customer)
	re, _ := h.customerUsecase.CreateCustomer(context.Background(), &customer)
	ctx.JSON(200, re)
}

func (h *CustomerHandler) DeleteCustomer(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := h.customerUsecase.DeleteCustomer(context.Background(), uint(id))
	if err != nil {
		ctx.JSON(200, id)
	}
}

func (h *CustomerHandler) SearchCustomerByName(ctx *gin.Context) {
	name, ok := ctx.GetQuery("name")
	if !ok {
		return
	}
	re, _ := h.customerUsecase.SearchCustomerByName(context.Background(), name)
	ctx.JSON(200, re)
}

func (h *CustomerHandler) SearchCustomerByPhone(ctx *gin.Context) {
	phone, ok := ctx.GetQuery("phone")
	if !ok {
		return
	}
	re, _ := h.customerUsecase.SearchCustomerByPhone(context.Background(), phone)
	ctx.JSON(200, re)
}

func AddCustomerHandler(r *gin.Engine, repo domain.CustomerUsecase) {
	handler := &CustomerHandler{
		customerUsecase: repo,
	}
	r.GET(customerEndpoint+"/:id", handler.GetCustomer)
	r.PUT(customerEndpoint, handler.UpdateCustomer)
	r.POST(customerEndpoint, handler.CreateCustomer)
	r.DELETE(customerEndpoint+"/:id", handler.DeleteCustomer)

	r.GET(searchCustomerByNameEndpoint, handler.SearchCustomerByName)
	r.GET(searchCustomerByPhoneEndpoint, handler.SearchCustomerByPhone)
}
