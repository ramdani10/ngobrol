package repository

import (
	"github.com/im7mortal/kmutex"
	"github.com/imansuparman/ngobrol/internal/app/model"
)

// IMessageRepo interface for message repository
type IMessageRepo interface {
	GetAll() (messages []model.Message, err error)
	Create(message *model.Message) (err error)
}

type messageRepo struct {
	opt    Option
	kmutex *kmutex.Kmutex
}

// NewMessageRepository initiate message repository
func NewMessageRepository(opt Option) IMessageRepo {
	return &messageRepo{
		opt:    opt,
		kmutex: kmutex.New(),
	}
}

func (m *messageRepo) GetAll() (messages []model.Message, err error) {
	_, err = m.opt.DbPostgre.Select(&messages, "SELECT * FROM messages")
	return
}

func (m *messageRepo) Create(message *model.Message) (err error) {
	err = m.opt.DbPostgre.Insert(message)
	return
}
