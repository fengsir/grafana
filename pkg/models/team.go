package models

import (
	"errors"
	"time"
)

// Typed errors
var (
	ErrTeamNotFound  = errors.New("Team not found")
	ErrTeamNameTaken = errors.New("Team name is taken")
)

// Team model
type Team struct {
	Id    int64  `json:"id"`
	OrgId int64  `json:"orgId"`
	Name  string `json:"name"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

// ---------------------
// COMMANDS

type CreateTeamCommand struct {
	Name  string `json:"name" binding:"Required"`
	OrgId int64  `json:"-"`

	Result Team `json:"-"`
}

type UpdateTeamCommand struct {
	Id   int64
	Name string
}

type DeleteTeamCommand struct {
	Id int64
}

type GetTeamByIdQuery struct {
	Id     int64
	Result *Team
}

type GetTeamsByUserQuery struct {
	UserId int64   `json:"userId"`
	Result []*Team `json:"teams"`
}

type SearchTeamsQuery struct {
	Query string
	Name  string
	Limit int
	Page  int
	OrgId int64

	Result SearchTeamQueryResult
}

type SearchTeamDto struct {
	Id          int64  `json:"id"`
	OrgId       int64  `json:"orgId"`
	Name        string `json:"name"`
	MemberCount int64  `json:"memberCount"`
}

type SearchTeamQueryResult struct {
	TotalCount int64            `json:"totalCount"`
	Teams      []*SearchTeamDto `json:"teams"`
	Page       int              `json:"page"`
	PerPage    int              `json:"perPage"`
}
