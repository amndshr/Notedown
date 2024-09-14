package main

import (
	"fmt"
	"time"
)

type Note struct {
	Title       string
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Font        string
	Pin         bool
	Attachments []string
}

type User struct {
	UserId        string
	Name          string
	Notes         []Note
	NumberOfNotes int
}

func main() {
	user1 := User{
		UserId: "amndshr",
		Name:   "Amanda",
		Notes: []Note{
			{
				Title:       "How to wake up early",
				Content:     "make sure to sleep enough the night before and set an alarm ",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Font:        "Ariel",
				Pin:         true,
				Attachments: []string{},
			},
		},
		NumberOfNotes: 1,
	}

	displayNumberOfNotes(user1)
	displayTitles(user1)
}

func displayNumberOfNotes(u User) {
	if u.NumberOfNotes > 1 {
		fmt.Println(u.NumberOfNotes, "Notes")
	} else {
		fmt.Println(u.NumberOfNotes, "Note")
	}
}

func displayTitles(u User) {
	for _, note := range u.Notes {
		fmt.Println(note.Title)
	}
}

// func compare (a user, b user){}
// func user input
// functionality that belongs to a note - addtext, addphoto, pin it, print list of title,
