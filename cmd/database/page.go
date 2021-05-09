package database

type GenericPage struct {
	Page     int `json:"int64"`
	PageSize int `json:"page_size"`
	Count    int `json:"count"`
}

type PageOfFolders struct {
	GenericPage
	Items []Folder `json:"items"`
}
type PageOfNotes struct {
	GenericPage
	Items []Note `json:"items"`
}
