package main

import (
	"fmt"
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

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1998, length("4m37s")},
	{"Go Ahead", "Alicia Key", "GRAPH", 2007, length("1m52s")},
	{"Python", "Martin Solveig", "BUPT", 2011, length("4m36s")},
	{"Java", "Shengjun Yu", "Smash", 2024, length("5m16s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
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

type byArtist []*Track

func (x byArtist) Len() int { return len(x) }

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (c customSort) Len() int {
	return len(c.t)
}

func (c customSort) Less(i, j int) bool {
	return c.less(c.t[i], c.t[j])
}

func (c customSort) Swap(i, j int) {
	c.t[i], c.t[j] = c.t[j], c.t[i]
}

func main() {
	printTracks(tracks)

	sort.Sort(byArtist(tracks))

	fmt.Println("===========Sort=============")

	printTracks(tracks)

	fmt.Println("=========Reverse===========")

	sort.Sort(sort.Reverse(byArtist(tracks)))

	printTracks(tracks)

	fmt.Println("=========customSort===========")

	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)
	values := []int{3, 1, 4, 6, 2, 8, 9, 0}
	fmt.Println(sort.IntsAreSorted(values))
	sort.Ints(values)
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	a := make([]int, 0)

	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()

	c := a

	b := append(a, 13)

	b = append(b, 14)

	fmt.Println(b)

	c = append(c, 17)

	
	fmt.Println(c)

}
