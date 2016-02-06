package engine

import (
    "testing"
    "os"
)

func TestDocumentOperations(t *testing.T) {
    fsStorage := initStorage(t)

    project, err := fsStorage.CreateProject("project1")

    if err != nil || project == "" {
        t.Errorf("Failed to create project %s", err)
    }

    err = fsStorage.CreateDocument("project1", "document1", []byte("document1"))
    if err != nil {
        t.Errorf("Failed to create document %s", err)
    }
    
    err = fsStorage.CreateDocument("project1", "document2", []byte("document2"))
    if err != nil {
        t.Errorf("Failed to create document %s", err)
    }

    files, err := fsStorage.Documents("project1")

    if err != nil {
        t.Errorf("Failed to get list of documents %s", err)
    }
    if len(files) != 2 {
        t.Errorf("Expected the project to contain 2 documents but instead it contains %v", len(files))
    }

    content, err := fsStorage.Document("project1", "document1")
    if err != nil {
        t.Errorf("Failed to get the contents of document1 %s", err)
    }
    if string(content) != "document1" {
        t.Errorf("Expect document1 to contain 'document1' but instead got %s", string(content))
    }

    content, err = fsStorage.Document("project1", "document2")
    if err != nil {
        t.Errorf("Failed to get the contents of document2 %s", err)
    }
    if string(content) != "document2" {
        t.Errorf("Expect document2 to contain 'document2' but instead got %s", string(content))
    }

    err = fsStorage.DeleteDocument("project1", "document1")
    if err != nil {
        t.Errorf("Failed to delete document1 %s", err)
    }

    content, err = fsStorage.Document("project1", "document1")
    if os.IsNotExist(err) == false {
        t.Errorf("Expected document1 to be removed", err)
    }
    removeStorage(t, fsStorage)
}

func initStorage(t *testing.T) *FilesystemStorage {
    tempDir := os.TempDir()
    storageHome := tempDir + ".huskydocs/"

    fsStorage := &FilesystemStorage{storageHome}
    fsStorage.Init()
    if _, err := os.Stat(storageHome); os.IsNotExist(err) {
        t.Errorf("Filesystem storage was initialized at '%s' but the directory doesn't exist", storageHome)
    }
    return fsStorage
}

func removeStorage(t *testing.T, fsStorage *FilesystemStorage) {
    err := os.RemoveAll(fsStorage.storageHome)
    if err != nil {
        t.Errorf("Failed to remove storage dir from %s", fsStorage.storageHome)
    }
}
