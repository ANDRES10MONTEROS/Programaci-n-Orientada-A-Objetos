package models

import (
	"errors"
	"time"
)

type ContentType string

const (
	Video ContentType = "video"
	Audio ContentType = "audio"
)

type Content struct {
	id          string
	title       string
	ctype       ContentType
	durationSec int
	createdAt   time.Time
}

func NewContent(id, title string, ctype ContentType, durationSec int) (*Content, error) {
	if id == "" || title == "" {
		return nil, errors.New("id and title required")
	}
	if durationSec <= 0 {
		return nil, errors.New("duration must be positive")
	}
	return &Content{
		id:          id,
		title:       title,
		ctype:       ctype,
		durationSec: durationSec,
		createdAt:   time.Now(),
	}, nil
}

func (c *Content) GetID() string { return c.id }
func (c *Content) GetTitle() string { return c.title }
func (c *Content) GetType() ContentType { return c.ctype }
func (c *Content) GetDuration() int { return c.durationSec }
func (c *Content) GetCreatedAt() time.Time { return c.createdAt }

func (c *Content) SetTitle(title string) error {
	if title == "" { return errors.New("title cannot be empty") }
	c.title = title
	return nil
}
func (c *Content) SetDuration(sec int) error {
	if sec <= 0 { return errors.New("duration must be positive") }
	c.durationSec = sec
	return nil
}
