package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Project struct {
	bun.BaseModel `bun:"table:projects,alias:p"`

	Id            uuid.UUID             `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	CreatedAt     timestamppb.Timestamp `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	Name          string                `bun:"name,notnull" json:"name"`
	Description   string                `bun:"description" json:"description"`
	AuthorId      uuid.UUID             `bun:"author_id,type:uuid,notnull" json:"author_id"`
	Collaborators []uuid.UUID           `bun:"collaborators,array" json:"collaborators"`
	AuthorName    string                `bun:"author_name,notnull" json:"author_name"`
}
