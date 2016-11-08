package auth

import (
	"net/http"

	"github.com/thoas/stats"
)

func (p *Engine) getAdminStatus(w http.ResponseWriter, r *http.Request) {
	stats := stats.New().Data()
	p.Render.JSON(w, http.StatusOK, stats)
}
