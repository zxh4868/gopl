package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type TrackTable struct {
	t []*Track
	c []string
}

// 添加一个方法来访问 t 字段
func (tt *TrackTable) GetT() []*Track {
	return tt.t
}

func (x TrackTable) Len() int {
	return len(x.t)
}

func (x TrackTable) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

func (x TrackTable) Less(i, j int) bool {

	for pos := len(x.c) - 1; pos >= 0; pos-- {
		switch x.c[pos] {
		case "title":
			if x.t[i].Title != x.t[j].Title {
				return x.t[i].Title < x.t[j].Title
			}
		case "artist":
			if x.t[i].Artist != x.t[j].Artist {
				return x.t[i].Artist < x.t[j].Artist
			}
		case "album":
			if x.t[i].Album != x.t[j].Album {
				return x.t[i].Album < x.t[j].Album
			}
		case "year":
			if x.t[i].Year != x.t[j].Year {
				return x.t[i].Year < x.t[j].Year
			}
		case "length":
			if x.t[i].Length != x.t[j].Length {
				return x.t[i].Length < x.t[j].Length
			}
		}
	}

	return x.t[i].Title < x.t[j].Title

}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func (x *TrackTable) Click(coloum string) {
	for i, v := range x.c {
		if v == coloum {
			x.c = append(x.c[:i], x.c[i+1:]...)
		}
	}
	x.c = append(x.c, coloum)
	sort.Sort(x)
}

func main() {
	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1998, length("4m37s")},
		{"Go Ahead", "Alicia Key", "GRAPH", 2007, length("1m52s")},
		{"Python", "Martin Solveig", "BUPT", 2011, length("4m36s")},
		{"Java", "Shengjun Yu", "Smash", 2024, length("5m16s")},
	}

	table := TrackTable{tracks, make([]string, 0)}

	tmpl := template.Must(template.ParseFiles("./table.html"))

	// 设置路由

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		head := r.FormValue("column")
		if head == "" {
			head = "title"
		}
		table.Click(head)
		err := tmpl.Execute(w, &tracks)
		if err != nil {
			log.Printf("template error: %s", err)
		}

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)

	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")

	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()

}
