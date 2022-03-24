package plumbing

import (
	. "gopkg.in/check.v1"
)

type ReflogSuite struct{}

var _ = Suite(&ReflogSuite{})

func (s *ReflogSuite) TestReflogEntryMessageAction(c *C) {
	var tests = []struct {
		message  ReflogEntryMessage
		expected string
	}{
		{"pull: Fast-forward", "pull"},
		{"checkout: moving from dev to tech-fix-transacting-domain-transaction", "checkout"},
		{"rebase (start): checkout dev", "rebase"},
		{"rebase (continue) (pick): tech(api): add a getMany in learning content DS and use it in challenge repository", "rebase"},
		{"rebase (continue) (finish): returning to refs/heads/pix-4039-flash-save-theta-errorrate", "rebase"},
		{"reset: moving to origin/dev", "reset"},
		{"commit: sr(scripts): fix eslint-plugin-eslint-comments errors", "commit"},
		{"commit (amend): feat(mon-pix): saved tutorials empty", "commit"},
	}

	for _, t := range tests {
		c.Assert(t.message.Action(), Equals, t.expected)
	}
}

func (s *ReflogSuite) TestReflogEntryMessageOptions(c *C) {
	var tests = []struct {
		message  ReflogEntryMessage
		expected []string
	}{
		{"pull: Fast-forward", nil},
		{"checkout: moving from dev to tech-fix-transacting-domain-transaction", nil},
		{"rebase (start): checkout dev", []string{"start"}},
		{"rebase (continue) (pick): tech(api): add a getMany in learning content DS and use it in challenge repository", []string{"continue", "pick"}},
		{"rebase (continue) (finish): returning to refs/heads/pix-4039-flash-save-theta-errorrate", []string{"continue", "finish"}},
		{"reset: moving to origin/dev", nil},
		{"commit: sr(scripts): fix eslint-plugin-eslint-comments errors", nil},
		{"commit (amend): feat(mon-pix): saved tutorials empty", []string{"amend"}},
	}

	for _, t := range tests {
		c.Assert(t.message.Options(), DeepEquals, t.expected)
	}
}

func (s *ReflogSuite) TestReflogEntryMessageBody(c *C) {
	var tests = []struct {
		message  ReflogEntryMessage
		expected string
	}{
		{"pull: Fast-forward", "Fast-forward"},
		{"checkout: moving from dev to tech-fix-transacting-domain-transaction", "moving from dev to tech-fix-transacting-domain-transaction"},
		{"rebase (start): checkout dev", "checkout dev"},
		{"rebase (continue) (pick): tech(api): add a getMany in learning content DS and use it in challenge repository", "tech(api): add a getMany in learning content DS and use it in challenge repository"},
		{"rebase (continue) (finish): returning to refs/heads/pix-4039-flash-save-theta-errorrate", "returning to refs/heads/pix-4039-flash-save-theta-errorrate"},
		{"reset: moving to origin/dev", "moving to origin/dev"},
		{"commit: sr(scripts): fix eslint-plugin-eslint-comments errors", "sr(scripts): fix eslint-plugin-eslint-comments errors"},
		{"commit (amend): feat(mon-pix): saved tutorials empty", "feat(mon-pix): saved tutorials empty"},
	}

	for _, t := range tests {
		c.Assert(t.message.Body(), Equals, t.expected)
	}
}
