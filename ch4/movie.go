package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		Movie{Title: "Go", Year: 2021, Color: false,
			Actors: []string{"吴京", "叶凡"}},
		Movie{Title: "Python", Year: 2022, Color: true,
			Actors: []string{"姬紫月", "庞博"}},
		Movie{Title: "C", Year: 2023, Color: false,
			Actors: []string{"无始大帝", "元始天尊"}},
		Movie{Title: "C++", Year: 2024},
		Movie{Title: "Java", Year: 2025},
	}
	data1, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Println(string(data1))
	data2, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Println(string(data2))
	var titles []struct{ Title string }
	if err := json.Unmarshal(data2, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
