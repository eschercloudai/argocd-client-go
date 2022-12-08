/*
Copyright 2022 EscherCloud.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/eschercloudai/argocd-client-go/pkg/api"
)

var (
	// ErrCertificate is raised when a supplied certificate isn't
	// in the right format etc.
	ErrCertificate = errors.New("certificate error")

	// ErrAuthentication is raised when something goes wrong with authentication.
	ErrAuthentication = errors.New("authentication error")
)

// Options define client options to define the connection.
type Options struct {
	// Host is the hostname to connect to.
	Host string

	// Username is the user to connect as.
	Username string

	// Password is the user's password.
	Password string

	// CA is the certificate to verify the connection against.
	// This is a single PEM encoded X.509 certificate.
	CA []byte

	// UserAgent allows the user agent to be overridden.
	UserAgent string
}

// parseCA takes a raw PEM encoded certificate and turns it into an x509
// native certificate type.
func parseCA(ca []byte) (*x509.Certificate, error) {
	block, next := pem.Decode(ca)
	if block == nil {
		return nil, fmt.Errorf("%w: unable to decode CA PEM", ErrCertificate)
	}

	if len(next) != 0 {
		return nil, fmt.Errorf("%w: CA PEM contains extra data", ErrCertificate)
	}

	if block.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("%w: CA PEM not of type CERTIFICATE", ErrCertificate)
	}

	caCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("unable to parse CA certficate: %w", err)
	}

	return caCert, nil
}

// httpsClient gets a new HTTPS client.  This is a default client with
// ostensibly a copy of go's DefaultTransport, and TLS applied to it,
// because we trust Google knows how to network.
func httpsClient(ca *x509.Certificate) *http.Client {
	caCertPool := x509.NewCertPool()
	caCertPool.AddCert(ca)

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:    caCertPool,
				MinVersion: tls.VersionTLS13,
			},
			DialContext:           dialer.DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: time.Second,
		},
	}

	return client
}

// New returns a new ArgoCD client initialized with a service token.
func New(ctx context.Context, options *Options) (*api.APIClient, error) {
	ca, err := parseCA(options.CA)
	if err != nil {
		return nil, err
	}

	cfg := api.NewConfiguration()
	cfg.Host = options.Host
	cfg.Scheme = "https"
	cfg.UserAgent = options.UserAgent
	cfg.HTTPClient = httpsClient(ca)

	client := api.NewAPIClient(cfg)

	// Do the Basic auth to Bearer token auth dance, and set the Authorization header.
	sessionCreateRequest := api.NewSessionSessionCreateRequest()
	sessionCreateRequest.SetUsername(options.Username)
	sessionCreateRequest.SetPassword(options.Password)

	sessionServiceCreate := client.SessionServiceApi.SessionServiceCreate(ctx)
	sessionServiceCreate = sessionServiceCreate.Body(*sessionCreateRequest)

	response, httpResponse, err := sessionServiceCreate.Execute()
	defer httpResponse.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("unable to create sessison: %w", err)
	}

	token, ok := response.GetTokenOk()
	if !ok {
		return nil, fmt.Errorf("%w: no token in session create response", ErrAuthentication)
	}

	cfg.AddDefaultHeader("Authorization", "Bearer "+*token)

	return client, nil
}
