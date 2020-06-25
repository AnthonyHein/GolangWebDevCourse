package main

import (
    "os"
    "log"
    "text/template"
    "encoding/csv"
    "strconv"
    "sort"
)

var tpl *template.Template

type stat struct {
    Open float64
    Close float64
}

type ByDiff []stat

func (b ByDiff) Len() int { return len(b) }

func (b ByDiff) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func (b ByDiff) Less(i, j int) bool { return b[i].Close - b[i].Open < b[j].Close - b[j].Open }

func diff(s stat) float64 {
    return s.Close - s.Open
}

var fm = template.FuncMap{
    "diff" : diff,
}

func init() {
    tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {

    src, err := os.Open("table.csv")
    if err != nil {
        log.Fatalln(err)
    }
    defer src.Close()

    rdr := csv.NewReader(src)
    rows, err := rdr.ReadAll()
    if err != nil {
        log.Fatalln(err)
    }

    records := make([]stat, 0, len(rows))

    for i, row := range rows {
        if i == 0 {
            continue
        }

        open, err := strconv.ParseFloat(row[1], 64)
        if err != nil {
            log.Fatalln(err)
        }
        close, err := strconv.ParseFloat(row[4], 64)
        if err != nil {
            log.Fatalln(err)
        }

        records = append(records, stat{
			Open: open,
			Close: close,
		})
    }

    sort.Sort(sort.Reverse(ByDiff(records)))

    err = tpl.Execute(os.Stdout, records)
    if err != nil {
        log.Fatalln(err)
    }
}
