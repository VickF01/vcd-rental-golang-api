package vcd

type VCDRequest struct {
	Title       string `json:"title" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateVCDRequest struct {
	Title       string `json:"title"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
}
