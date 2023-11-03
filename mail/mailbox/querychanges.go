package mailbox

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
)

// Get changes on a mailbox query
// https://www.rfc-editor.org/rfc/rfc8621.html#section-2.4
type QueryChanges struct {
	Account jmap.ID `json:"accountId,omitempty"`

	Filter Filter `json:"filter,omitempty"`

	Sort []*SortComparator `json:"sort,omitempty"`

	SinceQueryState string `json:"sinceQueryState,omitempty"`

	MaxChanges uint64 `json:"maxChanges,omitempty"`

	UpToID jmap.ID `json:"upToId,omitempty"`

	CalculateTotal bool `json:"calculateTotal,omitempty"`
}

func (m *QueryChanges) Name() string { return "Mailbox/queryChanges" }

func (m *QueryChanges) Requires() []jmap.URI { return []jmap.URI{mail.URI} }

type QueryChangesResponse struct {
	Account jmap.ID `json:"accountId,omitempty"`

	OldQueryState string `json:"oldQueryState,omitempty"`

	NewQueryState string `json:"newQueryState,omitempty"`

	Removed []jmap.ID `json:"removed,omitempty"`

	Added []*jmap.AddedItem `json:"added,omitempty"`
}

func newQueryChangesResponse() jmap.MethodResponse { return &QueryChangesResponse{} }
