package transport

import (
	"net/http"
	"time"
)

var HTTP = &http.Client{
	Timeout: 60 * time.Second,
}
