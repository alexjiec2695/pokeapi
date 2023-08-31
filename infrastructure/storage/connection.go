package storage

import (
	"context"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"pokeapi/entities"
)

type store struct {
	db *gorm.DB
}

func NewConnectionStorage(strConnection string) (*store, error) {
	db, err := gorm.Open(mysql.Open(strConnection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entities.Users{})
	if err != nil {
		return nil, err
	}

	return &store{
		db: db,
	}, nil
}

func (s *store) CreateUsers(ctx context.Context, user entities.Users) error {
	tx := s.db.Create(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (s *store) GetLogin(ctx context.Context, user entities.Users) (string, error) {
	result := entities.Users{}

	tx := s.db.Where("Email= ? and Password = ?", user.Email, user.Password).Take(&result)

	if tx.Error != nil {
		if tx.Error.Error() == "record not found" {
			return "", errors.New("not found")
		}
		return "", tx.Error
	}

	if tx.RowsAffected != 1 {
		return "", errors.New("user or password is wrong")
	}

	return result.ID, nil
}

func (s *store) UpdateUsers(ctx context.Context, user entities.Users) error {
	tx := s.db.Model(&entities.Users{}).Where("ID = ?", user.ID).Updates(user)
	return tx.Error
}

func (s *store) GetUsers(ctx context.Context, id string) (entities.Users, error) {
	result := entities.Users{}
	tx := s.db.Where("ID = ?", id).Take(&result)
	return result, tx.Error
}
