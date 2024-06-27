package models

type Coords struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Speed     float64 `json:"speed"`
	Heading   float64 `json:"heading"`
	Accuracy  float64 `json:"accuracy"`
	Altitude  float64 `json:"altitude"`
}
