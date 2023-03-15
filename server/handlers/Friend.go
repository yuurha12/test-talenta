package handlers

import (
	"net/http"
	friendsdto "server/dto/friend"
	dto "server/dto/result"
	"server/models"
	"server/repositories"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handler struct {
	FriendRepository repositories.FriendRepository
}

func HandlerFriend(FriendRepository repositories.FriendRepository) *handler {
	return &handler{FriendRepository}
}

func (h *handler) FindFriends(c echo.Context) error {
	friends, err := h.FriendRepository.FindFriends()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: friends})
}

func (h *handler) GetFriend(c echo.Context) error {
	id := c.Param("id")

	friendID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	friend, err := h.FriendRepository.GetFriend(int(friendID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(friend)})
}

func (h *handler) CreateFriend(c echo.Context) error {
	request := new(friendsdto.CreateFriendRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	friend := models.Friend{
		Name:   request.Name,
		Gender: request.Gender,
		Age:    int(request.Age),
	}

	data, err := h.FriendRepository.CreateFriend(friend)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

func (h *handler) UpdateFriend(c echo.Context) error {
	request := new(friendsdto.UpdateFriendRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	friendID := int(id)

	friend, err := h.FriendRepository.GetFriend(friendID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.FriendRepository.UpdateFriend(friend, friendID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

func (h *handler) DeleteFriend(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	friend, err := h.FriendRepository.GetFriend(int(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	err = h.FriendRepository.DeleteFriend(int(friend.ID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(friend)})
}

func (h *handler) GetFriendStats(c echo.Context) error {
	stats, err := h.FriendRepository.GetFriendStats()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	result := friendsdto.ResultStats{
		Code: http.StatusOK,
		Data: stats,
	}

	return c.JSON(http.StatusOK, result)
}

func convertResponse(u models.Friend) models.Friend {
	return models.Friend{
		ID:     u.ID,
		Name:   u.Name,
		Gender: u.Gender,
		Age:    u.Age,
	}
}
