package payloads

import (
	"github.com/asaskevich/govalidator"
)

type MessageRequest struct {
	Message string `json:"message" valid:"required"`
}

func GeneralValidator(s interface{}, isRequired bool) (bool, error) {
	govalidator.SetFieldsRequiredByDefault(isRequired)
	return govalidator.ValidateStruct(s)
}
