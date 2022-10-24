package channels

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1alpha1"
	"kubegems.io/kubegems/pkg/apis/gems"
)

var (
	alertProxyReceiverHost = fmt.Sprintf("alertproxy.%s:9094", gems.NamespaceMonitor)
	alertProxyFeishu       = "feishu"
)

type Feishu struct {
	ChannelType `json:"channelType"`
	URL         string `json:"url"`        // feishu robot webhook url
	At          string `json:"at"`         // 要@的用户id，所有人则是 all
	SignSecret  string `json:"signSecret"` // 签名校验key
}

func (f *Feishu) ToReceiver(name string) v1alpha1.Receiver {
	q := url.Values{}
	q.Add("type", string(TypeFeishu))
	q.Add("url", f.URL)
	q.Add("at", f.At)
	q.Add("signSecret", f.SignSecret)
	u := fmt.Sprintf("http://%s?%s", alertProxyReceiverHost, q.Encode())
	return v1alpha1.Receiver{
		Name: name,
		WebhookConfigs: []v1alpha1.WebhookConfig{
			{
				URL: &u,
			},
		},
	}
}

func (f *Feishu) Check() error {
	if !strings.Contains(f.URL, "open.feishu.cn") {
		return fmt.Errorf("feishu robot url not valid")
	}
	return nil
}

func (f *Feishu) Test() error {
	return nil
}
