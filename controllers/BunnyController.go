package controllers

import (
	"github.com/akdhanala/bunny/entity"
)

type BunnyController struct{
	cfg *entity.BunnyConfig
}

func NewBunnyController(cfg *entity.BunnyConfig) *BunnyController {	
	return &BunnyController{
		cfg: cfg,
	}
}

func (c *BunnyController) ResolveDestination(r entity.ResolveDestinationRequest) (string, error) {
	command := c.locateConfig(r.Command)
	if (command == nil) {
		return "", &entity.CommandNotFound{
			Command: r.Command,
		}
	}

	return c.queryOnlyResolver(r, command)
}

func (c *BunnyController) locateConfig(
	command string, 
) *entity.BunnyCommandOpts {
	if (c.cfg == nil) {
		return nil
	}

	for _, bunnyCommandOpts := range c.cfg.BunnyCommands {
		for _, alias := range bunnyCommandOpts.Aliases {
			if (command == alias) {
				return &bunnyCommandOpts
			}
		}
	}

	return nil
}

func (c *BunnyController) queryOnlyResolver(
	r entity.ResolveDestinationRequest, 
	command *entity.BunnyCommandOpts,
) (string, error) {
	if (len(r.Query) == 0) {
		return command.BaseURL, nil
	}
	return command.BaseURL + command.QueryPath + r.Query, nil
}
