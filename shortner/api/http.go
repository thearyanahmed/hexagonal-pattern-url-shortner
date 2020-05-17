package api

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"

	"github.com/thearyanahmed/url-shortner/serializer/json"
	"github.com/thearyanahmed/url-shortner/serializer/msgpack"
	"github.com/thearyanahmed/url-shortner/shortener"
	"github.com/thearyanahmed/url-shortner/shortner"
)

type RedirectHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	redirectService shortner.RedirectService
}

func newHandler(handlerService shortner.RedirectSerializer) RedirectHandler {
	return &handlerService{ redirectService: redirectService}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type",contentType)
	w.WriteHeader(statusCode)

	_, err := w.Write(body)

	if err != nil {
		log.Println(err)
	}
}

func (h *handler) serializer(contentType string) shortener.RedirectSerializer {
	if contentType == "appliaction/x-msgpack" {
		return &msgpack.Redirect{}
	}

	return &json.Redirect{}
}

func (h *handler) Get(w http.ResponseWriter, r http.Request) {
	code := chi.URLParam(r,"code")

	redirect, err := h.redirectService.Find(code)

	if err != nil {
		if errors.Cause(err) == shortner.RedirectNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r , redirect.URL, http.StatusMovedPermanently)
}

func (h *handler) Post (w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		return
	}

	redirect, err := h.serializer(contentType).Decode(body)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = h.redirectService.Store(redirect)

	if err != nil {
		if errors.Cause(err) == shortner.InvalidRedirect {
			http.Error(w, http.StatusText(http.StatusBadRequest),http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseBody, err := h.serializer(contentType).Encode(redirect)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(w,contentType,responseBody,http.StatusCreated)
}