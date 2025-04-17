package enduser

type CreateEndUserRequest struct {
	Name string `json:"name" validate:"required"`
}
