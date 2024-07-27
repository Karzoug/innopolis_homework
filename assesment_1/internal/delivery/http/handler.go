package http

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"assesment_1/internal/model"
	"assesment_1/internal/token"

	"github.com/rs/zerolog"
)

type handler struct {
	tv      token.Validator
	service FileMessageService
	logger  zerolog.Logger
}

func (h *handler) CreateFileMessage(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Warn().Err(err).Msg("read body error")
		http.Error(w, "Read body error", http.StatusInternalServerError)
		return
	}

	var msg model.Message
	if err := json.Unmarshal(body, &msg); err != nil {
		h.logger.Warn().Err(err).Msg("unmarshal body error")
		http.Error(w, "Invalid JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if !h.tv.Validate(msg.Token) {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	if err := h.service.Create(r.Context(), msg); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w,
				http.StatusText(http.StatusServiceUnavailable),
				http.StatusServiceUnavailable) // it's a service error, not a client wrong request, so TooManyRequests is not suitable
		}
		return
	}
}
