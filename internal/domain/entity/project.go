package entity 

import (
	"time"

	"github.com/google/uuid"
)

type ProjectId uuid.UUID

func NewProjectId() ProjectId {
	return ProjectId(uuid.New())
}

func (p ProjectId) UUID() uuid.UUID {
    return uuid.UUID(p)
}


func (p ProjectId) String() string {
    return uuid.UUID(p).String()
} 


type Project struct {
    Id        ProjectId
    Title     string
    StatusId Status
    CreatorId UserId
    Description string
    CreatedAt *time.Time
    UpdatedAt *time.Time
}


func NewProject(
    title string,
    creator UserId,
    description string,
) *Project {
    t := time.Now()
    return &Project{
        Id:          NewProjectId(),
        Title:       title,
        StatusId:    StatusCreate,
        CreatorId:   creator,
        Description: description,
        CreatedAt:   &t,
        UpdatedAt:   &t,
    }
}


func (p *Project) UpdateTitle(title string) {
    if title == "" || p.Title == title {
        return
    }
    t := time.Now()
    p.Title = title
    p.UpdatedAt = &t
}

func (p *Project) UpdateDescription(description string) {
    if description == "" || p.Description == description {
        return
    }
    t := time.Now()
    p.Description = description
    p.UpdatedAt = &t
}

func (p *Project) Delete() {
    t := time.Now()
    p.UpdatedAt = &t
    p.StatusId = StatusClosed
}


