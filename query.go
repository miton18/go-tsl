package tsl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/miton18/go-warp10/base"
)

// Query is a TSL query
type Query struct {
	raw        string
	endpoint   string
	token      string
	httpClient *http.Client
}

// Execute the query on Backend
func (q *Query) Execute() (base.GTSList, error) {
	if q.httpClient == nil {
		q.httpClient = http.DefaultClient
	}

	r := strings.NewReader(q.raw)

	req, err := http.NewRequest("POST", q.endpoint+"/v0/query", r)
	if err != nil {
		return nil, fmt.Errorf("Cannot build TSL query: %s", err.Error())
	}

	req.SetBasicAuth("tsl-user", q.token)

	res, err := q.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Cannot perform TSL query: %s", err.Error())
	}

	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("Unexpected TSL query response: (%d) %s", res.StatusCode, string(b))
	}

	stack := []json.RawMessage{}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&stack)
	if err != nil {
		return nil, fmt.Errorf("Cannot decode TSL response: %s", err.Error())
	}

	if len(stack) == 0 {
		return nil, nil
	}

	var gts base.GTSList
	err = json.Unmarshal(stack[0], &gts)
	if err != nil {
		return gts, fmt.Errorf("Cannot parse TSL stack: %s", err.Error())
	}

	err = res.Body.Close()
	if err != nil {
		return gts, fmt.Errorf("Cannot close TSL response body: %s", err.Error())
	}

	return gts, nil
}

// Dump output the TSL query
func (q *Query) Dump() string {
	return q.raw
}

// Select a metric
func (q *Query) Select(metric string) *Query {
	if metric == "" {
		metric = "*"
	}

	q.raw += fmt.Sprintf("select(\"%s\")", metric)
	return q
}

// Where filter metric labels
// value must be set using Eq(), NotEq(), Like() or NotLike() functions
func (q *Query) Where(label, value string) *Query {
	q.raw += fmt.Sprintf(".where(\"%s%s\")", label, value)
	return q
}

// From set times limits
func (q *Query) From(start, end time.Time) *Query {
	if end.IsZero() {
		start = time.Now()
	}

	q.raw += fmt.Sprintf(".from(%d, %d)", toMicroSeconds(start), toMicroSeconds(end))
	return q
}

// Last set time duration
// TODO: handle shift parameter
func (q *Query) Last(d time.Duration, at time.Time) *Query {
	if at.IsZero() {
		q.raw += fmt.Sprintf(".last(%s)", shortDur(d))
	} else {
		q.raw += fmt.Sprintf(".last(%s, timestamp=%d)", shortDur(d), toMicroSeconds(at))
	}
	return q
}

// LastN for last N datapoints
func (q *Query) LastN(n int64, at time.Time) *Query {
	if at.IsZero() {
		q.raw += fmt.Sprintf(".last(%d)", n)
	} else {
		q.raw += fmt.Sprintf(".last(%d, timestamp=%d)", n, toMicroSeconds(at))
	}
	return q
}

// SampleBy sample the metrics in time buckets
func (q *Query) SampleBy(d time.Duration, aggregator Aggregator) *Query {
	q.raw += fmt.Sprintf(".sampleBy(%s, %s)", shortDur(d), aggregator)
	return q
}

// SampleByN sample the metrics in buckets count
func (q *Query) SampleByN(n int64, aggregator Aggregator) *Query {
	q.raw += fmt.Sprintf(".sampleBy(%d, %s)", n, aggregator)
	return q
}

// Group sample the metrics in buckets count
func (q *Query) Group(aggregator Aggregator) *Query {
	q.raw += fmt.Sprintf(".group(%s)", aggregator)
	return q
}

// GroupBy sample the metrics in buckets count
func (q *Query) GroupBy(labels []string, aggregator Aggregator) *Query {
	q.raw += fmt.Sprintf(".groupBy(%s, %s)", toStringArray(labels), aggregator)
	return q
}
