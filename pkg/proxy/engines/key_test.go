/*
 * Copyright 2018 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package engines

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	oo "github.com/tricksterproxy/trickster/pkg/backends/options"
	"github.com/tricksterproxy/trickster/pkg/cache/key"
	tl "github.com/tricksterproxy/trickster/pkg/logging"
	ct "github.com/tricksterproxy/trickster/pkg/proxy/context"
	"github.com/tricksterproxy/trickster/pkg/proxy/headers"
	po "github.com/tricksterproxy/trickster/pkg/proxy/paths/options"
	"github.com/tricksterproxy/trickster/pkg/proxy/request"
	tu "github.com/tricksterproxy/trickster/pkg/util/testing"
)

const testMultipartBoundary = `; boundary=------------------------d0509edbe55938c0`
const testMultipartBody = `--------------------------d0509edbe55938c0
Content-Disposition: form-data; name="field1"

value1
--------------------------d0509edbe55938c0
Content-Disposition: form-data; name="field2"

value2
--------------------------d0509edbe55938c0--
`

const testJSONDocument = `
{
	"requestType": "query",
	"query": {
		"table": "movies",
		"fields": "eidr,title",
		"filter": "year=1979",
		"options": {
			"batchSize": 20,
			"someArray": [ "test" ],
			"booleanHere": true
		}
	},
	"field1": "value1"
}
`

func TestDeepSearch(t *testing.T) {

	var document map[string]interface{}
	err := json.Unmarshal([]byte(testJSONDocument), &document)
	if err != nil {
		t.Error(err)
	}

	val, err := deepSearch(document, "query/table")
	if err != nil {
		t.Error(err)
	}

	if val != "movies" {
		t.Errorf("expected %s got %s", "movies", val)
	}

	_, err = deepSearch(document, "")
	if err == nil {
		t.Errorf("expected error: %s", "could not find key")
	}

	_, err = deepSearch(document, "missingKey")
	if err == nil {
		t.Errorf("expected error: %s", "could not find key")
	}

	_, err = deepSearch(document, "query/filter/nottamap")
	if err == nil {
		t.Errorf("expected error: %s", "could not find key")
	}

	_, err = deepSearch(document, "query/options/batchSize")
	if err != nil {
		t.Error(err)
	}

	_, err = deepSearch(document, "query/options/booleanHere")
	if err != nil {
		t.Error(err)
	}

	_, err = deepSearch(document, "query/options/someArray")
	if err == nil {
		t.Errorf("expected error: %s", "could not find key")
	}
}

func TestDeriveCacheKey(t *testing.T) {

	rpath := &po.Options{
		Path:               "/",
		CacheKeyParams:     []string{"query", "step", "time"},
		CacheKeyHeaders:    []string{},
		CacheKeyFormFields: []string{"field1"},
	}

	cfg := &oo.Options{
		Paths: map[string]*po.Options{
			"root": rpath,
		},
	}

	newResources := func() *request.Resources {
		return request.NewResources(cfg, cfg.Paths["root"], nil, nil, nil, nil, tl.ConsoleLogger("error"))
	}

	tr := httptest.NewRequest("GET", "http://127.0.0.1/?query=12345&start=0&end=0&step=300&time=0", nil)
	tr = tr.WithContext(ct.WithResources(context.Background(), newResources()))

	pr := newProxyRequest(tr, nil)
	ck := pr.DeriveCacheKey(nil, "extra")

	if ck != "52dc11456c84506d3444e53ee4c99777" {
		t.Errorf("expected %s got %s", "52dc11456c84506d3444e53ee4c99777", ck)
	}

	cfg.Paths["root"].CacheKeyParams = []string{"*"}

	pr = newProxyRequest(tr, nil)
	ck = pr.DeriveCacheKey(pr.URL, "extra")
	if ck != "407aba34f02c87f6898a6d80b01f38a4" {
		t.Errorf("expected %s got %s", "407aba34f02c87f6898a6d80b01f38a4", ck)
	}

	const expected = "cb84ad010abb4d0f864470540a46f137"

	tr = httptest.NewRequest(http.MethodPost, "http://127.0.0.1/", bytes.NewReader([]byte("field1=value1")))
	tr = tr.WithContext(ct.WithResources(context.Background(), newResources()))
	tr.Header.Set(headers.NameContentType, headers.ValueXFormURLEncoded)
	pr = newProxyRequest(tr, nil)
	ck = pr.DeriveCacheKey(nil, "extra")
	if ck != expected {
		t.Errorf("expected %s got %s", expected, ck)
	}

	tr = httptest.NewRequest(http.MethodPut, "http://127.0.0.1/", bytes.NewReader([]byte(testMultipartBody)))
	tr = tr.WithContext(ct.WithResources(context.Background(), newResources()))
	tr.Header.Set(headers.NameContentType, headers.ValueMultipartFormData+testMultipartBoundary)
	tr.Header.Set(headers.NameContentLength, strconv.Itoa(len(testMultipartBody)))
	pr = newProxyRequest(tr, nil)
	ck = pr.DeriveCacheKey(nil, "extra")
	if ck != "4766201eee9ef1916f57309deae22f90" {
		t.Errorf("expected %s got %s", "4766201eee9ef1916f57309deae22f90", ck)
	}

	_, _, tr, _, _ = tu.NewTestInstance("", nil, 0, "", nil, "rpc", "http://127.0.0.1/", "INFO")
	tr.Method = http.MethodPost
	tr.Body = ioutil.NopCloser(bytes.NewReader([]byte(testJSONDocument)))
	tr = tr.WithContext(ct.WithResources(context.Background(), newResources()))
	tr.Header.Set(headers.NameContentType, headers.ValueApplicationJSON)
	tr.Header.Set(headers.NameContentLength, strconv.Itoa(len(testJSONDocument)))
	pr = newProxyRequest(tr, nil)

	ck = pr.DeriveCacheKey(nil, "extra")
	if ck != "82c1d86126a02b96b8d0fcb94a9f486a" {
		t.Errorf("expected %s got %s", "82c1d86126a02b96b8d0fcb94a9f486a", ck)
	}

	// Test Custom KeyHasher Integration
	rpath.KeyHasher = []key.HasherFunc{exampleKeyHasher}
	ck = pr.DeriveCacheKey(nil, "extra")
	if ck != "test-key" {
		t.Errorf("expected %s got %s", "test-key", ck)
	}

	tr = httptest.NewRequest(http.MethodPost, "http://127.0.0.1/", nil)
	tr.Body = ioutil.NopCloser(bytes.NewReader([]byte(testJSONDocument)))
	tr = tr.WithContext(ct.WithResources(context.Background(), newResources()))
	tr.Header.Set(headers.NameContentType, headers.ValueApplicationJSON)
	tr.Header.Set(headers.NameContentLength, strconv.Itoa(len(testJSONDocument)))
	pr = newProxyRequest(tr, nil)
	pr.upstreamRequest.URL = nil
	ck = pr.DeriveCacheKey(nil, "extra")
	if ck != "test-key" {
		t.Errorf("expected %s got %s", expected, ck)
	}
}

func exampleKeyHasher(path string, params url.Values, headers http.Header,
	body io.ReadCloser, extra string) (string, io.ReadCloser) {
	return "test-key", nil
}

func TestDeriveCacheKeyAuthHeader(t *testing.T) {

	client := &TestClient{
		config: &oo.Options{
			Paths: map[string]*po.Options{
				"root": {
					Path:            "/",
					CacheKeyParams:  []string{"query", "step", "time"},
					CacheKeyHeaders: []string{"X-Test-Header"},
				},
			},
		},
	}

	tr := httptest.NewRequest("GET", "http://127.0.0.1/?query=12345&start=0&end=0&step=300&time=0", nil)
	tr = tr.WithContext(ct.WithResources(context.Background(),
		request.NewResources(client.Configuration(), client.Configuration().Paths["root"],
			nil, nil, nil, nil, tl.ConsoleLogger("error"))))

	tr.Header.Add("Authorization", "test")
	tr.Header.Add("X-Test-Header", "test2")

	pr := newProxyRequest(tr, nil)

	ck := pr.DeriveCacheKey(nil, "extra")

	if ck != "60257fa6b18d6072b90a294269a8e6e1" {
		t.Errorf("expected %s got %s", "60257fa6b18d6072b90a294269a8e6e1", ck)
	}

}

func TestDeriveCacheKeyNoPathConfig(t *testing.T) {

	client := &TestClient{
		config: &oo.Options{
			Paths: map[string]*po.Options{
				"root": {
					Path:            "/",
					CacheKeyParams:  []string{"query", "step", "time"},
					CacheKeyHeaders: []string{},
				},
			},
		},
	}

	tr := httptest.NewRequest("GET", "http://127.0.0.1/?query=12345&start=0&end=0&step=300&time=0", nil)
	tr = tr.WithContext(ct.WithResources(context.Background(),
		request.NewResources(client.Configuration(), nil, nil, nil, nil, nil, tl.ConsoleLogger("error"))))

	pr := newProxyRequest(tr, nil)
	ck := pr.DeriveCacheKey(nil, "extra")

	if ck != "f53b04ce5c434a7357804ae15a64ee6c" {
		t.Errorf("expected %s got %s", "f53b04ce5c434a7357804ae15a64ee6c", ck)
	}

}

func TestDeriveCacheKeyNilURL(t *testing.T) {

	_, w, r, _, _ := tu.NewTestInstance("", nil, 0, "", nil, "rpc",
		"http://127.0.0.1/?query=12345&start=0&end=0&step=300&time=0", "INFO")

	pr := newProxyRequest(r, w)
	pr.upstreamRequest.URL = nil
	k := pr.DeriveCacheKey(nil, "")
	if k != "c04284eb2c269dd939d54437d4efb071" {
		t.Errorf("unexpected cache key: %s", k)
	}
}
