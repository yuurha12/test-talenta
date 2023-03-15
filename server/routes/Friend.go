package routes

import (
	"server/handlers"
	"server/pkg/mysql"
	"server/repositories"

	"github.com/labstack/echo/v4"
)

func FriendRoutes(e *echo.Group) {
	friendRepository := repositories.NewFriendRepository(mysql.DB)
	h := handlers.HandlerFriend(friendRepository)

	e.GET("/friends", h.FindFriends)
	e.GET("/friend/:id", h.GetFriend)
	e.POST("/friend", h.CreateFriend)
	e.PATCH("/friend/:id", h.UpdateFriend)
	e.DELETE("/friend/:id", h.DeleteFriend)
	e.GET("/friendstats", h.GetFriendStats)
}
