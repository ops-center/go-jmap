package email

import "git.sr.ht/~rockorager/go-jmap"

// This is a standard "/changes" method as described in [RFC8620], Section 5.2.
// If generating intermediate states for a large set of changes, it is
// recommended that newer changes be returned first, as these are generally of
// more interest to users.
type Changes struct {
	// The id of the account to use.
	Account jmap.ID `json:"accountId"`

	// The current state of the client. This is the string that was
	// returned as the state argument in the Foo/get response. The server
	// will return the changes that have occurred since this state.
	SinceState string `json:"sinceState"`

	// The maximum number of ids to return in the response. The server MAY
	// choose to return fewer than this value but MUST NOT return more. If
	// not given by the client, the server may choose how many to return.
	// If supplied by the client, the value MUST be a positive integer
	// greater than 0. If a value outside of this range is given, the
	// server MUST reject the call with an invalidArguments error.
	MaxChanges uint64 `json:"maxChanges"`
}

func (m *Changes) Name() string { return "Mailbox/changes" }

func (m *Changes) Uses() string { return MailCapability }

func (m *Changes) NewResponse() interface{} { return &ChangesResponse{} }

// This is a standard "/changes" method as described in [RFC8620], Section 5.2.
// If generating intermediate states for a large set of changes, it is
// recommended that newer changes be returned first, as these are generally of
// more interest to users.
type ChangesResponse struct {
	// The id of the account used for the call.
	Account jmap.ID `json:"accountId"`

	// This is the sinceState argument echoed back; it’s the state from
	// which the server is returning changes.
	OldState string `json:"oldState"`

	// This is the state the client will be in after applying the set of
	// changes to the old state.
	NewState string `json:"newState"`

	// If true, the client may call Foo/changes again with the newState
	// returned to get further updates. If false, newState is the current
	// server state.
	HasMoreChanges bool `json:"hasMoreChanges"`

	// An array of ids for records that have been created since the old
	// state.
	Created []jmap.ID `json:"created"`

	// An array of ids for records that have been updated since the old
	// state.
	Updated []jmap.ID `json:"updated"`

	// An array of ids for records that have been destroyed since the old
	// state.
	Destroyed []jmap.ID `json:"destroyed"`
}
