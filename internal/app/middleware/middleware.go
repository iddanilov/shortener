package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type appHandler func(context *gin.Context) error

func Middleware(h appHandler) gin.HandlerFunc {
	return func(context *gin.Context) {
		var appErr *AppError
		w := context.Writer
		err := h(context)
		if err != nil {
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusBadRequest)
					_, err := w.Write(ErrNotFound.Marshal())
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
					}
					return
				}
				err = err.(*AppError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(ErrNotFound.Marshal())
			}
			w.WriteHeader(http.StatusTeapot)
			w.Write(systemError(err).Marshal())
		}

	}

}
