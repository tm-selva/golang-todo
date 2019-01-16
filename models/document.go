package model

type Document struct {
	DefaultProps
	DocName        string `json:"doc_name"`
	FileExtension  string `json:"file_extentsion"`
	FilePath       string `json:"file_path"`
	FileSizeInByte string `json:"file_size_in_byte"`
	OriginalName   string `json:"original_name"`
	UserId         string `json:"user_id"`
}
