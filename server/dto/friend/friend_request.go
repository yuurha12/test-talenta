package frienddto

type CreateFriendRequest struct {
	Name   string `json:"name" form:"name" validate:"required"`
	Gender string `json:"gender" form:"gender" validate:"required"`
	Age    int    `json:"age" form:"age" validate:"required"`
}

type UpdateFriendRequest struct {
	Name   string `json:"name" form:"name"`
	Gender string `json:"gender" form:"gender"`
	Age    int    `json:"age" form:"age"`
}