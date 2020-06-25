package main

import (
    "text/template"
    "log"
    "os"
)

type menu struct {
    Breakfast []item
    Lunch []item
    Dinner []item
}

type item struct {
    Name string
}

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

    m := menu{
        Breakfast : []item{
            item{
                "Orange Juice",
            },
            item{
                "Milk",
            },
            item{
                "Eggs",
            },
            item{
                "Toast",
            },
        },
        Lunch : []item{
            item{
                "Turkey Sandwich",
            },
            item{
                "Grilled Cheese",
            },
            item{
                "Chicken Noodle Soup",
            },
            item{
                "Hamburger",
            },
        },
        Dinner : []item{
            item{
                "Filet Mignon",
            },
            item{
                "Crab Cakes",
            },
            item{
                "Spaghetti",
            },
            item{
                "Grilled Chicken",
            },
        },
    }

    err := tpl.Execute(os.Stdout, m)
    if err != nil {
        log.Fatalln(err)
    }
}
