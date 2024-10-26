package filelink

import (
	"xseon-zero/domain/model"

	"github.com/google/uuid"
)

type FileLinkUseCase interface {
	CreateFileLink(link, caption, category string) (*model.FileLink, error)
	GetFileLinkByID(id uuid.UUID) (*model.FileLink, error)
	GetAllFileLinks() ([]model.FileLink, error)
	UpdateFileLink(id uuid.UUID, link, caption, category string) (*model.FileLink, error)
	DeleteFileLink(id uuid.UUID) error
}
