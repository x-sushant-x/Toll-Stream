package main

import (
	"math"

	"github.com/sushant102004/Traffic-Toll-Microservice/types"
)

type CalculatorServicer interface {
	CalculateDistance(types.OBUData) (int, error)
}

type CalculateService struct {
	points [][]float64
}

func NewCalculateService() *CalculateService {
	return &CalculateService{
		points: make([][]float64, 0),
	}
}

func (s *CalculateService) CalculateDistance(data types.OBUData) (int, error) {
	distance := 0.0
	if len(s.points) > 0 {
		prevPoint := s.points[len(s.points)-1]
		distance = calculateDistance(prevPoint[0], prevPoint[1], data.Lat, data.Long, "K")
	}
	s.points = append(s.points, []float64{data.Lat, data.Long})
	return int(distance), nil
}

func calculateDistance(lat1, lng1, lat2, lng2 float64, unit ...string) float64 {
	radlat1 := float64(math.Pi * lat1 / 180)
	radlat2 := float64(math.Pi * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}
