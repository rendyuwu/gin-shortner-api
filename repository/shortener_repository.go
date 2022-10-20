package repository

import (
	"github.com/rendyuwu/gin-shortner-api/domain/model"
	"gorm.io/gorm"
)

type ShortenerRepository interface {
	FindAll() ([]model.Shortener, error)
	FindById(id int) (model.Shortener, error)
	FindByCode(code string) (model.Shortener, error)
	FindByCustomCode(customCode string) (model.Shortener, error)
	Create(book model.Shortener) (model.Shortener, error)
	Update(book model.Shortener) (model.Shortener, error)
	Delete(book model.Shortener) error
}

type ShortenerRepositoryImpl struct {
	DB *gorm.DB
}

func NewShortenerRepository(DB *gorm.DB) ShortenerRepository {
	return &ShortenerRepositoryImpl{DB: DB}
}

func (repository *ShortenerRepositoryImpl) FindAll() ([]model.Shortener, error) {
	var shorteners []model.Shortener

	err := repository.DB.Find(&shorteners).Error

	return shorteners, err
}

func (repository *ShortenerRepositoryImpl) FindById(id int) (model.Shortener, error) {
	var shortener model.Shortener

	err := repository.DB.Find(&shortener, id).Error

	return shortener, err
}

func (repository *ShortenerRepositoryImpl) FindByCode(code string) (model.Shortener, error) {
	var shortener model.Shortener

	err := repository.DB.Where("code = ?", code).Find(&shortener).Error

	return shortener, err
}

func (repository *ShortenerRepositoryImpl) FindByCustomCode(customCode string) (model.Shortener, error) {
	var shortener model.Shortener

	err := repository.DB.Where("custom_code = ?", customCode).Find(&shortener).Error

	return shortener, err
}

func (repository *ShortenerRepositoryImpl) Create(book model.Shortener) (model.Shortener, error) {
	err := repository.DB.Create(&book).Error

	return book, err
}

func (repository *ShortenerRepositoryImpl) Update(book model.Shortener) (model.Shortener, error) {
	err := repository.DB.Save(&book).Error

	return book, err
}

func (repository *ShortenerRepositoryImpl) Delete(book model.Shortener) error {
	err := repository.DB.Delete(&book).Error

	return err
}
