package controllers

import "github.com/akdhanala/bunny/entity"

type BunnyController struct{}

func NewBunnyController() *BunnyController {
	return &BunnyController{}
}

func (c *BunnyController) ResolveDestination(r entity.ResolveDestinationRequest) (string, error) {
	if (r.Command != "g") {
		return "", entity.CommandNotFound{
			Command: r.Command,
		}
	}

	if (len(r.Query) == 0) {
		return "https://www.google.com/", nil
	}
	return "https://www.google.com/search?q=" + r.Query, nil
}
