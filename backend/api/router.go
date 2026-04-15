package api

import (
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/websocket"

	"bobmediamanager/config"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func NewRouter(cfg *config.Config, staticFS fs.FS) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)

	h := &Handler{cfg: cfg}

	r.Route("/api", func(r chi.Router) {
		// Config
		r.Get("/config", h.GetConfig)
		r.Put("/config", h.PutConfig)

		// Sources
		r.Get("/sources", h.GetSources)
		r.Post("/sources", h.PostSource)
		r.Put("/sources/{id}", h.PutSource)
		r.Delete("/sources/{id}", h.DeleteSource)
		r.Post("/sources/{slug}/test", h.TestSource)

		// Scan & match
		r.Post("/scan", h.PostScan)
		r.Post("/match", h.PostMatch)
		r.Post("/search", h.PostSearch)
		r.Get("/metadata/{source}/{id}", h.GetMetadata)
		r.Put("/files/{fileID}/id", h.PutFileID)

		// Operations
		r.Post("/rename/preview", h.PostRenamePreview)
		r.Post("/rename/execute", h.PostRenameExecute)

		// Companions
		r.Post("/nfo", h.PostNFO)
		r.Post("/images", h.PostImages)
		r.Post("/subtitles/search", h.PostSubtitlesSearch)
		r.Post("/subtitles/download", h.PostSubtitlesDownload)
		r.Post("/verify", h.PostVerify)

		// History
		r.Get("/history", h.GetHistory)
		r.Post("/history/{id}/undo", h.PostHistoryUndo)
	})

	// WebSocket
	r.Get("/ws/progress", h.WSProgress)

	// Static frontend (SPA)
	r.Handle("/*", http.FileServer(http.FS(staticFS)))

	return r
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
