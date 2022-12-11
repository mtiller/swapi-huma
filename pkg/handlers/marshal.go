package handlers

import (
	"encoding/json"
	"strings"

	"github.com/mtiller/go-claxon"
)

type Preamble struct {
	Schema string `json:"$schema,omitempty"`
	Self   string `json:"$self,omitempty"`
}

type ShortLink struct {
	Href  string `json:"href"`
	Title string `json:"title,omitempty"`
	Type  string `json:"type,omitempty"`
}

type LinksByRel map[string][]ShortLink

type Postamble struct {
	Links   LinksByRel      `json:"$links,omitempty"`
	Actions []claxon.Action `json:"$actions,omitempty"`
}

func Marshal(v interface{}, c claxon.Claxon) ([]byte, error) {
	pre := Preamble{
		Schema: c.Schema,
		Self:   c.Self,
	}
	preamble, err := json.Marshal(pre)
	if err != nil {
		return preamble, err
	}
	links := LinksByRel{}
	for _, link := range c.Links {
		rel, exists := links[link.Rel]
		if !exists {
			rel = []ShortLink{}
		}
		rel = append(rel, ShortLink{
			Title: link.Title,
			Href:  link.Href,
			Type:  link.Type,
		})
		links[link.Rel] = rel
	}
	post := Postamble{
		Links:   links,
		Actions: c.Actions,
	}
	postamble, err := json.Marshal(post)
	data, err := json.Marshal(v)
	if err != nil {
		return data, err
	}
	start := string(preamble[1 : len(preamble)-1])
	middle := string(data[1 : len(data)-1])
	end := string(postamble[1 : len(postamble)-1])

	parts := []string{}
	if strings.TrimSpace(start) != "" {
		parts = append(parts, start)
	}
	if strings.TrimSpace(middle) != "" {
		parts = append(parts, middle)
	}
	if strings.TrimSpace(end) != "" {
		parts = append(parts, end)
	}
	return []byte("{" + strings.Join(parts, ",") + "}"), nil
}
