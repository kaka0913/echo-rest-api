package validator

import(
	"go-rest-api/model"

	valdation "github.com/go-ozzo/ozzo-validation"
)

type ITaskValdator interface {
	TaskValdate(task model.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValdator {
	return &taskValidator{}
}	

func (tv *taskValidator) TaskValdate(task model.Task) error {
	return valdation.ValidateStruct(&task,
		valdation.Field(
				&task.Title, 
				valdation.Required.Error("title is required"),
				valdation.Length(1, 10).Error("limited max 10 characters"),
		),
	)
}