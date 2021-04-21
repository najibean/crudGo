package app

import "github.com/najibean/crudGo/app/models"

type Model struct {
	Model interface {}
}

func RegisterModels() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Address{}},
		{Model: models.Product{}},
		{Model: models.ProductImage{}},
		{Model: models.Section{}},
		{Model: models.Category{}},

	}
}