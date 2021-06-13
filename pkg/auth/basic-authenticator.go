package auth

import "net/http"

// Validate HTTP request credentials
func (h *HtpasswdFile) Validate(req *http.Request, res http.ResponseWriter) bool {
	user, passwd, ok := req.BasicAuth()
	if !ok || !h.validateCredentials(user, passwd) {
		res.Header().Set("WWW-Authenticate", `Basic realm="Ah ah ah, you didn't say the magic word"`)
		return false
	}
	return true
}

func (h *HtpasswdFile) Issuer() string {
	return "basic"
}
