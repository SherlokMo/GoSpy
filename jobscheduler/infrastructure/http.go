package infrastructure

import (
	"crypto/tls"
	"net/http"
	"net/http/httptrace"
	"time"
)

type ResponseTracer struct {
	DNS          time.Duration
	Connection   time.Duration
	TLSHandshake time.Duration
	TotalTime    time.Duration
}

func TraceHttpConnection(method string, url string) (*ResponseTracer, error) {
	req, _ := http.NewRequest(method, url, nil)

	var start, connect, dns, tlsHandshake time.Time
	var dnsDone, connectTime, tlsHandShakeTime, totalTime time.Duration
	trace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) { dns = time.Now() },
		DNSDone: func(ddi httptrace.DNSDoneInfo) {
			dnsDone = time.Since(dns)
		},

		TLSHandshakeStart: func() { tlsHandshake = time.Now() },
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			tlsHandShakeTime = time.Since(tlsHandshake)
		},

		ConnectStart: func(network, addr string) { connect = time.Now() },
		ConnectDone: func(network, addr string, err error) {
			connectTime = time.Since(connect)
		},

		GotFirstResponseByte: func() {
			totalTime = time.Since(start)
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	start = time.Now()
	_, err := http.DefaultTransport.RoundTrip(req)
	return &ResponseTracer{
		dnsDone,
		connectTime,
		tlsHandShakeTime,
		totalTime,
	}, err
}
