package enfins

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"net/url"
	"fmt"
)

type QueryBuilder struct {
	url    *url.URL
	cfg    *Configuration
	params url.Values
}

// Create new QueryBuilder with a configuration
func NewQuery(path string, cfg *Configuration) (*QueryBuilder, error) {
	u, err := url.Parse(cfg.Scheme+ "://" + cfg.Host + "/" + cfg.Version + "/")
	if err != nil {
		return nil, err
	}
	qb := &QueryBuilder{
		u,
		cfg,
		url.Values{},
	}
	qb.url.Path = qb.url.Path + path
	qb.AddParam("ident", qb.cfg.Ident)
	return qb, nil
}

// Add param to a query map if name is exist it will be replaced
func (qb *QueryBuilder) AddParam(name string, value string) {
	qb.params.Set(name, value)
	// update Sort by Key
	qb.url.RawQuery = qb.params.Encode()
}

// Adds sign to query
func (qb *QueryBuilder) signQuery() {
	qb.params.Del("sign") // if exist
	mac := hmac.New(sha1.New, []byte(qb.cfg.SecretKey))
	mac.Write([]byte(qb.params.Encode()))
	qb.AddParam("sign", hex.EncodeToString(mac.Sum(nil)))
}

// returns URL for request
func (qb *QueryBuilder) Url() string {
	qb.signQuery()
	fmt.Println(qb.url.String())
	return qb.url.String()
}


