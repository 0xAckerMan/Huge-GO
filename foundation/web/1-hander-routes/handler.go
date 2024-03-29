package main
import(
    "fmt"
    "net/http"
)

func (app *Application) healthcheck(w http.ResponseWriter, r *http.Request){
    if r.Method != http.MethodGet {
        http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
        return
    }
    fmt.Fprintln(w,"Status: Active")
    fmt.Fprintf(w, "Environment: %s\n", "dev")
    fmt.Fprintf(w, "Version: %s\n", version)
}
