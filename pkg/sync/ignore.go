package sync

import (
	pathpkg "path"
	"strings"

	"github.com/pkg/errors"

	"github.com/bmatcuk/doublestar"
)

// UnmarshalText implements the text unmarshalling interface used when loading
// from TOML files.
func (m *IgnoreVCSMode) UnmarshalText(textBytes []byte) error {
	// Convert the bytes to a string.
	text := string(textBytes)

	// Convert to a VCS mode.
	switch text {
	case "true":
		*m = IgnoreVCSMode_IgnoreVCS
	case "false":
		*m = IgnoreVCSMode_PropagateVCS
	default:
		return errors.Errorf("unknown VCS ignore specification: %s", text)
	}

	// Success.
	return nil
}

// Supported indicates whether or not a particular VCS ignore mode is a valid,
// non-default value.
func (m IgnoreVCSMode) Supported() bool {
	switch m {
	case IgnoreVCSMode_IgnoreVCS:
		return true
	case IgnoreVCSMode_PropagateVCS:
		return true
	default:
		return false
	}
}

// Description returns a human-readable description of a VCS ignore mode.
func (m IgnoreVCSMode) Description() string {
	switch m {
	case IgnoreVCSMode_IgnoreVCSDefault:
		return "Default"
	case IgnoreVCSMode_IgnoreVCS:
		return "Ignore"
	case IgnoreVCSMode_PropagateVCS:
		return "Propagate"
	default:
		return "Unknown"
	}
}

// DefaultVCSIgnores is the default set of ignores to use when ignoring VCS
// directories.
var DefaultVCSIgnores = []string{
	".git/",
	".svn/",
	".hg/",
	".bzr/",
	"_darcs/",
}

type ignorePattern struct {
	negated       bool
	directoryOnly bool
	matchLeaf     bool
	pattern       string
}

func newIgnorePattern(pattern string) (*ignorePattern, error) {
	// Check for invalid patterns, or at least those that would leave us with an
	// empty string after parsing. Obviously we can't perform general complete
	// validation for all patterns, but if they pass this parsing, they should
	// be sane enough to at least try to match.
	if pattern == "" || pattern == "!" {
		return nil, errors.New("empty pattern")
	} else if pattern == "/" || pattern == "!/" {
		return nil, errors.New("root pattern")
	} else if pattern == "//" || pattern == "!//" {
		return nil, errors.New("root directory pattern")
	}

	// Check if this is a negated pattern. If so, remove the exclamation point
	// prefix, since it won't enter into pattern matching.
	negated := false
	if pattern[0] == '!' {
		negated = true
		pattern = pattern[1:]
	}

	// Check if this is an absolute pattern. If so, remove the forward slash
	// prefix, since it won't enter into pattern matching.
	absolute := false
	if pattern[0] == '/' {
		absolute = true
		pattern = pattern[1:]
	}

	// Check if this is a directory-only pattern. If so, remove the trailing
	// slash, since it won't enter into pattern matching.
	directoryOnly := false
	if pattern[len(pattern)-1] == '/' {
		directoryOnly = true
		pattern = pattern[:len(pattern)-1]
	}

	// Determine whether or not the pattern contains a slash.
	containsSlash := strings.IndexByte(pattern, '/') >= 0

	// Attempt to do a match with the pattern to ensure validity. We have to
	// match against a non-empty path (we choose something simple), otherwise
	// bad pattern errors won't be detected.
	if _, err := doublestar.Match(pattern, "a"); err != nil {
		return nil, errors.Wrap(err, "unable to validate pattern")
	}

	// Success.
	return &ignorePattern{
		negated:       negated,
		directoryOnly: directoryOnly,
		matchLeaf:     (!absolute && !containsSlash),
		pattern:       pattern,
	}, nil
}

func (i *ignorePattern) matches(path string, directory bool) (bool, bool) {
	// If this pattern only applies to directories and this is not a directory,
	// then this is not a match.
	if i.directoryOnly && !directory {
		return false, false
	}

	// Check if there is a direct match. Since we've already validated the
	// pattern in the constructor, we know match can't fail with an error (it's
	// only return code is on bad patterns).
	if match, _ := doublestar.Match(i.pattern, path); match {
		return true, i.negated
	}

	// If it makes sense, attempt to match on the last component of the path,
	// assuming the path is non-empty (non-root).
	if i.matchLeaf && path != "" {
		if match, _ := doublestar.Match(i.pattern, pathpkg.Base(path)); match {
			return true, i.negated
		}
	}

	// No match.
	return false, false
}

// ValidIgnorePattern checks whether or not a given pattern is a valid ignore
// specification.
func ValidIgnorePattern(pattern string) bool {
	// Verify that we can parse the ignore.
	_, err := newIgnorePattern(pattern)
	return err == nil
}

type ignorer struct {
	patterns []*ignorePattern
}

func newIgnorer(patterns []string) (*ignorer, error) {
	// Parse patterns.
	ignorePatterns := make([]*ignorePattern, len(patterns))
	for i, p := range patterns {
		if ip, err := newIgnorePattern(p); err != nil {
			return nil, errors.Wrap(err, "unable to parse pattern")
		} else {
			ignorePatterns[i] = ip
		}
	}

	// Success.
	return &ignorer{ignorePatterns}, nil
}

func (i *ignorer) ignored(path string, directory bool) bool {
	// Nothing is initially ignored.
	ignored := false

	// Run through patterns, keeping track of the ignored state as we reach more
	// specific rules.
	for _, p := range i.patterns {
		if match, negated := p.matches(path, directory); !match {
			continue
		} else {
			ignored = !negated
		}
	}

	// Done.
	return ignored
}
