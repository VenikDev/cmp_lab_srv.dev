package controllers

import (
	"github.com/charmbracelet/log"
	"github.com/valyala/fasthttp"
	"parser_labs/models/store"
)

func ProcessRequest(params store.StrStore) (err error) {
	url, err := GetLink(params)
	if err != nil {
		return err
	}

	var body []byte
	// client
	{
		client := &fasthttp.Client{}
		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		// Set the request method and URL
		req.SetRequestURI(url)

		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		err := client.Do(req, resp)
		if err != nil {
			log.Error("[process_req/to_lab] " + err.Error())
			return err
		}

		body = resp.Body()
	}

	// parse
	{

	}

	return nil
}
