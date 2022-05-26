package main

import (
	"context"
	"crypto/x509"
	"net/url"
	"testing"

	"github.com/sigstore/fulcio/pkg/ca/x509ca"
)

type principal1 struct {
	// Subject ('sub') from ID token
	subject string

	// Issuer ('iss') from ID token
	issuer string

	// URI to be set in certificate. URI is of the form
	// https://kubernetes.io/namespaces/<namespace>/serviceaccounts/<serviceaccount>.
	uri string
}

func (p principal1) Name(context.Context) string {
	return p.subject
}

func (p principal1) Embed(ctx context.Context, cert *x509.Certificate) error {
	parsed, err := url.Parse(p.uri)
	if err != nil {
		return err
	}
	cert.URIs = []*url.URL{parsed}

	cert.ExtraExtensions, err = x509ca.Extensions{
		Issuer: p.issuer,
	}.Render()
	if err != nil {
		return err
	}

	return nil
}

type principal2 struct {
	// Subject ('sub') from ID token
	subject string

	// Issuer ('iss') from ID token
	issuer string

	// URI to be set in certificate. URI is of the form
	// https://kubernetes.io/namespaces/<namespace>/serviceaccounts/<serviceaccount>.
	uri string
}

func (p *principal2) Name(context.Context) string {
	return p.subject
}

func (p *principal2) Embed(ctx context.Context, cert *x509.Certificate) error {
	parsed, err := url.Parse(p.uri)
	if err != nil {
		return err
	}
	cert.URIs = []*url.URL{parsed}

	cert.ExtraExtensions, err = x509ca.Extensions{
		Issuer: p.issuer,
	}.Render()
	if err != nil {
		return err
	}

	return nil
}

func BenchmarkPrincipal1Name(b *testing.B) {
	p := principal1{
		subject: "foo",
		issuer:  "https://example.com",
		uri:     "https://example.com/foo/bar/baz",
	}
	for n := 0; n < b.N; n++ {
		p.Name(context.TODO())
	}
}

func BenchmarkPrincipal1Embed(b *testing.B) {
	p := principal1{
		subject: "foo",
		issuer:  "https://example.com",
		uri:     "https://example.com/foo/bar/baz",
	}
	for n := 0; n < b.N; n++ {
		p.Embed(context.TODO(), &x509.Certificate{})
	}
}

func BenchmarkPrincipal2Name(b *testing.B) {
	p := principal2{
		subject: "foo",
		issuer:  "https://example.com",
		uri:     "https://example.com/foo/bar/baz",
	}
	for n := 0; n < b.N; n++ {
		p.Name(context.TODO())
	}
}

func BenchmarkPrincipal2Embed(b *testing.B) {
	p := principal2{
		subject: "foo",
		issuer:  "https://example.com",
		uri:     "https://example.com/foo/bar/baz",
	}
	for n := 0; n < b.N; n++ {
		p.Embed(context.TODO(), &x509.Certificate{})
	}
}
