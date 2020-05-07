package handler

import (
	"fmt"
	"github.com/kitabisa/perkakas/v2/log"
	"html/template"
	"net/http"
	"os"
)

type IndexHandler struct {
	HandlerOption
	http.Handler
}

func (h *IndexHandler) Index(w http.ResponseWriter, r *http.Request) {
	dir, err := os.Getwd()
	if err != nil {
		h.Logger.AddMessage(log.ErrorLevel, err)
		return
	}

	h.Logger.AddMessage(log.InfoLevel, dir)
	var tmpl = template.Must(template.ParseFiles(
		fmt.Sprintf("%v/%v", dir, "views/Index.tmpl"),
		fmt.Sprintf("%v/%v", dir, "views/components/Header.tmpl"),
		fmt.Sprintf("%v/%v", dir, "views/components/Footer.tmpl"),
	))

	data, err := h.Services.Message.GetAll()
	if err != nil {
		h.Logger.AddMessage(log.ErrorLevel, err)
		return
	}

	fmt.Println("Data", data)

	err = tmpl.ExecuteTemplate(w, "Index", data)
	if err != nil {
		h.Logger.AddMessage(log.ErrorLevel, err)
		return
	}
}