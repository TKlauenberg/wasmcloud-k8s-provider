// Generated by `wit-bindgen-wrpc-go` 0.11.0. DO NOT EDIT!
package handler

import (
	bytes "bytes"
	context "context"
	errors "errors"
	fmt "fmt"
	k8s__provider__types "github.com/TKlauenberg/wasmcloud-k8s-provider/bindings/k8s/provider/types"
	slog "log/slog"
	wrpc "wrpc.io/go"
)

type ResourceEvent = k8s__provider__types.ResourceEvent

// Handle a Kubernetes resource event.
// @param event The resource event to handle.
// @returns A status indicating the result of the operation.
func Handle(ctx__ context.Context, wrpc__ wrpc.Invoker, event *k8s__provider__types.ResourceEvent) (err__ error) {
	var buf__ bytes.Buffer
	write0__, err__ := (event).WriteToIndex(&buf__)
	if err__ != nil {
		err__ = fmt.Errorf("failed to write `event` parameter: %w", err__)
		return
	}
	if write0__ != nil {
		err__ = errors.New("unexpected deferred write for synchronous `event` parameter")
		return
	}
	var w__ wrpc.IndexWriteCloser
	var r__ wrpc.IndexReadCloser
	w__, r__, err__ = wrpc__.Invoke(ctx__, "k8s:provider/handler@0.1.0", "handle", buf__.Bytes())
	if err__ != nil {
		err__ = fmt.Errorf("failed to invoke `handle`: %w", err__)
		return
	}
	defer func() {
		if err := r__.Close(); err != nil {
			slog.ErrorContext(ctx__, "failed to close reader", "instance", "k8s:provider/handler@0.1.0", "name", "handle", "err", err)
		}
	}()
	if cErr__ := w__.Close(); cErr__ != nil {
		slog.DebugContext(ctx__, "failed to close outgoing stream", "instance", "k8s:provider/handler@0.1.0", "name", "handle", "err", cErr__)
	}
	return
}
