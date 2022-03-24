package filesystem

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
	"github.com/go-git/go-git/v5/storage/filesystem/dotgit"
)

type ReflogStorage struct {
	dir *dotgit.DotGit
}

var _ storer.ReflogStorer = &ReflogStorage{}

func (s *ReflogStorage) IterReflog(ref plumbing.ReferenceName) (storer.ReflogIter, error) {
	reflog, err := s.dir.Reflog(ref)
	if err != nil {
		return nil, err
	}

	return storer.NewReflogEntrySliceIter(reflog), nil
}
