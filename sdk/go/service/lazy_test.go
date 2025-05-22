package service_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/1backend/1backend/sdk/go/service"
)

type noOpStarter struct {
}

func (s *noOpStarter) LazyStart() error {
	return nil
}

func TestLazy_DoesNotBlockConcurrentHandlers(t *testing.T) {
	starter := &noOpStarter{}

	slowHandler := service.Lazy(starter, func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.Write([]byte("slow done"))
	})

	fastHandler := service.Lazy(starter, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("fast done"))
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/svc/slow", slowHandler)
	mux.HandleFunc("/svc/fast", fastHandler)

	server := httptest.NewServer(mux)
	defer server.Close()

	go func() {
		_, err := http.Get(server.URL + "/svc/slow")
		if err != nil {
			t.Errorf("slow handler error: %v", err)
		}
	}()

	time.Sleep(5 * time.Millisecond)

	start := time.Now()
	resp, err := http.Get(server.URL + "/svc/fast")
	if err != nil {
		t.Errorf("fast handler error: %v", err)
		return
	}
	defer resp.Body.Close()
	fastDuration := time.Since(start)

	if fastDuration > 5*time.Millisecond {
		t.Errorf("fast handler took too long (was blocked?): %s", fastDuration)
	} else {
		t.Logf("fast handler returned quickly: %s", fastDuration)
	}
}
