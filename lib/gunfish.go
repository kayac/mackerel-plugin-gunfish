package gunfish

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

// Plugin mackerel plugin for gunfish
type Plugin struct {
	Target   string
	Tempfile string
	Prefix   string
}

// MetricKeyPrefix interface for PluginWithPrefix
func (m Plugin) MetricKeyPrefix() string {
	if m.Prefix == "" {
		m.Prefix = "gunfish"
	}
	return m.Prefix
}

// FetchMetrics interface for mackerelplugin
func (m Plugin) FetchMetrics() (map[string]float64, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/stats/app", m.Target))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	s := struct {
		Uptime                 float64 `json:"uptime"`
		StartAt                float64 `json:"start_at"`
		ServiceUnavailableAt   float64 `json:"su_at"`
		Workers                float64 `json:"workers"`
		QueueSize              float64 `json:"queue_size"`
		RetryQueueSize         float64 `json:"retry_queue_size"`
		WorkersQueueSize       float64 `json:"workers_queue_size"`
		CommandQueueSize       float64 `json:"cmdq_queue_size"`
		RetryCount             float64 `json:"retry_count"`
		RequestCount           float64 `json:"req_count"`
		SentCount              float64 `json:"sent_count"`
		ErrCount               float64 `json:"err_count"`
		CertificateExpireUntil float64 `json:"certificate_expire_until"`
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return nil, err
	}
	ret := make(map[string]float64, 13)
	ret["uptime"] = s.Uptime
	ret["workers"] = s.Workers
	ret["queue_size"] = s.QueueSize
	ret["retry_queue_size"] = s.RetryQueueSize
	ret["workers_queue_size"] = s.WorkersQueueSize
	ret["cmdq_queue_size"] = s.CommandQueueSize
	ret["retry_count"] = s.RetryCount
	ret["req_count"] = s.RequestCount
	ret["sent_count"] = s.SentCount
	ret["err_count"] = s.ErrCount
	ret["certificate_expire_until"] = s.CertificateExpireUntil

	return ret, nil
}

// GraphDefinition interface for mackerelplugin
func (m Plugin) GraphDefinition() map[string]mp.Graphs {
	labelPrefix := strings.Title(m.Prefix)

	var graphdef = map[string]mp.Graphs{
		"messages": {
			Label: (labelPrefix + " Messages"),
			Unit:  mp.UnitFloat,
			Metrics: []mp.Metrics{
				{Name: "req_count", Label: "Requests", Diff: true},
				{Name: "sent_count", Label: "Sents", Diff: true},
				{Name: "retry_count", Label: "Retries", Diff: true},
				{Name: "err_count", Label: "Errors", Diff: true},
			},
		},
		"queue": {
			Label: (labelPrefix + "Server"),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "queue_size", Label: "Incoming Queue size", Diff: false},
				{Name: "retry_queue_size", Label: "Retrying Queue size", Diff: false},
				{Name: "cmdq_queue_size", Label: "Command Queue size", Diff: false},
				{Name: "workers_queue_size", Label: "Workers queue size", Diff: false},
			},
		},
		"workers": {
			Label: (labelPrefix + " Workers"),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "workers", Label: "Workers", Diff: false},
			},
		},
		"certificates": {
			Label: (labelPrefix + " Certificates"),
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "certificate_expire_until", Label: "Expires until", Diff: false},
			},
		},
	}
	return graphdef
}

// Do the plugin
func Do() {
	optHost := flag.String("host", "localhost", "Hostname")
	optPort := flag.String("port", "8003", "Port")
	optPrefix := flag.String("metric-key-prefix", "gunfish", "Metric key prefix")
	optTempfile := flag.String("tempfile", "", "Temp file name")
	flag.Parse()

	var gunfish Plugin
	gunfish.Prefix = *optPrefix
	gunfish.Target = fmt.Sprintf("%s:%s", *optHost, *optPort)

	helper := mp.NewMackerelPlugin(gunfish)
	helper.Tempfile = *optTempfile
	helper.Run()
}
