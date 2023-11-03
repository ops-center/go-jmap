package email

import (
	"encoding/json"
	"time"

	"git.sr.ht/~rockorager/go-jmap"
)

type Filter interface {
	implementsFilter()
}

type FilterOperator struct {
	Operator   jmap.Operator `json:"operator,omitempty"`
	Conditions []Filter      `json:"conditions,omitempty"`
}

func (fo *FilterOperator) implementsFilter() {}

// EmailFilterCondition is an interface that represents FilterCondition
// objects. A filter condition object can be either a named struct, ie
// EmailFilterConditionName, or an EmailFilter itself. EmailFilters can
// be used to create complex filtering
type FilterCondition struct {
	// A Mailbox id.  An Email must be in this Mailbox to match the condition.
	InMailbox jmap.ID `json:"inMailbox,omitempty"`

	// A list of Mailbox ids.  An Email must be in at least one Mailbox not in this
	// list to match the condition.  This is to allow messages solely in trash/spam
	// to be easily excluded from a search.
	InMailboxOtherThan []jmap.ID `json:"inMailboxOtherThan,omitempty"`

	// The "receivedAt" date-time of the Email must be before this date- time to
	// match the condition.
	Before *time.Time `json:"before,omitempty"`

	// The "receivedAt" date-time of the Email must be the same or after this
	// date-time to match the condition.
	After *time.Time `json:"after,omitempty"`

	// The "size" property of the Email must be equal to or greater than this
	// number to match the condition.
	MinSize uint64 `json:"minSize,omitempty"`

	// The "size" property of the Email must be less than this number to match the
	// condition.
	MaxSize uint64 `json:"maxSize,omitempty"`

	// All Emails (including this one) in the same Thread as this Email must have
	// the given keyword to match the condition.
	AllInThreadHaveKeyword string `json:"allInThreadHaveKeyword,omitempty"`

	// At least one Email (possibly this one) in the same Thread as this Email must
	// have the given keyword to match the condition.
	SomeInThreadHaveKeyword string `json:"someInThreadHaveKeyword,omitempty"`

	// All Emails (including this one) in the same Thread as this Email must *not*
	// have the given keyword to match the condition.
	NoneInThreadHaveKeyword string `json:"noneInThreadHaveKeyword,omitempty"`

	// This Email must have the given keyword to match the condition.
	HasKeyword string `json:"hasKeyword,omitempty"`

	// This Email must not have the given keyword to match the condition.
	NotKeyword string `json:"notKeyword,omitempty"`

	// The "hasAttachment" property of the Email must be identical to the value
	// given to match the condition.
	HasAttachment bool `json:"hasAttachment,omitempty"`

	// Looks for the text in Emails.  The server MUST look up text in the From, To,
	// Cc, Bcc, and Subject header fields of the message and SHOULD look inside any
	// "text/*" or other body parts that may be converted to text by the server.
	// The server MAY extend the search to any additional textual property.
	Text string `json:"text,omitempty"`

	// Looks for the text in the From header field of the message.
	From string `json:"from,omitempty"`

	// Looks for the text in the To header field of the message.
	To string `json:"to,omitempty"`

	// Looks for the text in the Cc header field of the message.
	Cc string `json:"cc,omitempty"`

	// Looks for the text in the Bcc header field of the message.
	Bcc string `json:"bcc,omitempty"`

	// Looks for the text in the Subject header field of the message.
	Subject string `json:"subject,omitempty"`

	// Looks for the text in one of the body parts of the message.  The server MAY
	// exclude MIME body parts with content media types other than "text/*" and
	// "message/*" from consideration in search matching.  Care should be taken to
	// match based on the text content actually presented to an end user by viewers
	// for that media type or otherwise identified as appropriate for search
	// indexing. Matching document metadata uninteresting to an end user (e.g.,
	// markup tag and attribute names) is undesirable.
	Body string `json:"body,omitempty"`

	// The array MUST contain either one or two elements.  The first element is the
	// name of the header field to match against.  The second (optional) element is
	// the text to look for in the header field value.  If not supplied, the
	// message matches simply if it has a header field of the given name.
	Header []string `json:"header,omitempty"`

	// When true, only messages where smimeStatus is not null match
	//
	// Requires server to support urn:ietf:jmap:smimeverify
	HasSMIME bool `json:"hasSmime,omitempty"`

	// When true, only messages with successfully verified SMIME match
	//
	// Requires server to support urn:ietf:jmap:smimeverify
	HasVerifiedSMIME bool `json:"hasVerifiedSmime,omitempty"`

	// When true, only messages with successfully verified SMIME at the time
	// of delivery match
	//
	// Requires server to support urn:ietf:jmap:smimeverify
	HasVerifiedSMIMEAtDelivery bool `json:"hasVerifiedSmimeAtDelivery,omitempty"`
}

func (fc *FilterCondition) implementsFilter() {}

func (fc *FilterCondition) MarshalJSON() ([]byte, error) {
	if fc.Before != nil && fc.Before.Location() != time.UTC {
		utc := fc.Before.UTC()
		fc.Before = &utc
	}
	if fc.After != nil && fc.After.Location() != time.UTC {
		utc := fc.After.UTC()
		fc.After = &utc
	}
	// create a type alias to avoid infinite recursion
	type Alias FilterCondition
	return json.Marshal((*Alias)(fc))
}
