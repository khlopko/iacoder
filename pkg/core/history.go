package core

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type History struct {
	db *sql.DB
}

func NewHistory() (*History, error) {
	db, err := sql.Open("sqlite3", "history.db")
	if err != nil {
		return nil, err
	}
	return &History{db}, nil
}

type HistoryEntryAuthor int

const (
	User HistoryEntryAuthor = iota
	Assistant HistoryEntryAuthor = iota
)

type HistoryEntry struct {
	Id int
	Author HistoryEntryAuthor
	Message string
}

func (self *History) LoadWithLimit(limit int) []HistoryEntry {
	return []HistoryEntry{}
}

func (self *History) Update(newEntries []HistoryEntry) {
}

