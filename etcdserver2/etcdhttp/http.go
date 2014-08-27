package etcdhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"code.google.com/p/go.net/context"
	etcdserver "github.com/coreos/etcd/etcdserver2"
	"github.com/coreos/etcd/store"
)

const DefaultTimeout = 500 * time.Millisecond

type Handler struct {
	Timeout time.Duration
	Server  etcdserver.Server
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: set read/write timeout?

	timeout := h.Timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	ctx, _ := context.WithTimeout(context.Background(), timeout)
	// TODO(bmizerany): watch the closenotify chan in another goroutine can
	// call cancel when it closes. be sure to watch ctx.Done() too so we
	// don't leak a ton of these goroutines.

	rr, err := parseRequest(r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	resp, err := h.Server.Do(ctx, rr)
	if err != nil {
		// TODO(bmizerany): switch on store errors and etcdserver.ErrUnknownMethod
		panic("TODO")
	}

	encodeResponse(ctx, w, resp)
}

func parseRequest(r *http.Request) (etcdserver.Request, error) {
	return etcdserver.Request{}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, resp etcdserver.Response) error {
	var ev *store.Event
	switch {
	case resp.Event != nil:
		ev = resp.Event
	case resp.Watcher != nil:
		// TODO(bmizerany): support streaming?
		defer resp.Watcher.Remove()
		select {
		case ev = <-resp.Watcher.EventChan:
		case <-ctx.Done():
			return ctx.Err()
		}
	default:
		panic("should not be rechable")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("X-Etcd-Index", fmt.Sprint(ev.Index()))

	if ev.IsCreated() {
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	return json.NewEncoder(w).Encode(ev)
}