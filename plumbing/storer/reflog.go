package storer

import (
	"io"

	"github.com/go-git/go-git/v5/plumbing"
)

type ReflogStorer interface {
	IterReflog(plumbing.ReferenceName) (ReflogIter, error)
}

type ReflogIter interface {
	Next() (*plumbing.ReflogEntry, error)
	ForEach(func(*plumbing.ReflogEntry) error) error
	Close()
}

type ReflogEntrySliceIter struct {
	entries []*plumbing.ReflogEntry
	pos     int
}

func NewReflogEntrySliceIter(entries []*plumbing.ReflogEntry) ReflogIter {
	return &ReflogEntrySliceIter{
		entries: entries,
	}
}

func (iter *ReflogEntrySliceIter) Next() (*plumbing.ReflogEntry, error) {
	if iter.pos >= len(iter.entries) {
		return nil, io.EOF
	}

	obj := iter.entries[iter.pos]
	iter.pos++
	return obj, nil
}

func (iter *ReflogEntrySliceIter) ForEach(cb func(*plumbing.ReflogEntry) error) error {
	defer iter.Close()
	for {
		obj, err := iter.Next()
		if err != nil {
			if err == io.EOF {
				return nil
			}

			return err
		}

		if err := cb(obj); err != nil {
			if err == ErrStop {
				return nil
			}

			return err
		}
	}
}

func (iter *ReflogEntrySliceIter) Close() {
	iter.pos = len(iter.entries)
}
