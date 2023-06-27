package cci

type State struct {
	Name string  `json:"name"`
	Code string  `json:"code"`
	City []*City `json:"city"`
}

type City struct {
	Name   string    `json:"name"`
	Code   string    `json:"code"`
	Region []*Region `json:"region,omitempty"`
}

type Region struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}
