package core

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

type History struct {
	Ready bool
	db *sql.DB
}

func NewHistory() *History {
	return &History{false, nil}
}

func (self *History) Prepare() error {
	time.Sleep(3 * time.Second)
	if self.Ready {
		if self.db == nil {
			return errors.New("Invalid state: history is ready, but it's actually not.")
		}
		return nil
	}
	db, err := sql.Open("sqlite3", "history.db")
	if err != nil {
		return err
	}
	m := migrations{db}
	m.run()
	self.Ready = true
	return nil
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
	self.assertIfReady()
	return []HistoryEntry{}
}

func (self *History) Update(newEntries []HistoryEntry) {
	self.assertIfReady()
}

func (self *History) assertIfReady() {
	if self.Ready {
		return
	}
	fmt.Printf("History isn't ready! Please, report the bug.")
	os.Exit(1)
}

type migrations struct {
	db *sql.DB
}

func (self *migrations) run() {
}

