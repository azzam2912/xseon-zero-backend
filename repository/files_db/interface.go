package files_db

import (
	"xseon-zero/domain/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilesDBInterface interface {
	BeginTransaction() *gorm.DB
	CreateFileLink(tx *gorm.DB, fileLink *model.FileLink) error
	GetFileLinkByID(id uuid.UUID) (*model.FileLink, error)
	GetAllFileLinks() ([]model.FileLink, error)
	UpdateFileLink(tx *gorm.DB, fileLink *model.FileLink) error
	DeleteFileLink(tx *gorm.DB, id uuid.UUID) error
}
