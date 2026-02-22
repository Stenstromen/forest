package models

import "time"

/*
{
  "occurred_at": "2026-02-23T08:30:00Z",

  "distance": { "value": 10, "unit": "km" },
  "duration": { "value": 50, "unit": "min" },

  "calories": 650
}
*/

type Distance struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

type Derived struct {
	PaceSecondsPerKm   float64 `json:"pace_seconds_per_km"`
	PaceSecondsPerMile float64 `json:"pace_seconds_per_mile"`
}

type Run struct {
	OccurredAt time.Time `json:"occurred_at"`
	Distance   Distance  `json:"distance"`
	Duration   Distance  `json:"duration"`
	Calories   float64   `json:"calories"`
}

type RunResponse struct {
	ID         string    `json:"id"`
	OccurredAt time.Time `json:"occurred_at"`
	Distance   Distance  `json:"distance_m"`
	Duration   Distance  `json:"duration_s"`
	Calories   float64   `json:"calories_kcal"`
	Derived    Derived   `json:"derived"`
}
