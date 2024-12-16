package entity

import (
	"time"

	"github.com/google/uuid"
)

type RequirementId uuid.UUID

func (r RequirementId) String() string {
    return uuid.UUID(r).String()
}

func (r RequirementId) UUID() uuid.UUID {
    return uuid.UUID(r)
}

func NewRequirementId() RequirementId {
	return RequirementId(uuid.New())
}

type Requirement struct {
	Id              RequirementId
	Title           string
	StatusId        Status
	Description     string
    Status          Status
	ExecutorId      UserId
	ProjectId       ProjectId
	SpecificationId *SpecificationId
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}

func NewRequirement(
	title string,
	description string,
	executor UserId,
	projectId ProjectId,
) *Requirement {
	t := time.Now()
	return &Requirement{
		Id:          NewRequirementId(),
		Title:       title,
		Description: description,
		ExecutorId:  executor,
        Status:      StatusCreate,
		ProjectId:   projectId,
		CreatedAt:   &t,
		UpdatedAt:   &t,
	}
}


func (r *Requirement) SetSpecification(specificationId SpecificationId) {
    t := time.Now()
    r.SpecificationId = &specificationId
    r.UpdatedAt = &t
}

func (r *Requirement) UpdateExecutor(executor UserId) {
    t := time.Now()
    r.ExecutorId = executor
    r.UpdatedAt = &t
}

func (r *Requirement) UpdateTitle(title string) {
    t := time.Now()
    r.Title = title
    r.UpdatedAt = &t
}

func (r *Requirement) UpdateDescription(description string) {
    t := time.Now()
    r.Description = description
    r.UpdatedAt = &t
}


func (r *Requirement) Delete() {
    t := time.Now()
    r.UpdatedAt = &t
    r.Status = StatusClosed
}
