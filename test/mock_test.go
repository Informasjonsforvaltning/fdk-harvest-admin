package test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"go.mongodb.org/mongo-driver/bson"
)

type MockDataSourceRepository struct {
	MockDataSource *model.DataSource
	MockError      error
}

func (m *MockDataSourceRepository) GetDataSources(ctx context.Context, query bson.D) ([]model.DataSource, error) {
	var dataSources []model.DataSource
	if m.MockDataSource != nil {
		dataSources = append(dataSources, *m.MockDataSource)
	}
	return dataSources, m.MockError
}

func (m *MockDataSourceRepository) GetDataSource(ctx context.Context, id string) (*model.DataSource, error) {
	return m.MockDataSource, m.MockError
}

func (m *MockDataSourceRepository) DeleteDataSource(ctx context.Context, id string) error {
	return m.MockError
}

func (m *MockDataSourceRepository) CreateDataSource(ctx context.Context, dataSource model.DataSource) error {
	return m.MockError
}

type MockPublisher struct {
	MockError error
}

func (m *MockPublisher) Publish(routingKey string, data []byte) error {
	return m.MockError
}

type MockResponseWriter struct {
	MockHeader         map[string][]string
	MockStatusCode     int
	MockError          error
	CurrentStatusCode  int
	CurrentWriteOutput []byte
}

var MockRsaKey, _ = rsa.GenerateKey(rand.Reader, 2048)

func MockJwkStore() *httptest.Server {
	key, err := jwk.New(MockRsaKey)
	if err != nil {
		fmt.Println(err)
	}

	server := httptest.NewServer(
		http.HandlerFunc(
			func(rw http.ResponseWriter, r *http.Request) {
				rw.Header().Add("Content-Type", "application/json")

				key.Set(jwk.KeyIDKey, "testkid")

				buf, err := json.MarshalIndent(key, "", "  ")
				if err != nil {
					fmt.Printf("failed to marshal key into JSON: %s\n", err)
					return
				}

				fmt.Fprintf(rw, `{"keys":[%s]}`, buf)
			},
		),
	)

	return server
}

func CreateMockJwt(expiresAt int64, auth *string, audience *[]string) *string {
	t := jwt.New()
	t.Set(jwt.SubjectKey, `https://github.com/lestrrat-go/jwx/jwt`)
	t.Set(jwt.IssuedAtKey, time.Now().Unix())
	t.Set(jwt.ExpirationKey, expiresAt)
	if auth != nil {
		t.Set(`authorities`, auth)
	}
	if audience != nil {
		t.Set(jwt.AudienceKey, *audience)
	}

	jwk_key, _ := jwk.New(MockRsaKey)

	jwk_key.Set(jwk.KeyIDKey, "testkid")

	signed, _ := jwt.Sign(t, jwa.RS256, jwk_key)

	signed_string := string(signed)

	return &signed_string
}
