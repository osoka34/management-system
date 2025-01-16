package project

import (
	"time"

	"service/internal/app/command/project"
	"service/internal/domain/entity"
)

type CreateProjectRequest struct {
	CreatorId   string `json:"creator_id"  required:"true"`
	Title       string `json:"title"       required:"true"`
	Description string `json:"description" required:"true"`
}

func (r CreateProjectRequest) ToCmd() *project.CreateProjectCmd {
	return &project.CreateProjectCmd{
		CreatorId:   r.CreatorId,
		Title:       r.Title,
		Description: r.Description,
	}
}

type UpdateProjectRequest struct {
	Id          string `json:"id"          required:"true"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (r UpdateProjectRequest) ToCmd() *project.UpdateProjectCmd {
	return &project.UpdateProjectCmd{
		Id:          r.Id,
		Title:       r.Title,
		Description: r.Description,
	}
}

type DeleteProjectRequest struct {
	Id string `json:"id" required:"true"`
}

func (r DeleteProjectRequest) ToCmd() *project.DeleteProjectCmd {
	return &project.DeleteProjectCmd{
		Id: r.Id,
	}
}

type GetAllProjectsResponse struct {
	Projects []*Project `json:"projects"`
}


func NewGetAllProjectsResponse(projects []*entity.Project) GetAllProjectsResponse {
    var projectsDto = make([]*Project, 0, len(projects))
    for _, project := range projects {
        projectsDto = append(projectsDto, FromEntity(project))
    }
    return GetAllProjectsResponse{
        Projects: projectsDto,
    }
}

type Project struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	StatusId    int        `json:"status_id"`
	CreatorId   string     `json:"creator_id"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func FromEntity(project *entity.Project) *Project {
	return &Project{
		Id:          project.Id.String(),
		Title:       project.Title,
		StatusId:    int(project.StatusId),
		CreatorId:   project.CreatorId.String(),
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}
}

type CreateProjectResponse struct {
	Id string `json:"id"`
}

type UpdateProjectResponse struct {
	Id string `json:"id"`
}

type DeleteProjectResponse struct {
	Id string `json:"id"`
}
