package service

import (
	"github.com/imansuparman/ngobrol/internal/app/commons"
	"github.com/imansuparman/ngobrol/internal/app/model"
	payloads "github.com/imansuparman/ngobrol/internal/app/payload"
	"github.com/kitabisa/perkakas/v2/log"
	"time"
)

// IMessage interface for message service
type IMessage interface {
	GetAll() (messages[]payloads.MessageResponse, err error)
	Create(request payloads.MessageRequest) (response payloads.MessageResponse, err error)
}

type message struct {
	opt Option
}

// NewMessage create message service instance with option as param
func NewMessage(opt Option) IMessage {
	return &message{
		opt: opt,
	}
}

func (h *message) GetAll() (response []payloads.MessageResponse, err error) {
	var messages []model.Message
	messages, err = h.opt.Repository.Message.GetAll()
	if err != nil {
		h.opt.Logger.AddMessage(log.ErrorLevel, err)
		return
	}

	for _, message := range messages {
		response = append(response, payloads.MessageResponse{
			ID:      message.ID,
			Message: message.Message,
			Created: message.Created,
			Updated: message.Updated,
		})
	}
	return
}

func (h *message) Create(request payloads.MessageRequest) (response payloads.MessageResponse, err error) {
	message := model.Message{
		ID:             commons.GenerateUUID(),
		Message:        request.Message,
		Created:        time.Now(),
		Updated:        time.Now(),
	}
	err = h.opt.Repository.Message.Create(&message)
	if err != nil {
		h.opt.Logger.AddMessage(log.ErrorLevel, err)
		return
	}

	response = payloads.MessageResponse{
		ID:      message.ID,
		Message: message.Message,
		Created: message.Created,
		Updated: message.Updated,
	}

	return
}