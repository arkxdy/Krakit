package proxy

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func New(target string) (*httputil.ReverseProxy, error) {
	u, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	p := httputil.NewSingleHostReverseProxy(u)

	p.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)

		_ = json.NewEncoder(w).Encode(map[string]any{
			"error": "service unavailable",
		})
	}

	return p, nil
}
