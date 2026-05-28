package controllers

import "github.com/akdhanala/bunny/entity"

type BunnyController struct{}

func NewBunnyController() *BunnyController {
	return &BunnyController{}
}

func (c *BunnyController) ResolveDestination(r entity.ResolveDestinationRequest) string {
	if (len(r.Query) == 0) {
		return "https://www.google.com/"
	}
	return "https://www.google.com/search?q=" + r.Query
}