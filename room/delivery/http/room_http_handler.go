package http

import (
	"context"
	"hotel-booking/room/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
	roomUsecase domain.RoomUsecase
}

const (
	roomEndpoint               = "/room"
	searchRoomByStatusEndpoint = "/room-status"
	searchRoomByPriceEndpoint  = "/room-price"
)

func (h *RoomHandler) GetRoom(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	re, _ := h.roomUsecase.GetRoom(context.Background(), uint(id))
	ctx.JSON(200, re)
}

func (h *RoomHandler) UpdateRoom(ctx *gin.Context) {
	var room domain.Room
	_ = ctx.BindJSON(&room)
	re, _ := h.roomUsecase.UpdateRoom(context.Background(), &room)
	ctx.JSON(200, re)
}

func (h *RoomHandler) CreateRoom(ctx *gin.Context) {
	var room domain.Room
	ctx.BindJSON(&room)
	re, _ := h.roomUsecase.CreateRoom(context.Background(), &room)
	ctx.JSON(200, re)
}

func (h *RoomHandler) DeleteRoom(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := h.roomUsecase.DeleteRoom(context.Background(), uint(id))
	if err != nil {
		ctx.JSON(200, id)
	}
}

func (h *RoomHandler) SearchRoomByStatus(ctx *gin.Context) {
	status, ok := ctx.GetQuery("status")
	if !ok {
		return
	}
	re, _ := h.roomUsecase.SearchRoomByStatus(context.Background(), domain.RoomStatus(status))
	ctx.JSON(200, re)
}

func (h *RoomHandler) SearchRoomByPrice(ctx *gin.Context) {
	minPrice, ok1 := ctx.GetQuery("minPrice")
	maxPrice, ok2 := ctx.GetQuery("maxPrice")
	if !ok1 || !ok2 {
		return
	}
	min, _ := strconv.ParseFloat(minPrice, 64)
	max, _ := strconv.ParseFloat(maxPrice, 64)
	re, _ := h.roomUsecase.SearchRoomByPrice(context.Background(), min, max)
	ctx.JSON(200, re)
}

func AddRoomHandler(r *gin.Engine, repo domain.RoomUsecase) {
	handler := &RoomHandler{
		roomUsecase: repo,
	}
	r.GET(roomEndpoint+"/:id", handler.GetRoom)
	r.PUT(roomEndpoint, handler.UpdateRoom)
	r.POST(roomEndpoint, handler.CreateRoom)
	r.DELETE(roomEndpoint+"/:id", handler.DeleteRoom)

	r.GET(searchRoomByPriceEndpoint, handler.SearchRoomByPrice)
	r.GET(searchRoomByStatusEndpoint, handler.SearchRoomByStatus)
}
