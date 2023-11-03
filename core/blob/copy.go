package blob

import (
	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/core"
)

// Copy copies data between accounts
type Copy struct {
	// The ID of the account to copy blobs from
	FromAccount jmap.ID `json:"fromAccountId,omitempty"`

	// The ID of the account to copy blobs to
	Account jmap.ID `json:"accountId,omitempty"`

	// A list of IDs of blobs to copy
	IDs []jmap.ID `json:"blobIds,omitempty"`
}

func (m *Copy) Name() string { return "Blob/copy" }

func (m *Copy) Requires() []jmap.URI { return []jmap.URI{core.URI} }

type CopyResponse struct {
	// The ID of the account blobs were copied from
	FromAccount jmap.ID `json:"fromAccountId,omitempty"`

	// The ID of the account blobs were copied to
	Account jmap.ID `json:"accountId,omitempty"`

	// A map of the blobId in the fromAccount to the ID of the blob in the
	// account it was copied to. Map is null if no blobs were copied
	Copied map[jmap.ID]jmap.ID `json:"blobIds,omitempty"`

	// A map of blobId to a SetError object for each blob that failed to be
	// copied, or null if none.
	NotCopied map[jmap.ID]*jmap.SetError `json:"notCopied,omitempty"`
}

func newCopyResponse() jmap.MethodResponse { return &CopyResponse{} }
