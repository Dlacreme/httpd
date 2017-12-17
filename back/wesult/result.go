package wesult

import (
	"encoding/json"
	"net/http"

	"github.com/Dlacreme/httpd/back/flight"
	"github.com/Dlacreme/httpd/back/werror"
	"github.com/Dlacreme/httpd/view"
)

// New will create a new Result type
func New(output IOutput, err *werror.Error) Result {
	return Result{output, err}
}

// Result will be return by all YANA-Core query
type Result struct {
	Output IOutput
	Error  *werror.Error
}

// IOutput is used a output of all YANA-core query
type IOutput interface {
}

// ToJson will format the result as Json file
func (res *Result) ToJson(w http.ResponseWriter) {
	if res.Error != nil {
		w.WriteHeader(res.Error.Code)
		w.Write([]byte(res.Error.Message))
		return
	}
	json.NewEncoder(w).Encode(res.Output)
}

// RenderView will display appropriate view
func (res *Result) RenderView(w http.ResponseWriter, r *http.Request, v *view.Info) {
	c := flight.Context(w, r)
	v.Vars["user"] = c.User
	v.Vars["model"] = res.Output
	v.Vars["error"] = res.Error
	v.Render(w, r)
}
