package service

import (
	"github.com/FlyCynomys/tools/log"
)

//all assert should take in handles ,not here

type Result struct {
	Status      int         `json:"status,omitempty"`
	Description string      `json:"description,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

type ResultList struct {
	Status      int           `json:"status,omitempty"`
	Description string        `json:"description,omitempty"`
	Data        []interface{} `json:"data,omitempty"`
}

type ResultMap struct {
	Status      int                    `json:"status,omitempty"`
	Description string                 `json:"description,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
}

func Init() {
	log.Info("service init start")

	log.Info("service init over")
}
