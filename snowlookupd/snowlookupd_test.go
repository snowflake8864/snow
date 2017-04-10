package snowlookupd

import (
	"testing"
)

func TestNoLogger(t *testing.T) {
	opts := NewOptions()
	opts.Logger = nil
	opts.TCPAddress = "127.0.0.1:0"
	nsqlookupd := New(opts)
	nsqlookupd.logf("should never be logged")

}
