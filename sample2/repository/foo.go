package repository

import (
	"entity"
)

type FooRepository interface {
	FindByID(id int64) (*entity.Foo, error)
	Insert(foo *entity.Foo) error
	Update(foo *entity.Foo) error
}
