package discord

import (
	"regexp"
	"strings"
)

var domainURLPattern = regexp.MustCompile(`(?i)\b((?:https?://)?(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z]{2,24})(?::\d{1,5})?(?:/[^\s]*)?`)
var inviteURLPattern = regexp.MustCompile(`(?i)(?:discord\.gg|discord(?:app)?\.com/invite|discord\.com/invite|discord\.gg)/\S+`)

// FindURLs returns URL or bare-domain tokens from text.
func FindURLs(text string) []string { return domainURLPattern.FindAllString(text, -1) }

// HasDiscordInvite reports whether text contains a Discord invite URL.
func HasDiscordInvite(text string) bool { return FindDiscordInvite(text) != "" }

// FindDiscordInvite returns the first Discord invite URL in text.
func FindDiscordInvite(text string) string { return inviteURLPattern.FindString(text) }

// ExtractHost returns the lowercase host from a URL or domain token.
func ExtractHost(raw string) string {
	s := strings.ToLower(strings.TrimSpace(raw))
	if i := strings.Index(s, "://"); i != -1 {
		s = s[i+3:]
	}
	if i := strings.IndexAny(s, "/?#"); i != -1 {
		s = s[:i]
	}
	if i := strings.Index(s, "@"); i != -1 {
		s = s[i+1:]
	}
	if i := strings.Index(s, ":"); i != -1 {
		s = s[:i]
	}
	s = strings.TrimPrefix(s, "www.")
	return strings.Trim(s, ".")
}
