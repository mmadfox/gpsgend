package http

import (
	"errors"
	"time"
)

var (
	errBadRequest = errors.New("bad request")
)

type addTrackerRequest struct {
	Model       string                 `json:"model"`
	Color       string                 `json:"color"`
	CustomID    string                 `json:"customId"`
	Descr       string                 `json:"description"`
	Props       map[string]interface{} `json:"props"`
	SkipOffline bool                   `json:"skipOffline"`
	Offline     struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"offline"`
	Elevation struct {
		Min       float64 `json:"min"`
		Max       float64 `json:"max"`
		Amplitude int     `json:"amplitude"`
		Mode      int     `json:"mode"`
	} `json:"elevation"`
	Battery struct {
		Min        float64       `json:"min"`
		Max        float64       `json:"max"`
		ChargeTime time.Duration `json:"chargeTime"`
	} `json:"battery"`
	Speed struct {
		Min       float64 `json:"min"`
		Max       float64 `json:"max"`
		Amplitude int
	} `json:"speed"`
}

type updateTrackerRequest struct {
	Model  string `json:"model"`
	Color  string `json:"color"`
	UserID string `json:"customId"`
	Descr  string `json:"description"`
}

type sensor struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Min       float64 `json:"min"`
	Max       float64 `json:"max"`
	Amplitude int     `json:"amplitude"`
	Mode      int     `json:"mode"`
}
