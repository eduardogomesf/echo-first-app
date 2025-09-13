package dtos

type AddProductDto struct {
	Name       string   `json:"name"`
	Price      float64  `json:"price"`
	Categories []string `json:"categories"`
	IsDisabled bool     `json:"isDisabled"`
}
