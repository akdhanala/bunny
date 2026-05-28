package controllers

type BunnyController struct{}

func NewBunnyController() *BunnyController {
	return &BunnyController{}
}

func (c *BunnyController) ResolveDestination() string {
	return "https://www.google.com/search?q=what+is+golang"
}