package clevermining

import (
	"github.com/dghubble/sling"
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"log"
	"strings"
)

type CleverMiningClient struct {
	sling      *sling.Sling
	httpClient *cleverminingHttpClient
}

// server send the api response with text/html content type
// we fix this: change content type to json
type cleverminingHttpClient struct {
	client    *http.Client
	debug     bool
	useragent string
}

func (d cleverminingHttpClient) Do(req *http.Request) (*http.Response, error) {
	if d.debug {
		d.dumpRequest(req)
	}
	if d.useragent != "" {
		req.Header.Set("User-Agent", d.useragent)
	}
	client := func() (*http.Client) {
		if d.client != nil {
			return d.client
		} else {
			return http.DefaultClient
		}
	}()
	if client.Transport != nil {
		if transport, ok := client.Transport.(*http.Transport); ok {
			if transport.TLSClientConfig != nil {
				transport.TLSClientConfig.InsecureSkipVerify = true;
			} else {
				transport.TLSClientConfig = &tls.Config{
					InsecureSkipVerify: true,
				}
			}
		}
	} else {
		if transport, ok := http.DefaultTransport.(*http.Transport); ok {
			if transport.TLSClientConfig != nil {
				transport.TLSClientConfig.InsecureSkipVerify = true;
			} else {
				transport.TLSClientConfig = &tls.Config{
					InsecureSkipVerify: true,
				}
			}
		}
	}
	resp, err := client.Do(req)
	if d.debug {
		d.dumpResponse(resp)
	}
	if err == nil {
		contenttype := resp.Header.Get("Content-Type");
		if len(contenttype) == 0 || strings.HasPrefix(contenttype, "text/html") {
			resp.Header.Set("Content-Type", "application/json")
		}
	}
	return resp, err
}

func (d cleverminingHttpClient) dumpRequest(r *http.Request) {
	if r == nil {
		log.Print("dumpReq ok: <nil>")
		return
	}
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Print("dumpReq err:", err)
	} else {
		log.Print("dumpReq ok:", string(dump))
	}
}

func (d cleverminingHttpClient) dumpResponse(r *http.Response) {
	if r == nil {
		log.Print("dumpResponse ok: <nil>")
		return
	}
	dump, err := httputil.DumpResponse(r, true)
	if err != nil {
		log.Print("dumpResponse err:", err)
	} else {
		log.Print("dumpResponse ok:", string(dump))
	}
}

func NewCleverMiningClient(client *http.Client, BaseURL string, UserAgent string) *CleverMiningClient {
	if len(BaseURL) == 0 {
		BaseURL = "https://www.clevermining.com/"
	}
	cleverminingclient := &cleverminingHttpClient{client:client, useragent:UserAgent}
	return &CleverMiningClient{
		httpClient: cleverminingclient,
		sling: sling.New().Doer(cleverminingclient).Base(strings.TrimRight(BaseURL, "/") + "/").Path("api/v1/"),
	}
}

func (client CleverMiningClient) SetDebug(debug bool) {
	client.httpClient.debug = debug
}
