package types

import "fmt"

// Feed is an single twtxt.txt feed with a cannonical Nickname and URL for the feed
type Feed struct {
	Nick string
	URL  string
}

// String implements the Stringer interface and returns the Feed represented
// as a twtxt.txt URI in the form @<nick url>
func (f Feed) String() string {
	return fmt.Sprintf("@<%s %s>", f.Nick, f.URL)
}

// Feeds is a mappping of Feed to booleans used to ensure unique feeds
type Feeds map[Feed]bool
