package routers

import "net/http"

type MethodRouter map[string]http.HandlerFunc

func (m MethodRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, ok := m[r.Method]
	if !ok {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	handler(w, r)
}
