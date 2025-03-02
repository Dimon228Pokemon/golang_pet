package storage

import (
	"fmt"

	"github.com/Dimon228Pokemon/golang_pet/internal/file"
	"github.com/google/uuid"
)

// в этом пакете находится то что можно использовать только внутри проекта(кароче пакетики находятся)
// вся бизнес логика ,все тута
type Storage struct {
	files map[uuid.UUID]*file.File
}

// создали как бы конструктор
func NewStrorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}
	s.files[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	foundFile, ok := s.files[fileID]
	if !ok {
		return nil, fmt.Errorf("file %v not found", fileID)
	}
	return foundFile, nil
}
