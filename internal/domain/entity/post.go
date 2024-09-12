package entity

type Post struct {
	ID         string
	User       User
	Message    string
	Coordinate Coordinate
	CreatedAt  int64 // Unix time
}

type Coordinate struct {
	Latitude  float64
	Longitude float64
}
