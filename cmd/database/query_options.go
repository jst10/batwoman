package database

type QueryOptions struct {
	Page           int
	PageSize       int
	OrderBy        string
	OrderDirection string
}

type FolderQueryOptions struct {
	QueryOptions
	Name    string
	OwnerId string
}
type NoteQueryOptions struct {
	FolderQueryOptions
	SharedOption   string
	FolderId   string
	OrPublic   bool
}
