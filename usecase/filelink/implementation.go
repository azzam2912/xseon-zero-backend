package filelink

import (
	"log"
	"xseon-zero/domain/model"
	"xseon-zero/repository/files_db"

	"github.com/google/uuid"
)

type fileLinkImpl struct {
	filesDB files_db.FilesDBInterface
}

func NewFileLinkImpl(filesDB files_db.FilesDBInterface) FileLinkUseCase {
	return &fileLinkImpl{
		filesDB: filesDB,
	}
}

func (f *fileLinkImpl) CreateFileLink(link, caption, category string) (*model.FileLink, error) {
	fileLink := &model.FileLink{
		ID:       uuid.New(),
		Link:     link,
		Caption:  caption,
		Category: category,
	}

	tx := f.filesDB.BeginTransaction()
	err := f.filesDB.CreateFileLink(tx, fileLink)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction: %v", err)
		return nil, err
	}

	return fileLink, nil
}

func (f *fileLinkImpl) GetFileLinkByID(id uuid.UUID) (*model.FileLink, error) {
	return f.filesDB.GetFileLinkByID(id)
}

func (f *fileLinkImpl) GetAllFileLinks() ([]model.FileLink, error) {
	return f.filesDB.GetAllFileLinks()
}

func (f *fileLinkImpl) UpdateFileLink(id uuid.UUID, link, caption, category string) (*model.FileLink, error) {
	fileLink, err := f.filesDB.GetFileLinkByID(id)
	if err != nil {
		return nil, err
	}

	fileLink.Link = link
	fileLink.Caption = caption
	fileLink.Category = category

	tx := f.filesDB.BeginTransaction()
	err = f.filesDB.UpdateFileLink(tx, fileLink)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction: %v", err)
		return nil, err
	}

	return fileLink, nil
}

func (f *fileLinkImpl) DeleteFileLink(id uuid.UUID) error {
	tx := f.filesDB.BeginTransaction()
	err := f.filesDB.DeleteFileLink(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Error committing transaction: %v", err)
		return err
	}

	return nil
}
