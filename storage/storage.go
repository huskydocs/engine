package storage

import (
    "fmt"
    "os"
    "io/ioutil"
    "path/filepath"
)

type Storage interface {
    Init()
}

type ProjectStorage interface {
    CreateProject(projectId string) (string, error)
}

type DocumentStorage interface {
    Documents(projectId string) ([]os.FileInfo, error)
    Document(projectId, documentId string) ([]byte, error)
    CreateDocument(projectId, documentId string, content []byte) error
    UpdateDocument(projectId, documentId string, content []byte) error
    DeleteDocument(projectId, documentId string) error
}

type FilesystemStorage struct {
    storageHome string
}

func (fs FilesystemStorage) Init() {
    err := os.Mkdir(fs.storageHome, 0700)
    if err == nil || os.IsExist(err) {
       fmt.Printf("Created Huskydocs Filesystem storage at %s", fs.storageHome)
       return
    }
    panic(err)
}

func (fs FilesystemStorage) CreateProject(projectId string) (string, error) {
    project := filepath.Join(fs.storageHome, projectId)
    err := os.Mkdir(project, 0700)
    if err == nil || os.IsExist(err) {
       fmt.Printf("Created Huskydocs project at %s\n", project)
       return project, nil
    }
    return "", err
}

func (fs FilesystemStorage) Documents(projectId string) ([]os.FileInfo, error) {
    files, error := ioutil.ReadDir(filepath.Join(fs.storageHome, projectId))
    if (error != nil) {
        return nil, error
    }

    return files, nil
}

func (fs FilesystemStorage) Document(projectId, documentId string) ([]byte, error) {
    content, error := ioutil.ReadFile(filepath.Join(fs.storageHome, projectId, documentId))
    if (error != nil) {
        return nil, error
    }

    return content, nil
}

func (fs FilesystemStorage) CreateDocument(projectId, documentId string, content []byte) error {
    err := ioutil.WriteFile(filepath.Join(fs.storageHome, projectId, documentId), content, 0644)
    return err
}

func (fs FilesystemStorage) UpdateDocument(projectId, documentId string, content []byte) error {
    error := ioutil.WriteFile(filepath.Join(fs.storageHome, projectId, documentId), content, 0600)
    return error
}

func (fs FilesystemStorage) DeleteDocument(projectId, documentId string) error {
    error := os.Remove(filepath.Join(fs.storageHome, projectId, documentId))
    return error
}
