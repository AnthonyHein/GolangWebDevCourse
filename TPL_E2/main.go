package main

import (
    "text/template"
    "os"
    "log"
)

type hotel struct {
    Name string
    Address string
    City string
    Zip int
    Region string
}

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
    hotels := []hotel{
        hotel{
            Name : "Mohegan Sun",
            Address : "101 Hullfish Lane",
            City : "Las Vegas",
            Zip : 8675,
            Region : "Southern",
        },
        hotel{
            Name : "Wachovia",
            Address : "1600 Penn Ave",
            City : "Sacramento",
            Zip : 309,
            Region : "Northern",
        },
        hotel{
            Name : "Marriot",
            Address : "42 Wallaby Way",
            City : "San Andreas",
            Zip : 06706,
            Region : "Central",
        },
    }

    err := tpl.Execute(os.Stdout, hotels)
    if err != nil {
        log.Println(err)
    }
}
