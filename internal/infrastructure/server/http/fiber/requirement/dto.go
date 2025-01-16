package requirement

import (
	"time"

	"service/internal/app/command/requirement"
	"service/internal/domain/entity"
)

type CreateRequirementRequest struct {
	ProjectId   string `json:"project_id"  required:"true"`
	Title       string `json:"title"       required:"true"`
	Description string `json:"description" required:"true"`
	ExecutorId  string `json:"executor_id" required:"true"`
}

func (r CreateRequirementRequest) ToCmd() *requirement.CreateRequirementCmd {
	return &requirement.CreateRequirementCmd{
		ProjectId:   r.ProjectId,
		Title:       r.Title,
		Description: r.Description,
		ExecutorId:  r.ExecutorId,
	}
}

type UpdateRequirementRequest struct {
	Id              string `json:"id"               required:"true"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	ExecutorId      string `json:"executor_id"`
	SpecificationId string `json:"specification_id"`
}

func (r UpdateRequirementRequest) ToCmd() *requirement.UpdateRequirementCmd {
	return &requirement.UpdateRequirementCmd{
		Id:              r.Id,
		Title:           r.Title,
		Description:     r.Description,
		ExecutorId:      r.ExecutorId,
		SpecificationId: r.SpecificationId,
	}
}

type DeleteRequirementRequest struct {
	Id string `json:"id" required:"true"`
}

func (r DeleteRequirementRequest) ToCmd() *requirement.DeleteRequirementCmd {
	return &requirement.DeleteRequirementCmd{
		Id: r.Id,
	}
}

type CreateRequirementResponse struct {
	Id string `json:"id"`
}

type UpdateRequirementResponse struct {
	Id string `json:"id"`
}

type DeleteRequirementResponse struct {
	Id string `json:"id"`
}

type AddInSpecRequest struct {
	Ids             []string `json:"ids" "required:"true"`
	SpecificationId string   `json:"specification_id" required:"true"`
}

func (r AddInSpecRequest) ToCmd() *requirement.AddInSpecificationCmd {
	return &requirement.AddInSpecificationCmd{
		Ids:             r.Ids,
		SpecificationId: r.SpecificationId,
	}
}

type AddInSpecResponse struct {
	Ids []string `json:"ids"`
}

type Requirement struct {
	Id              string     `json:"id"`
	Title           string     `json:"title"`
	StatusId        int        `json:"status_id"`
	Description     string     `json:"description"`
	ExecutorId      string     `json:"executor_id"`
	ProjectId       string     `json:"project_id"`
	SpecificationId *string    `json:"specification_id"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

type GetProjectRequirementsRequest struct {
	ProjectId string `json:"project_id" required:"true"`
}

func (r GetProjectRequirementsRequest) ToCmd() *requirement.GetProjectRepuirementsCmd {
	return &requirement.GetProjectRepuirementsCmd{
		ProjectId: r.ProjectId,
	}
}

type GetProjectRequirementsResponse struct {
	Requirements []*Requirement `json:"requirements"`
}

func NewGetProjectRequirementsResponse(
	requirements []*entity.Requirement,
) *GetProjectRequirementsResponse {
	var requirementsDto []*Requirement
	for _, requirement := range requirements {
		requirementsDto = append(requirementsDto, FromEntity(requirement))
	}

	return &GetProjectRequirementsResponse{
		Requirements: requirementsDto,
	}
}

type GetSpecificationRequirementsRequest struct {
	SpecificationId string `json:"specification_id" required:"true"`
}

func (r GetSpecificationRequirementsRequest) ToCmd() *requirement.GetSpecRequirementsCmd {
	return &requirement.GetSpecRequirementsCmd{
		SpecificationId: r.SpecificationId,
	}
}

type GetSpecificationRequirementsResponse struct {
	Requirements []*Requirement `json:"requirements"`
}

func NewGetSpecificationRequirementsResponse(
	requirements []*entity.Requirement,
) *GetSpecificationRequirementsResponse {
	var requirementsDto []*Requirement
	for _, requirement := range requirements {
		requirementsDto = append(requirementsDto, FromEntity(requirement))
	}

	return &GetSpecificationRequirementsResponse{
		Requirements: requirementsDto,
	}
}

func FromEntity(requirement *entity.Requirement) *Requirement {
	var specId *string
	if requirement.SpecificationId != nil {
		id := requirement.SpecificationId.String()
		specId = &id
	}

	return &Requirement{
		Id:              requirement.Id.String(),
		Title:           requirement.Title,
		StatusId:        int(requirement.StatusId),
		Description:     requirement.Description,
		ExecutorId:      requirement.ExecutorId.String(),
		ProjectId:       requirement.ProjectId.String(),
		CreatedAt:       requirement.CreatedAt,
		UpdatedAt:       requirement.UpdatedAt,
		SpecificationId: specId,
	}
}
