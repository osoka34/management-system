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
		StatusId:      StatusCreate,
		ProjectId:   projectId,
		CreatedAt:   &t,
		UpdatedAt:   &t,
	}
}

func (r *Requirement) SetSpecification(specificationId SpecificationId) {
	if specificationId.String() == "" {
		return
	}
	t := time.Now()
	r.SpecificationId = &specificationId
	r.UpdatedAt = &t
}

func (r *Requirement) SetExecutor(executor UserId) {
    if executor.String() == "" {
        return
    }
	t := time.Now()
	r.ExecutorId = executor
	r.UpdatedAt = &t
}

func (r *Requirement) SetTitle(title string) {
    if title == "" || r.Title == title {
        return
    }
	t := time.Now()
	r.Title = title
	r.UpdatedAt = &t
}

func (r *Requirement) SetDescription(description string) {
    if description == "" || r.Description == description {
        return
    }
	t := time.Now()
	r.Description = description
	r.UpdatedAt = &t
}

func (r *Requirement) Delete() {
	t := time.Now()
	r.UpdatedAt = &t
	r.StatusId = StatusClosed
}
