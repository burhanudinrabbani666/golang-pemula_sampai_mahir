package repository

import "golang-pemula_sampai_mahir/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
