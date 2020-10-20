// Code generated by go generate; DO NOT EDIT.
/*
Copyright 2019 HAProxy Technologies

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

package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/haproxytech/config-parser/v3/parsers"
)

func TestServer(t *testing.T) {
	tests := map[string]bool{
		"server name 127.0.0.1:8080": true,
		"server name 127.0.0.1": true,
		"server addr 127.0.0.1": true,
		"server addr ::1": true,
		"server name 127.0.0.1 agent-check": true,
		"server name 127.0.0.1 agent-send name": true,
		"server name 127.0.0.1 agent-inter 1000ms": true,
		"server name 127.0.0.1 agent-addr 127.0.0.1": true,
		"server name 127.0.0.1 agent-addr site.com": true,
		"server name 127.0.0.1 agent-port 1": true,
		"server name 127.0.0.1 agent-port 65535": true,
		"server name 127.0.0.1 allow-0rtt": true,
		"server name 127.0.0.1 alpn h2": true,
		"server name 127.0.0.1 alpn http/1.1": true,
		"server name 127.0.0.1 alpn h2,http/1.1": true,
		"server name 127.0.0.1 backup": true,
		"server name 127.0.0.1 ca-file cert.crt": true,
		"server name 127.0.0.1 check": true,
		"server name 127.0.0.1 check-send-proxy": true,
		"server name 127.0.0.1 check-alpn http/1.0": true,
		"server name 127.0.0.1 check-alpn http/1.1,http/1.0": true,
		"server name 127.0.0.1 check-proto h2": true,
		"server name 127.0.0.1 check-ssl": true,
		"server name 127.0.0.1 check-via-socks4": true,
		"server name 127.0.0.1 ciphers ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA:ECDHE-RSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-RSA-AES256-SHA256:DHE-RSA-AES256-SHA:ECDHE-ECDSA-DES-CBC3-SHA:ECDHE-RSA-DES-CBC3-SHA:EDH-RSA-DES-CBC3-SHA:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:DES-CBC3-SHA:!DSS": true,
		"server name 127.0.0.1 ciphersuites ECDHE-ECDSA-CHACHA20-POLY1305:ECDHE-RSA-CHACHA20-POLY1305:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA:ECDHE-RSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-RSA-AES256-SHA256:DHE-RSA-AES256-SHA:ECDHE-ECDSA-DES-CBC3-SHA:ECDHE-RSA-DES-CBC3-SHA:EDH-RSA-DES-CBC3-SHA:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:DES-CBC3-SHA:!DSS": true,
		"server name 127.0.0.1 cookie value": true,
		"server name 127.0.0.1 crl-file file.pem": true,
		"server name 127.0.0.1 crt cert.pem": true,
		"server name 127.0.0.1 disabled": true,
		"server name 127.0.0.1 enabled": true,
		"server name 127.0.0.1 error-limit 50": true,
		"server name 127.0.0.1 fall 30": true,
		"server name 127.0.0.1 force-sslv3": true,
		"server name 127.0.0.1 force-tlsv10": true,
		"server name 127.0.0.1 force-tlsv11": true,
		"server name 127.0.0.1 force-tlsv12": true,
		"server name 127.0.0.1 force-tlsv13": true,
		"server name 127.0.0.1 init-addr last,libc,none": true,
		"server name 127.0.0.1 init-addr last,libc,none,127.0.0.1": true,
		"server name 127.0.0.1 inter 1500ms": true,
		"server name 127.0.0.1 fastinter 2500ms": true,
		"server name 127.0.0.1 fastinter unknown": true,
		"server name 127.0.0.1 downinter 3500ms": true,
		"server name 127.0.0.1 log-proto legacy": true,
		"server name 127.0.0.1 log-proto octet-count": true,
		"server name 127.0.0.1 maxconn 1": true,
		"server name 127.0.0.1 maxconn 50": true,
		"server name 127.0.0.1 maxqueue 0": true,
		"server name 127.0.0.1 maxqueue 1000": true,
		"server name 127.0.0.1 max-reuse -1": true,
		"server name 127.0.0.1 max-reuse 0": true,
		"server name 127.0.0.1 max-reuse 1": true,
		"server name 127.0.0.1 minconn 1": true,
		"server name 127.0.0.1 minconn 50": true,
		"server name 127.0.0.1 namespace test": true,
		"server name 127.0.0.1 no-agent-check": true,
		"server name 127.0.0.1 no-backup": true,
		"server name 127.0.0.1 no-check": true,
		"server name 127.0.0.1 no-check-ssl": true,
		"server name 127.0.0.1 no-send-proxy-v2": true,
		"server name 127.0.0.1 no-send-proxy-v2-ssl": true,
		"server name 127.0.0.1 no-send-proxy-v2-ssl-cn": true,
		"server name 127.0.0.1 no-ssl": true,
		"server name 127.0.0.1 no-ssl-reuse": true,
		"server name 127.0.0.1 no-sslv3": true,
		"server name 127.0.0.1 no-tls-tickets": true,
		"server name 127.0.0.1 no-tlsv10": true,
		"server name 127.0.0.1 no-tlsv11": true,
		"server name 127.0.0.1 no-tlsv12": true,
		"server name 127.0.0.1 no-tlsv13": true,
		"server name 127.0.0.1 no-verifyhost": true,
		"server name 127.0.0.1 no-tfo": true,
		"server name 127.0.0.1 non-stick": true,
		"server name 127.0.0.1 npn http/1.1,http/1.0": true,
		"server name 127.0.0.1 observe layer4": true,
		"server name 127.0.0.1 observe layer7": true,
		"server name 127.0.0.1 on-error fastinter": true,
		"server name 127.0.0.1 on-error fail-check": true,
		"server name 127.0.0.1 on-error sudden-death": true,
		"server name 127.0.0.1 on-error mark-down": true,
		"server name 127.0.0.1 on-marked-down shutdown-sessions": true,
		"server name 127.0.0.1 on-marked-up shutdown-backup-session": true,
		"server name 127.0.0.1 pool-max-conn -1": true,
		"server name 127.0.0.1 pool-max-conn 0": true,
		"server name 127.0.0.1 pool-max-conn 100": true,
		"server name 127.0.0.1 pool-purge-delay 0": true,
		"server name 127.0.0.1 pool-purge-delay 5": true,
		"server name 127.0.0.1 pool-purge-delay 500": true,
		"server name 127.0.0.1 port 27015": true,
		"server name 127.0.0.1 port 27016": true,
		"server name 127.0.0.1 proto h2": true,
		"server name 127.0.0.1 redir http://image1.mydomain.com": true,
		"server name 127.0.0.1 redir https://image1.mydomain.com": true,
		"server name 127.0.0.1 rise 2": true,
		"server name 127.0.0.1 rise 200": true,
		"server name 127.0.0.1 resolve-opts allow-dup-ip": true,
		"server name 127.0.0.1 resolve-opts ignore-weight": true,
		"server name 127.0.0.1 resolve-opts allow-dup-ip,ignore-weight": true,
		"server name 127.0.0.1 resolve-opts prevent-dup-ip,ignore-weight": true,
		"server name 127.0.0.1 resolve-prefer ipv4": true,
		"server name 127.0.0.1 resolve-prefer ipv6": true,
		"server name 127.0.0.1 resolve-net 10.0.0.0/8": true,
		"server name 127.0.0.1 resolve-net 10.0.0.0/8,10.0.0.0/16": true,
		"server name 127.0.0.1 resolvers mydns": true,
		"server name 127.0.0.1 send-proxy": true,
		"server name 127.0.0.1 send-proxy-v2": true,
		"server name 127.0.0.1 proxy-v2-options ssl": true,
		"server name 127.0.0.1 proxy-v2-options ssl,cert-cn": true,
		"server name 127.0.0.1 proxy-v2-options ssl,cert-cn,ssl-cipher,cert-sig,cert-key,authority,crc32c,unique-id": true,
		"server name 127.0.0.1 send-proxy-v2-ssl": true,
		"server name 127.0.0.1 send-proxy-v2-ssl-cn": true,
		"server name 127.0.0.1 slowstart 2000ms": true,
		"server name 127.0.0.1 sni TODO": true,
		"server name 127.0.0.1 source TODO": true,
		"server name 127.0.0.1 ssl": true,
		"server name 127.0.0.1 ssl-max-ver SSLv3": true,
		"server name 127.0.0.1 ssl-max-ver TLSv1.0": true,
		"server name 127.0.0.1 ssl-max-ver TLSv1.1": true,
		"server name 127.0.0.1 ssl-max-ver TLSv1.2": true,
		"server name 127.0.0.1 ssl-max-ver TLSv1.3": true,
		"server name 127.0.0.1 ssl-min-ver SSLv3": true,
		"server name 127.0.0.1 ssl-min-ver TLSv1.0": true,
		"server name 127.0.0.1 ssl-min-ver TLSv1.1": true,
		"server name 127.0.0.1 ssl-min-ver TLSv1.2": true,
		"server name 127.0.0.1 ssl-min-ver TLSv1.3": true,
		"server name 127.0.0.1 ssl-reuse": true,
		"server name 127.0.0.1 stick": true,
		"server name 127.0.0.1 socks4 127.0.0.1:81": true,
		"server name 127.0.0.1 tcp-ut 20ms": true,
		"server name 127.0.0.1 tfo": true,
		"server name 127.0.0.1 track TODO": true,
		"server name 127.0.0.1 tls-tickets": true,
		"server name 127.0.0.1 verify none": true,
		"server name 127.0.0.1 verify required": true,
		"server name 127.0.0.1 verifyhost site.com": true,
		"server name 127.0.0.1 weight 1": true,
		"server name 127.0.0.1 weight 128": true,
		"server name 127.0.0.1 weight 256": true,
		"server addr": false,
		"server": false,
		"---": false,
		"--- ---": false,
	}
	parser := &parsers.Server{}
	for command, shouldPass := range tests {
		t.Run(command, func(t *testing.T) {
		line :=strings.TrimSpace(command)
		lines := strings.SplitN(line,"\n", -1)
		var err error
		parser.Init()
		if len(lines)> 1{
			for _,line = range(lines){
			  line = strings.TrimSpace(line)
				if err=ProcessLine(line, parser);err!=nil{
					break
				}
			}
		}else{
			err = ProcessLine(line, parser)
		}
			if shouldPass {
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				result, err := parser.Result()
				if err != nil {
					t.Errorf(err.Error())
					return
				}
				var returnLine string
				if result[0].Comment == "" {
					returnLine = result[0].Data
				} else {
					returnLine = fmt.Sprintf("%s # %s", result[0].Data, result[0].Comment)
				}
				if command != returnLine {
					t.Errorf(fmt.Sprintf("error: has [%s] expects [%s]", returnLine, command))
				}
			} else {
				if err == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error for line [%s]", line))
				}
				_, parseErr := parser.Result()
				if parseErr == nil {
					t.Errorf(fmt.Sprintf("error: did not throw error on result for line [%s]", line))
				}
			}
		})
	}
}
