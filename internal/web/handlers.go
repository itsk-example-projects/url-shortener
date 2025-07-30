package web

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"

	"url-shortener/internal/config"
	"url-shortener/internal/models"
	"url-shortener/internal/storage"
	"url-shortener/internal/utils"
)

type Handlers struct {
	Templates *template.Template
	Storage   *storage.LinkStorage
	BaseURL   string
}

func RegisterHandlers(templates *template.Template, storage *storage.LinkStorage) *Handlers {
	return &Handlers{
		Templates: templates,
		Storage:   storage,
		BaseURL:   viper.GetString(config.BaseUrl),
	}
}

func (h *Handlers) Index(w http.ResponseWriter, r *http.Request) {
	// If "/:slug" - redirect
	if r.URL.Path != "/" {
		slug := strings.TrimPrefix(r.URL.Path, "/")
		if link, ok := h.Storage.GetLinkBySlug(slug); ok {
			http.Redirect(w, r, link.Long, http.StatusFound)
			return
		}
		http.NotFound(w, r)
		return
	}

	// If "/" - serve index page
	uid := GetUserID(w, r)
	data := map[string]any{
		"UserLinks": h.Storage.LinksByUser(uid),
		"BaseURL":   h.BaseURL,
		"UserID":    uid,
	}
	if err := h.Templates.ExecuteTemplate(w, "index.html", data); err != nil {
		log.Printf("error rendering index: %v", err)
	}
}

func (h *Handlers) ExportID(w http.ResponseWriter, r *http.Request) {
	uid := GetUserID(w, r)
	w.Header().Set("Content-Disposition", "attachment; filename=\"user-id.txt\"")
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(uid))
}

func (h *Handlers) ImportID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	file, _, err := r.FormFile("user_id_file")
	if err != nil {
		http.Error(w, "Failed to read uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	idBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file content", http.StatusInternalServerError)
		return
	}

	newID := strings.TrimSpace(string(idBytes))
	if newID == "" {
		http.Error(w, "Uploaded file is empty or invalid", http.StatusBadRequest)
		return
	}

	SetUserIDCookie(w, newID)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handlers) Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	uid := GetUserID(w, r)
	rawURL := strings.TrimSpace(r.FormValue("url"))
	if rawURL == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "http://" + rawURL
	}

	slug, err := utils.GenerateSlug(
		func(s string) bool { _, exists := h.Storage.GetLinkBySlug(s); return exists },
		6,
		100,
	)
	if err != nil {
		http.Error(w, "Failed to generate slug", http.StatusInternalServerError)
		return
	}

	link := models.Link{Short: slug, Long: rawURL, UserID: uid}
	if err := h.Storage.AddLink(link); err != nil {
		http.Error(w, "Failed to save link", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	slug := strings.TrimPrefix(r.URL.Path, "/delete/")
	uid := GetUserID(w, r)

	link, ok := h.Storage.GetLinkBySlug(slug)
	if !ok || link.UserID != uid {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if err := h.Storage.DeleteLinkBySlug(slug); err != nil {
		http.Error(w, "Failed to delete link", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handlers) Admin(w http.ResponseWriter, r *http.Request) {
	user, pass := BasicAuth(r)
	if user != viper.GetString(config.ControlPanelUsername) || pass != viper.GetString(config.ControlPanelPassword) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	data := map[string]any{
		"Links":   h.Storage.GetAllLinks(),
		"BaseURL": h.BaseURL,
	}
	if err := h.Templates.ExecuteTemplate(w, "panel.html", data); err != nil {
		log.Printf("error rendering admin: %v", err)
	}
}

func (h *Handlers) AdminDelete(w http.ResponseWriter, r *http.Request) {
	user, pass := BasicAuth(r)
	if user != viper.GetString(config.ControlPanelUsername) || pass != viper.GetString(config.ControlPanelPassword) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	slug := strings.TrimPrefix(r.URL.Path, "/admin/delete/")
	if err := h.Storage.DeleteLinkBySlug(slug); err != nil {
		http.Error(w, "Failed to delete link", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
