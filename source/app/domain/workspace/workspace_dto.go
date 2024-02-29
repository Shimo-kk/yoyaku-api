package workspace

import "time"

type Workspace struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

// エンティティをDTOに変換
func ToDtoFromEntity(ue *WorkspaceEntity) *Workspace {
	return &Workspace{
		Id:        ue.id,
		CreatedAt: ue.createdAt,
		UpdatedAt: ue.updatedAt,
		Name:      ue.name,
	}
}

// DTOをエンティティに変換
func (ud *Workspace) ToEntity() *WorkspaceEntity {
	return &WorkspaceEntity{
		id:        ud.Id,
		createdAt: ud.CreatedAt,
		updatedAt: ud.UpdatedAt,
		name:      ud.Name,
	}
}
