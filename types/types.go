package types

type OBUData struct {
	OBUID int     `json:"obuId"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

type CalculatedDistance struct {
	OBUID    int     `json:"obuID"`
	Distance float64 `json:"distance"`
	Date     string  `json:"date"`
}

type Invoice struct {
	OBUID         int     `json:"obuID"`
	TotalDistance float64 `json:"totalDistance"`
	TotalAmount   float64 `json:"totalAmount"`
	Date          string  `json:"date"`
}
