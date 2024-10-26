package files_db

import (
	"log"
	"xseon-zero/domain/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type filesDBImpl struct {
	db *gorm.DB
}

func NewFilesDBImpl(db *gorm.DB) FilesDBInterface {
	return &filesDBImpl{db: db}
}

func (f *filesDBImpl) BeginTransaction() *gorm.DB {
	return f.db.Begin()
}

func (f *filesDBImpl) CreateFileLink(tx *gorm.DB, fileLink *model.FileLink) error {
	result := tx.Create(fileLink)
	if result.Error != nil {
		log.Printf("Error creating file link: %v", result.Error)
		return result.Error
	}
	return nil
}

func (f *filesDBImpl) GetFileLinkByID(id uuid.UUID) (*model.FileLink, error) {
	var fileLink model.FileLink
	result := f.db.Where("id = ?", id).First(&fileLink)
	if result.Error != nil {
		log.Printf("Error getting file link by ID: %v", result.Error)
		return nil, result.Error
	}
	return &fileLink, nil
}

func (f *filesDBImpl) GetAllFileLinks() ([]model.FileLink, error) {
	var fileLinks []model.FileLink
	result := f.db.Find(&fileLinks)
	if result.Error != nil {
		log.Printf("Error getting all file links: %v", result.Error)
		return nil, result.Error
	}
	return fileLinks, nil
}

func (f *filesDBImpl) UpdateFileLink(tx *gorm.DB, fileLink *model.FileLink) error {
	result := tx.Save(fileLink)
	if result.Error != nil {
		log.Printf("Error updating file link: %v", result.Error)
		return result.Error
	}
	return nil
}

func (f *filesDBImpl) DeleteFileLink(tx *gorm.DB, id uuid.UUID) error {
	result := tx.Delete(&model.FileLink{}, id)
	if result.Error != nil {
		log.Printf("Error deleting file link: %v", result.Error)
		return result.Error
	}
	return nil
}
