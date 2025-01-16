package entity

import (
	"time"

	"github.com/google/uuid"
)

type SpecificationId uuid.UUID


func (s SpecificationId) UUID() uuid.UUID {
    return uuid.UUID(s)
}


func (s SpecificationId) String() string {
    return uuid.UUID(s).String()
}

func NewSpecificationId() SpecificationId {
	return SpecificationId(uuid.New())
}


type Specification struct {
    Id        SpecificationId
    Title     string
    Description string
    StatusId Status
    ProjectId ProjectId
    CreatedAt *time.Time
    UpdatedAt *time.Time
}


func NewSpecification(
    title string,
    description string,
    projectId ProjectId,
) *Specification {
    t := time.Now()
    return &Specification{
        Id:          NewSpecificationId(),
        Title:       title,
        Description: description,
        StatusId:    StatusCreate,
        ProjectId:   projectId,
        CreatedAt:   &t,
        UpdatedAt:   &t,
    }
}


func (s *Specification) SetTitle(title string) {
    t := time.Now()
    s.Title = title
    s.UpdatedAt = &t
}

func (s *Specification) SetDescription(description string) {
    t := time.Now()
    s.Description = description
    s.UpdatedAt = &t
}

func (s *Specification) Delete() {
    t := time.Now()
    s.UpdatedAt = &t
    s.StatusId = StatusClosed
}


    



