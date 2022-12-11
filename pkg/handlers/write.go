package handlers

import (
	"net/http"

	"github.com/danielgtaylor/huma"
	"github.com/mtiller/go-claxon"
	"github.com/mtiller/rfc8288"
)

func WriteModel(status int, ctx huma.Context, v interface{}, c claxon.Claxon) {
	links, err := claxon.ToRFC8288Links(c)
	if err != nil {
		ctx.WriteError(http.StatusInternalServerError, err.Error())
		return
	}

	val := rfc8288.LinkHeaderValue(links...)
	ctx.Header().Add("Link", val)

	ctx.WriteModel(status, v)
}
