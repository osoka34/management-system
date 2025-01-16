package specification

import (
	"service/internal/app/command/specification"
)

type CreateSpecificationRequest struct {
	ProjectId   string `required:"true" json:"project_id"`
	Title       string `required:"true" json:"title"`
	Description string `required:"true" json:"description"`
}

func (r CreateSpecificationRequest) ToCmd() *specification.CreateSpecificationCmd {
	return &specification.CreateSpecificationCmd{
		ProjectId:   r.ProjectId,
		Title:       r.Title,
		Description: r.Description,
	}
}

type UpdateSpecificationRequest struct {
	Id          string `json"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (r UpdateSpecificationRequest) ToCmd() *specification.UpdateSpecificationCmd {
	return &specification.UpdateSpecificationCmd{
		Id:          r.Id,
		Title:       r.Title,
		Description: r.Description,
	}
}


type DeleteSpecificationRequest struct {
    Id string `json:"id"`
}

func (r DeleteSpecificationRequest) ToCmd() *specification.DeleteSpecificationCmd {
    return &specification.DeleteSpecificationCmd{
        Id: r.Id,
    }
}



type CreateSpecificationResponse struct {
    Id string `json:"id"`
}

type UpdateSpecificationResponse struct {
    Id string `json:"id"`
}

type DeleteSpecificationResponse struct {
    Id string `json:"id"`
}
