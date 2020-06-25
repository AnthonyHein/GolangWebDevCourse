package main

import (
    "os"
    "text/template"
    "log"
)

var tpl *template.Template

type menu struct {
    Rest string
    Breakfast []item
    Lunch []item
    Dinner []item
}

type item struct {
    Name string
}

func init () {
    tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main () {

    rest := []menu{
        menu{
            Rest : "Pop's Deli",
            Breakfast : []item{
                item{
                    "Eggs",
                },
            },
            Lunch : []item{
                item{
                    "Chicken Noodle Soup",
                },
            },
            Dinner : []item{
                item{
                    "Spaghetti",
                },
            },
        },
        menu{
            Rest : "Late Meal",
            Breakfast : []item{
                item{
                    "Toast",
                },
            },
            Lunch : []item{
                item{
                    "Hamburger",
                },
            },
            Dinner : []item{
                item{
                    "Grilled Chicken",
                },
            },
        },
        menu{
            Rest : "Trattoria Hein",
            Breakfast : []item{
                item{
                    "Orange Juice",
                },
                item{
                    "Milk",
                },
            },
            Lunch : []item{
                item{
                    "Turkey Sandwich",
                },
                item{
                    "Grilled Cheese",
                },
            },
            Dinner : []item{
                item{
                    "Filet Mignon",
                },
                item{
                    "Crab Cakes",
                },
            },
        },
    }


    err := tpl.Execute(os.Stdout, rest)
    if err != nil {
        log.Fatalln(err)
    }
}
