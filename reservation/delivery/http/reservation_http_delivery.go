package http

import (
	"context"
	"hotel-booking/reservation/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReservationHandler struct {
	reservationUsecase domain.ReservationUsecase
}

const (
	reservationEndpoint      = "/reservation"
	searchByCustomerEndpoint = "/search-reservation"
	searchByStatusEndpoint   = "/search-reservation-status"
)

func (h *ReservationHandler) GetReservation(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	re, _ := h.reservationUsecase.GetReservation(context.Background(), uint(id))
	ctx.JSON(200, re)
}

func (h *ReservationHandler) UpdateReservation(ctx *gin.Context) {
	var reservation domain.Reservation
	ctx.BindJSON(&reservation)
	re, _ := h.reservationUsecase.UpdateReservation(context.Background(), &reservation)
	ctx.JSON(200, re)
}

func (h *ReservationHandler) CreateReservation(ctx *gin.Context) {
	var reservation domain.Reservation
	ctx.BindJSON(&reservation)
	re, _ := h.reservationUsecase.CreateReservation(context.Background(), &reservation)
	ctx.JSON(200, re)
}
func (h *ReservationHandler) DeleteReservation(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := h.reservationUsecase.DeleteReservation(context.Background(), uint(id))
	if err != nil {
		ctx.JSON(200, id)
	}
}

func (h *ReservationHandler) SearchReservationByCustomer(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	re, _ := h.reservationUsecase.SearchReservationByCustomer(context.Background(), uint(id))
	ctx.JSON(200, re)
}
func (h *ReservationHandler) SearchReservationByStatus(ctx *gin.Context) {
	status := ctx.Query("status")
	re, _ := h.reservationUsecase.SearchReservationByStatus(context.Background(), status)
	ctx.JSON(200, re)
}

func AddReservationHandler(r *gin.Engine, repo domain.ReservationUsecase) {
	handler := &ReservationHandler{
		reservationUsecase: repo,
	}
	r.GET(reservationEndpoint+"/:id", handler.GetReservation)
	r.PUT(reservationEndpoint, handler.UpdateReservation)
	r.POST(reservationEndpoint, handler.CreateReservation)
	r.DELETE(reservationEndpoint+"/:id", handler.DeleteReservation)

	r.GET(searchByCustomerEndpoint+"/:id", handler.SearchReservationByCustomer)
	r.GET(searchByStatusEndpoint, handler.SearchReservationByStatus)
}
