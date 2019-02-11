// Code generated by go generate; DO NOT EDIT.
package tests

import (
	"testing"

	"github.com/haproxytech/config-parser/parsers"
)


func TestMailerNormal0(t *testing.T) {
	err := ProcessLine("mailer smtp1 192.168.0.1:587", &parsers.Mailer{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestMailerNormal1(t *testing.T) {
	err := ProcessLine("mailer smtp1 192.168.0.1:587 # just some comment", &parsers.Mailer{})
	if err != nil {
		t.Errorf(err.Error())
	}
}
func TestMailerFail0(t *testing.T) {
	err := ProcessLine("mailer", &parsers.Mailer{})
	if err == nil {
		t.Errorf("no data")
	}
}
func TestMailerFail1(t *testing.T) {
	err := ProcessLine("---", &parsers.Mailer{})
	if err == nil {
		t.Errorf("no data")
	}
}
func TestMailerFail2(t *testing.T) {
	err := ProcessLine("--- ---", &parsers.Mailer{})
	if err == nil {
		t.Errorf("no data")
	}
}
