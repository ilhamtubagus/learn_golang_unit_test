package repository

import "github.com/ilhamtubagus/learn_golang_unit_test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
