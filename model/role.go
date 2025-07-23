package model

import "time"

type Skill struct {
	ID        string    `json:"id"`
	OrgID     string    `json:"orgId" db:"org_id"`
	SkillID   string    `json:"skillId" db:"skill_id"`
	En        string    `json:"en" db:"en"`
	Th        string    `json:"th" db:"th"`
	Active    bool      `json:"active" db:"active"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	CreatedBy string    `json:"createdBy" db:"created_by"`
	UpdatedBy string    `json:"updatedBy" db:"updated_by"`
}

type SkillInsert struct {
	En     string `json:"en" db:"en"`
	Th     string `json:"th" db:"th"`
	Active bool   `json:"active" db:"active"`
}

type SkillUpdate struct {
	En     string `json:"en" db:"en"`
	Th     string `json:"th" db:"th"`
	Active bool   `json:"active" db:"active"`
}
