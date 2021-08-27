package models

type TopSecret struct {
	Satellites []Satellite `json:"satellites,omitempy"`
}

type Satellite struct {
	Name     string   `json:"name,omitempy"`
	Distance float32  `json:"distance,omitempy"`
	Message  []string `json:"message,omitempy"`
}

type ResponseTopSecret struct {
	Position Position `json:"position,omitempy"`
	Message  string   `json:"message,omitempy"`
}

type Position struct {
	X float32 `json:"x,omitempy"`
	Y float32 `json:"y,omitempy"`
}

type ResponseError struct {
	Er string `json:"error,omitempy"`
}

type ResponseSplit map[string]interface{}

var SatellitesBD []Satellite
