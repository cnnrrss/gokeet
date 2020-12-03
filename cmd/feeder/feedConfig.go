package feeder

import (
	"regexp"
	"strings"
)

type FeedConfig struct {
	Host string `json:"host"`
	PathConstraints []string `json:"pathConstraints"`
	PathExclusions []string `json:"pathExclusions"`
}

// Constraints returns a comma-separated string containing the site's
// path constraints.
func (f FeedConfig) Constraints() string {
	return strings.Join(f.PathConstraints, ",")
}

// Exclusions returns a comma-separated string containing the parts of
// the site that must not be crawled
func (f FeedConfig) Exclusions() string {
	return strings.Join(f.PathExclusions, ",")
}


func pathMatchesRegexps(patterns []string, path string) bool {
	if len(patterns) == 0 {
		return true
	}

	for _, pattern := range patterns {
		if rx, err := regexp.Compile(pattern); err == nil {
			if !rx.MatchString(path) {
				return false
			}
		}
	}

	return true
}

// AllowPathByConstraints returns false if path doesn't match constraints list
func (f FeedConfig) AllowPathByConstraints(path string) bool {
	return pathMatchesRegexps(f.PathConstraints, path)
}


// AllowPathByExclusions returns false if path matches exclusion list
func (f FeedConfig) AllowPathByExclusions(path string) bool {
	return pathMatchesRegexps(f.PathExclusions, path)
}