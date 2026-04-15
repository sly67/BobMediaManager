package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"

	"bobmediamanager/config"
	"bobmediamanager/service"
)

type Handler struct {
	cfg *config.Config
}

// --- Config ---

func (h *Handler) GetConfig(w http.ResponseWriter, r *http.Request) {
	jsonOK(w, h.cfg)
}

func (h *Handler) PutConfig(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(h.cfg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := config.Save(h.cfg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonOK(w, h.cfg)
}

// --- Sources ---

func (h *Handler) GetSources(w http.ResponseWriter, r *http.Request) {
	jsonOK(w, h.cfg.CustomSources)
}

func (h *Handler) PostSource(w http.ResponseWriter, r *http.Request) {
	jsonOK(w, map[string]string{"status": "not implemented"})
}

func (h *Handler) PutSource(w http.ResponseWriter, r *http.Request) {
	_ = chi.URLParam(r, "id")
	jsonOK(w, map[string]string{"status": "not implemented"})
}

func (h *Handler) DeleteSource(w http.ResponseWriter, r *http.Request) {
	_ = chi.URLParam(r, "id")
	jsonOK(w, map[string]string{"status": "not implemented"})
}

func (h *Handler) TestSource(w http.ResponseWriter, r *http.Request) {
	_ = chi.URLParam(r, "slug")
	jsonOK(w, map[string]string{"status": "not implemented"})
}

// --- Scan ---

func (h *Handler) PostScan(w http.ResponseWriter, r *http.Request) {
	var req service.ScanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	files, err := service.Scan(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonOK(w, files)
}

// --- Match & Search ---

func (h *Handler) PostMatch(w http.ResponseWriter, r *http.Request) {
	jsonOK(w, map[string]string{"status": "not implemented"})
}

func (h *Handler) PostSearch(w http.ResponseWriter, r *http.Request) {
	jsonOK(w, map[string]string{"status": "not implemented"})
}

func (h *Handler) GetMetadata(w http.ResponseWriter, r *http.Request) {
	_ = chi.URLParam(r, "source")
	_ = chi.URLParam(r, "id")
	jsonOK(w, map[string]string{"status": "not implemented"})
}

func (h *Handler) PutFileID(w http.ResponseWriter, r *http.Request) {
	_ = chi.URLParam(r, "fileID")
	jsonOK(w, map[string]string{"status": "not implemented"})
}

// --- Rename ---

func (h *Handler) PostRenamePreview(w http.ResponseWriter, r *http.Request) {
	jsonOK(w, map[string]string{"status": "not implemented"})
}

func (h *Handler) PostRenameExecute(w http.ResponseWriter, r *http.Request) {
	jsonOK(w, map[string]string{"status": "not implemented"})
}

// --- Companions ---

func (h *Handler) PostNFO(w http.ResponseWriter, r *http.Request) {
	jsonOK(w, map[string]string{"status": "not implemented"})
}

func (h *Handler) PostImages(w http.ResponseWriter, r *http.Request) {
	jsonOK(w, map[string]string{"status": "not implemented"})
}

func (h *Handler) PostSubtitlesSearch(w http.ResponseWriter, r *http.Request) {
	jsonOK(w, map[string]string{"status": "not implemented"})
}

func (h *Handler) PostSubtitlesDownload(w http.ResponseWriter, r *http.Request) {
	jsonOK(w, map[string]string{"status": "not implemented"})
}

func (h *Handler) PostVerify(w http.ResponseWriter, r *http.Request) {
	jsonOK(w, map[string]string{"status": "not implemented"})
}

// --- History ---

func (h *Handler) GetHistory(w http.ResponseWriter, r *http.Request) {
	jsonOK(w, []interface{}{})
}

func (h *Handler) PostHistoryUndo(w http.ResponseWriter, r *http.Request) {
	_ = chi.URLParam(r, "id")
	jsonOK(w, map[string]string{"status": "not implemented"})
}

// --- WebSocket ---

func (h *Handler) WSProgress(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if err := conn.WriteMessage(mt, msg); err != nil {
			break
		}
	}
}

// --- helpers ---

func jsonOK(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

// unused placeholder kept so websocket import stays used in this file
var _ = (*websocket.Conn)(nil)
