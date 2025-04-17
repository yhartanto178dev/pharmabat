package drug

type CreateDrugRequest struct {
	Name string `json:"name" validate:"required"`
}
