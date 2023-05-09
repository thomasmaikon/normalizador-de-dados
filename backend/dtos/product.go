package dtos

import "encoding/json"

type ProductDTO struct {
	Description string      `json:description`
	Price       json.Number `json:price`
}
