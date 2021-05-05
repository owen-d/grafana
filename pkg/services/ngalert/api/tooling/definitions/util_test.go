package definitions

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapJSONKeys(t *testing.T) {
	input := `{
  "template_files": {},
  "alertmanager_config": {
    "global": {
      "resolve_timeout": "4m",
      "http_config": {
        "BasicAuth": null,
        "Authorization": null,
        "BearerToken": "",
        "BearerTokenFile": "",
        "ProxyURL": {},
        "TLSConfig": {
          "CAFile": "",
          "CertFile": "",
          "KeyFile": "",
          "ServerName": "",
          "InsecureSkipVerify": false
        },
        "FollowRedirects": true
      },
      "smtp_from": "youraddress@example.org",
      "smtp_hello": "localhost",
      "smtp_smarthost": "localhost:25",
      "smtp_require_tls": true,
      "pagerduty_url": "https://events.pagerduty.com/v2/enqueue",
      "opsgenie_api_url": "https://api.opsgenie.com/",
      "wechat_api_url": "https://qyapi.weixin.qq.com/cgi-bin/",
      "victorops_api_url": "https://alert.victorops.com/integrations/generic/20131114/alert/"
    },
    "route": {
      "receiver": "example-email"
    },
    "templates": [],
    "receivers": [
      {
        "name": "example-email",
        "email_configs": [
          {
            "auth_password": "shh",
            "send_resolved": false,
            "to": "youraddress@example.org",
            "smarthost": "",
            "html": "{{ template \"email.default.html\" . }}",
            "tls_config": {
              "CAFile": "",
              "CertFile": "",
              "KeyFile": "",
              "ServerName": "",
              "InsecureSkipVerify": false
            }
          }
        ]
      }
    ]
  }
}`

	expected := `{
  "alertmanager_config": {
    "global": {
      "http_config": {
        "authorization": null,
        "basic_auth": null,
        "bearer_token": "",
        "bearer_token_file": "",
        "follow_redirects": true,
        "proxy_url": {},
        "tls_config": {
          "ca_file": "",
          "cert_file": "",
          "insecure_skip_verify": false,
          "key_file": "",
          "server_name": ""
        }
      },
      "opsgenie_api_url": "https://api.opsgenie.com/",
      "pagerduty_url": "https://events.pagerduty.com/v2/enqueue",
      "resolve_timeout": "4m",
      "smtp_from": "youraddress@example.org",
      "smtp_hello": "localhost",
      "smtp_require_tls": true,
      "smtp_smarthost": "localhost:25",
      "victorops_api_url": "https://alert.victorops.com/integrations/generic/20131114/alert/",
      "wechat_api_url": "https://qyapi.weixin.qq.com/cgi-bin/"
    },
    "receivers": [
      {
        "email_configs": [
          {
            "auth_password": "shh",
            "html": "{{ template \"email.default.html\" . }}",
            "send_resolved": false,
            "smarthost": "",
            "tls_config": {
              "ca_file": "",
              "cert_file": "",
              "insecure_skip_verify": false,
              "key_file": "",
              "server_name": ""
            },
            "to": "youraddress@example.org"
          }
        ],
        "name": "example-email"
      }
    ],
    "route": {
      "receiver": "example-email"
    },
    "templates": []
  },
  "template_files": {}
}`

	var tmp PostableUserConfig
	// ensure it meets spec
	require.Nil(t, json.Unmarshal([]byte(input), &tmp))

	var ifc interface{}
	require.Nil(t, json.Unmarshal([]byte(input), &ifc))
	out, err := mapJSONKeys(ifc, ToSnakeCase)
	require.Nil(t, err)

	b, err := json.MarshalIndent(&out, "", "  ")
	require.Nil(t, err)
	require.Equal(t, expected, string(b))

}
