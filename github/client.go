package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

// Client http request client type
type Client struct {
	url        string
	httpClient *http.Client
}

// NewClient returns an new instance of client object.
func NewClient(timeout time.Duration) Client {
	return Client{httpClient: &http.Client{Timeout: timeout}}
}

func (c Client) getResponse(url string) (map[string]interface{}, error) {
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var content map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		return nil, err
	}

	for k, v := range content {
		fmt.Println(k, "::", v)
	}

	return content, nil
}

func (c Client) getMetricTypes(user string) ([]plugin.Metric, error) {
	url := "https://api.github.com/users/" + user
	m, err := c.getResponse(url)
	if err != nil {
		return nil, err
	}

	var metrics []plugin.Metric
	for k := range m {
		metric := plugin.Metric{
			Namespace: plugin.NewNamespace("intel", "github", "user").AddDynamicElement("userid", "github userid").AddStaticElement(k),
			Version:   1,
		}
		metrics = append(metrics, metric)
	}
	return metrics, nil
}

func (c Client) getMetricData(user string) (map[string]plugin.Metric, error) {
	url := "https://api.github.com/users/" + user
	rmp, err := c.getResponse(url)
	if err != nil {
		return nil, err
	}

	mp := map[string]plugin.Metric{}
	for k, v := range rmp {
		ns := plugin.NewNamespace("intel", "github", "user", user, k)
		metric := plugin.Metric{
			Namespace: ns,
			Version:   1,
			Data:      v,
			Timestamp: time.Now(),
		}
		mp[strings.Join(ns.Strings(), "/")] = metric
	}
	return mp, nil
}
