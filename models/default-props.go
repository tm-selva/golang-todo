package model

type DefaultProps struct {
	CreatedAt  string `json:"created_at"`
	CreatedBy  string `json:"created_by"`
	DeletedBy  string `json:"deleted_by"`
	IsDeleted  string `json:"is_deleted"`
	ModifiedAt string `json:"modified_at"`
	ModifiedBy string `json:"modified_by"`
}
