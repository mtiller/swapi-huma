package handlers

import (
	"net/http"

	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/negotiation"
	"github.com/mtiller/go-claxon"
	"github.com/mtiller/rfc8288"
)

func WriteModel(status int, accept string, ctx huma.Context, v interface{}, c claxon.Claxon) {
	ct := "application/json"
	if accept != "" {
		best := negotiation.SelectQValue(accept, []string{
			"application/claxon+json",
			"application/json",
		})

		if best != "" {
			ct = best
		}
	}

	ctx.Header().Set("Content-Type", ct)
	if ct == "application/claxon+json" {
		bytes, err := Marshal(v, c)
		if err != nil {
			ctx.WriteError(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.Write(bytes)
	} else {
		links, err := claxon.ToRFC8288Links(c)
		if err != nil {
			ctx.WriteError(http.StatusInternalServerError, err.Error())
			return
		}

		val := rfc8288.LinkHeaderValue(links...)
		ctx.Header().Add("Link", val)

		ctx.WriteModel(status, v)
	}
}
