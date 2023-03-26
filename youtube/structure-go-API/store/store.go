package store

import (
	"github.com/LeonLow97/types"
)

type DB struct {
}

func New() *DB {
	return &DB{}
}

func (d DB) InsertNewStudent(s *types.Student) error {
	types.StudentsList = append(types.StudentsList, *s)
	return nil
}

func (d DB) GetAllStudents() []types.Student {
	return types.StudentsList
}
