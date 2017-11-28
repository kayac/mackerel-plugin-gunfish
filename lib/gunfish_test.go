package gunfish

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphDefinition(t *testing.T) {
	var gunfish Plugin

	graphdef := gunfish.GraphDefinition()
	if len(graphdef) != 4 {
		t.Errorf("GetTempfilename: %d should be 4", len(graphdef))
	}
}

var statsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{
  "pid": 32611,
  "debug_port": 0,
  "uptime": 6477451,
  "start_at": 1505379870,
  "su_at": 0,
  "period": 13,
  "retry_after": 10,
  "workers": 8,
  "queue_size": 0,
  "retry_queue_size": 0,
  "workers_queue_size": 0,
  "cmdq_queue_size": 0,
  "retry_count": 0,
  "req_count": 9578,
  "sent_count": 615434,
  "err_count": 4297,
  "certificate_not_after": "2018-09-07T02:58:21Z",
  "certificate_expire_until": 24431779
}`)
})

func TestParse(t *testing.T) {
	var gunfish Plugin
	mux := http.NewServeMux()
	mux.HandleFunc("/stats/app", statsHandler)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	u, _ := url.Parse(ts.URL)
	gunfish.Target = u.Host

	stat, err := gunfish.FetchMetrics()
	fmt.Println(stat)
	assert.Nil(t, err)
	// Gunfish Stats
	assert.EqualValues(t, stat["uptime"], 6477451)
	assert.EqualValues(t, stat["workers"], 8)
	assert.EqualValues(t, stat["queue_size"], 0)
	assert.EqualValues(t, stat["retry_queue_size"], 0)
	assert.EqualValues(t, stat["workers_queue_size"], 0)
	assert.EqualValues(t, stat["cmdq_queue_size"], 0)
	assert.EqualValues(t, stat["retry_count"], 0)
	assert.EqualValues(t, stat["req_count"], 9578)
	assert.EqualValues(t, stat["sent_count"], 615434)
	assert.EqualValues(t, stat["err_count"], 4297)
	assert.EqualValues(t, stat["certificate_expire_until"], 24431779)
}
