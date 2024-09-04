package model

import (
	"time"

	"github.com/google/uuid"
)

const (
	DirectoryType = "CollectionType"
	DocumentType  = "DocumentType"
)

type Document struct {
	BlobURLGet        string
	BlobURLGetExpires string
	Bookmarked        bool
	CurrentPage       int
	FileTags          []string
	ID                string
	Message           string
	ModifiedClient    string
	PageTags          []string
	Parent            string
	Success           bool
	Type              string
	Version           int
	VissibleName      string
}

type MetadataDocument struct {
	FileTags       []string
	ID             string
	ModifiedClient string
	PageTags       []string
	Parent         string
	Type           string
	Version        int
	VissibleName   string
}

type DeleteDocument struct {
	ID      string
	Version int
}

type UploadDocumentRequest struct {
	ID      string
	Type    string
	Version int
}

type UploadDocumentResponse struct {
	ID                string
	Version           int
	Message           string
	Success           bool
	BlobURLPut        string
	BlobURLPutExpires string
}

type BlobRootStorageRequest struct {
	Method       string `json:"http_method"`
	Initial      bool   `json:"initial_sync,omitempty"`
	RelativePath string `json:"relative_path"`
	RootSchema   string `json:"root_schema,omitempty"`
	Generation   int64  `json:"generation"`
}

// BlobStorageRequest request
type BlobStorageRequest struct {
	Method       string `json:"http_method"`
	Initial      bool   `json:"initial_sync,omitempty"`
	RelativePath string `json:"relative_path"`
	ParentPath   string `json:"parent_path,omitempty"`
}

// BlobStorageResponse response
type BlobStorageResponse struct {
	Expires            string `json:"expires"`
	Method             string `json:"method"`
	RelativePath       string `json:"relative_path"`
	Url                string `json:"url"`
	MaxUploadSizeBytes int64  `json:"maxuploadsize_bytes,omitifempty"`
}

// SyncCompleteRequest payload of the sync completion
type SyncCompletedRequest struct {
	Generation int64 `json:"generation"`
}

func CreateDirDocument(parent, name string) MetadataDocument {
	id := uuid.New()

	return MetadataDocument{
		ID:             id.String(),
		Parent:         parent,
		VissibleName:   name,
		Type:           DirectoryType,
		Version:        1,
		ModifiedClient: time.Now().UTC().Format(time.RFC3339Nano),
	}
}

func CreateUploadDocumentRequest(id string, entryType string) UploadDocumentRequest {
	if id == "" {
		newId := uuid.New()

		id = newId.String()
	}

	return UploadDocumentRequest{
		id,
		entryType,
		1,
	}
}

func CreateUploadDocumentMeta(id string, entryType, parent, name string) MetadataDocument {

	return MetadataDocument{
		ID:             id,
		Parent:         parent,
		VissibleName:   name,
		Type:           entryType,
		Version:        1,
		ModifiedClient: time.Now().UTC().Format(time.RFC3339Nano),
	}
}

func (meta MetadataDocument) ToDocument() Document {
	return Document{
		ID:             meta.ID,
		Parent:         meta.Parent,
		VissibleName:   meta.VissibleName,
		Type:           meta.Type,
		Version:        1,
		ModifiedClient: meta.ModifiedClient,
		FileTags:       meta.FileTags,
		PageTags:       meta.PageTags,
	}
}

func (doc Document) ToMetaDocument() MetadataDocument {
	return MetadataDocument{
		ID:             doc.ID,
		Parent:         doc.Parent,
		VissibleName:   doc.VissibleName,
		Type:           doc.Type,
		FileTags:       doc.FileTags,
		PageTags:       doc.PageTags,
		Version:        doc.Version,
		ModifiedClient: time.Now().UTC().Format(time.RFC3339Nano),
	}
}

func (doc Document) ToDeleteDocument() DeleteDocument {
	return DeleteDocument{
		ID:      doc.ID,
		Version: doc.Version,
	}
}
