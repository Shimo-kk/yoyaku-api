package workspace

type IWorkspaceRepository interface {
	Insert(entity *WorkspaceEntity) (*WorkspaceEntity, error)
	NotExists(name string) error
	FindById(id int) (*WorkspaceEntity, error)
	FindByName(email string) (*WorkspaceEntity, error)
	DeleteById(id int) error
}
