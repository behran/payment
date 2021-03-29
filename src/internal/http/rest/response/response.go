package response

import (
	"encoding/json"
	"net/http"
	"os"

	"payment/internal/contracts"

	"github.com/valyala/fasthttp"
)

//Response ...
type Response struct {
	Meta     Meta        `json:"meta"`
	JSONAPI  JSONAPI     `json:"jsonapi"`
	Response interface{} `json:"data,omitempty"`
	Errors   []Err       `json:"errors,omitempty"`
}

//Response objects ...
type (
	//Meta ...
	Meta struct {
		Status bool `json:"status"`
		Code   int  `json:"status_code"`
	}
	//JSONAPI ...
	JSONAPI struct {
		Version string `json:"version"`
	}
	//Err ...
	Err struct {
		Tittle string `json:"title"`
		Detail string `json:"detail"`
	}
)

//Version ...
var Version = os.Getenv("API_VERSION")

//OK ...
func OK(body interface{}, ctx *fasthttp.RequestCtx) {
	var payload = Response{
		Meta: Meta{
			Code:   http.StatusOK,
			Status: true,
		},
		JSONAPI: JSONAPI{
			Version: Version,
		},
		Response: body,
	}
	send(payload, http.StatusOK, ctx)
}

//Error ...
func Error(err error, ctx *fasthttp.RequestCtx) {
	e := err.(contracts.IPaymentError)
	payload := Response{
		Meta: Meta{
			Code:   e.StatusCode(),
			Status: false,
		},
		JSONAPI: JSONAPI{
			Version: Version,
		},
		Errors: []Err{
			{
				Tittle: e.Error(),
				Detail: e.Details(),
			},
		},
	}
	send(payload, e.StatusCode(), ctx)
}

func send(r Response, code int, ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(code)
	json.NewEncoder(ctx).Encode(r)
}
