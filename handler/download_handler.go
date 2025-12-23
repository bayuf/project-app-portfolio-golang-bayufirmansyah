package handler

import (
	"net/http"

	"go.uber.org/zap"
)

type DownloadHandler struct {
	Logger *zap.Logger
}

func NewDownloadHandler(log *zap.Logger) *DownloadHandler {
	return &DownloadHandler{
		Logger: log,
	}
}

func (h *DownloadHandler) DownloadCV(w http.ResponseWriter, r *http.Request) {
	filePath := "public/files/cv.pdf"
	if filePath == "" {
		h.Logger.Info("filepath for cv pdf is nil")
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", `attachment; filename="cv.pdf"`)

	http.ServeFile(w, r, filePath)
}

func (h *DownloadHandler) SeeCV(w http.ResponseWriter, r *http.Request) {
	filePath := "public/files/cv.pdf"
	if filePath == "" {
		h.Logger.Info("filepath for cv pdf is nil")
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", `inline; filename="cv.pdf"`)

	http.ServeFile(w, r, filePath)
}
