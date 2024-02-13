package celestia

import (
	"context"
)

type DataAvailabilityWriter interface {
	Store(context.Context, []byte) (*BlobPointer, bool, error)
	Verify(ctx context.Context, blobPointer *BlobPointer) (bool, error)
	Serialize(blobPointer *BlobPointer) ([]byte, error)
}

type DataAvailabilityReader interface {
	Read(context.Context, *BlobPointer) ([]byte, *SquareData, error)
}
