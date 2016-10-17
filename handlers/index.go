package handlers

import (
	"fmt"
	"net/http"
)

// Index is the default Endpoint for ecofox insights user service
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ecofox insights user service")
	fmt.Fprintln(w, "---")
	fmt.Fprintln(w, " ")
	fmt.Fprintln(w, "Interactive Documentation available at: /ui")
}
