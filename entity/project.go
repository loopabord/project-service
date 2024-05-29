package entity

import (
	"time"

	"github.com/uptrace/bun"
)

type Project struct {
	bun.BaseModel `bun:"table:projects,alias:p"`

	Id                string    `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name              string    `bun:"name,notnull" json:"name"`
	Description       string    `bun:"description" json:"description"`
	AuthorId          string    `bun:"author_id,type:uuid,notnull" json:"author_id"`
	AuthorName        string    `bun:"author_name,notnull" json:"author_name"`
	CollaboratorIds   []string  `bun:"collaborator_ids,array" json:"collaborator_ids"`
	CollaboratorNames []string  `bun:"collaborator_names,array" json:"collaborator_names"`
	CreatedAt         time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
}
