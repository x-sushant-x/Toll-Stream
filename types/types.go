package types

type OBUData struct {
	OBUID int     `json:"obuId"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

type CalculatedDistance struct {
	OBUID     int     `json:"obuID"`
	Distance  float64 `json:"distance"`
	Timestamp int64   `json:"timestamp"`
}
