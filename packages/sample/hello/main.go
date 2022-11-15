package main

import (
	"time"

	"github.com/hebcal/hebcal-go/hdate"
)

const iso8601 = "2006-01-02"

func Main(args map[string]interface{}) map[string]interface{} {
	t := time.Now()
	dt, ok := args["dt"].(string)
	if ok {
		tmp, err := time.Parse(iso8601, dt)
		if err == nil {
			t = tmp
		}
	}
	msg := make(map[string]interface{})
	msg["headers"] = map[string]string{
		"Content-Type":  "application/json",
		"Cache-Control": "public",
	}

	hd := hdate.FromTime(t)

	msg["body"] = map[string]string{
		"date":  t.Format(iso8601),
		"hdate": hd.String(),
	}
	return msg
}
