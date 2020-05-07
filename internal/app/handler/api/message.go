package api

import (
	"encoding/json"
	"github.com/imansuparman/ngobrol/internal/app/commons"
	"github.com/imansuparman/ngobrol/internal/app/handler"
	payloads "github.com/imansuparman/ngobrol/internal/app/payload"
	"github.com/kitabisa/perkakas/v2/log"
	"io/ioutil"
	"net/http"
)

// MessageHandler object for message handler
type MessageHandler struct {
	handler.HandlerOption
	http.Handler
}

func (h *MessageHandler) Index(w http.ResponseWriter, r *http.Request) (data interface{}, pageToken *string, err error) {
	data, err = h.Services.Message.GetAll()
	if err != nil {
		h.Logger.AddMessage(log.ErrorLevel, err)
		return
	}
	return
}

func (h *MessageHandler) Create(w http.ResponseWriter, r *http.Request) (data interface{}, pageToken *string, err error) {
	var body []byte
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		h.Options.Logger.AddMessage(log.ErrorLevel, err)
		err = commons.ErrInvalidBody
		return
	}

	var request payloads.MessageRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		h.Options.Logger.AddMessage(log.ErrorLevel, err)
		err = commons.ErrInvalidBody
		return
	}

	_, err = payloads.GeneralValidator(request, true)
	if err != nil {
		h.Logger.AddMessage(log.ErrorLevel, err)
		err = commons.ErrInvalidBody
		return
	}

	data, err = h.Services.Message.Create(request)
	if err != nil {
		h.Options.Logger.AddMessage(log.ErrorLevel, err)
		return
	}

	h.Websocket.Emit("connection", data)

	return
}
