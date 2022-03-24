package plumbing

import (
	"regexp"
	"time"
)

type ReflogEntry struct {
	Reference ReferenceName
	Index     int
	OldHash   Hash
	NewHash   Hash
	Author    Signature
	Message   ReflogEntryMessage
}

type Signature struct {
	// Name represents a person name. It is an arbitrary string.
	Name string
	// Email is an email, but it cannot be assumed to be well-formed.
	Email string
	// When is the timestamp of the signature.
	When time.Time
}

type ReflogEntryMessage string

func (m ReflogEntryMessage) Action() string {
	parts := m.parts()
	if parts == nil {
		return ""
	}
	return parts[0]
}

var reflogEntryMessageOptionsRegexp = regexp.MustCompile(`\w+`)

func (m ReflogEntryMessage) Options() []string {
	parts := m.parts()
	if parts == nil || parts[1] == "" {
		return nil
	}
	return reflogEntryMessageOptionsRegexp.FindAllString(parts[1], -1)
}

func (m ReflogEntryMessage) Body() string {
	parts := m.parts()
	if parts == nil {
		return ""
	}
	return parts[2]
}

var reflogEntryMessageRegexp = regexp.MustCompile(`^(\w+)([^:]*): (.*)$`)

func (m ReflogEntryMessage) parts() []string {
	matches := reflogEntryMessageRegexp.FindStringSubmatch(string(m))
	if len(matches) == 0 {
		return nil
	}
	return matches[1:]
}
