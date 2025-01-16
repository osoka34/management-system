package specification

import (
	"service/internal/app/command/specification"
	"service/internal/domain/entity"
	"time"
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

type GetSpecByProjectIdRequest struct {
	ProjectId string `required:"true" json:"project_id"`
}

func (r GetSpecByProjectIdRequest) ToCmd() *specification.GetByProjectIdCmd {
	return &specification.GetByProjectIdCmd{
		ProjectId: r.ProjectId,
	}
}

type GetSpecByProjectIdResponse struct {
	Specs []*Specification `json:"specs"`
}

type Specification struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ProjectId   string     `json:"project_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type GetSpecByIdRequest struct {
	SpecId string `required:"true" json:"spec_id"`
}

func (r GetSpecByIdRequest) ToCmd() *specification.GetByIdCmd {
	return &specification.GetByIdCmd{
		Id: r.SpecId,
	}
}

type GetSpecByIdResponse struct {
	Spec *Specification `json:"spec"`
}

func NewGetSpecByIdResponse(spec *entity.Specification) *GetSpecByIdResponse {
	return &GetSpecByIdResponse{
		Spec: FromEntity(spec),
	}
}

func FromEntity(e *entity.Specification) *Specification {
	return &Specification{
		Id:          e.Id.String(),
		Title:       e.Title,
		Description: e.Description,
		ProjectId:   e.ProjectId.String(),
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func NewGetSpecByProjectIdResponse(specs []*entity.Specification) *GetSpecByProjectIdResponse {
	out := make([]*Specification, 0, len(specs))

	for _, spec := range specs {
		out = append(out, FromEntity(spec))
	}

	return &GetSpecByProjectIdResponse{Specs: out}
}
