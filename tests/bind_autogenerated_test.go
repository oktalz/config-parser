// Code generated by go generate; DO NOT EDIT.
package tests

import (
	"testing"

	"github.com/haproxytech/config-parser/parsers"
)


func TestBindNormal0(t *testing.T) {
	err := ProcessLine("bind :80,:443", &parsers.Bind{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestBindNormal1(t *testing.T) {
	err := ProcessLine("bind 10.0.0.1:10080,10.0.0.1:10443", &parsers.Bind{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestBindNormal2(t *testing.T) {
	err := ProcessLine("bind /var/run/ssl-frontend.sock user root mode 600 accept-proxy", &parsers.Bind{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestBindNormal3(t *testing.T) {
	err := ProcessLine("bind :80", &parsers.Bind{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestBindNormal4(t *testing.T) {
	err := ProcessLine("bind :443 ssl crt /etc/haproxy/site.pem", &parsers.Bind{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestBindNormal5(t *testing.T) {
	err := ProcessLine("bind ipv6@:80", &parsers.Bind{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestBindNormal6(t *testing.T) {
	err := ProcessLine("bind ipv4@public_ssl:443 ssl crt /etc/haproxy/site.pem", &parsers.Bind{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestBindNormal7(t *testing.T) {
	err := ProcessLine("bind unix@ssl-frontend.sock user root mode 600 accept-proxy", &parsers.Bind{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestBindFail0(t *testing.T) {
	err := ProcessLine("bind", &parsers.Bind{})
	if err == nil {
		t.Errorf("no data")
	}
}
func TestBindFail1(t *testing.T) {
	err := ProcessLine("---", &parsers.Bind{})
	if err == nil {
		t.Errorf("no data")
	}
}
func TestBindFail2(t *testing.T) {
	err := ProcessLine("--- ---", &parsers.Bind{})
	if err == nil {
		t.Errorf("no data")
	}
}
