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

package types

import "github.com/haproxytech/config-parser/v2/common"

//name:section
//dir:extra
//no-init:true
type Section struct {
	Name    string
	Comment string
}

//name:config-version
//dir:extra
//no-init:true
//no-get:true
type ConfigVersion struct {
	Value int64
}

//name:comments
//dir:extra
//is-multiple:true
//no-init:true
//no-parse:true
type Comments struct {
	Value string
}

//name:unprocessed
//dir:extra
//is-multiple:true
//no-init:true
//no-parse:true
//test:skip
type UnProcessed struct {
	Value string
}

//name:simple-option
//struct-name:Option
//dir:simple
//no-init:true
type SimpleOption struct {
	NoOption bool
	Comment  string
}

//name:simple-timeout
//struct-name:Timeout
//dir:simple
//no-init:true
type SimpleTimeout struct {
	Value   string
	Comment string
}

//name:simple-word
//struct-name:Word
//dir:simple
//parser-type:StringC
type SimpleWord struct{}

//name:simple-number
//struct-name:Number
//dir:simple
//parser-type:Int64C
type SimpleNumber struct{}

//name:simple-string
//struct-name:String
//dir:simple
//parser-type:StringC
type SimpleString struct{}

//name:simple-time
//struct-name:Time
//dir:simple
//parser-type:StringC
type SimpleTime struct{}

type Filter interface {
	Parse(parts []string, comment string) error
	Result() common.ReturnResultLine
}

//name:filter
//dir:filters
//is-multiple:true
//parser-type:Filter
//is-interface:true
//no-init:true
//no-parse:true
type Filters struct{}

type HTTPAction interface {
	Parse(parts []string, comment string) error
	String() string
	GetComment() string
}

//name:http-request
//struct-name:Requests
//dir:http
//is-multiple:true
//parser-type:HTTPAction
//is-interface:true
//no-init:true
//no-parse:true
//test:fail:http-request
//test:fail:http-request capture req.cook_cnt(FirstVisit),bool strlen 10
//test:ok:http-request capture req.cook_cnt(FirstVisit),bool len 10
//test:ok:http-request deny deny_status 0 unless { src 127.0.0.1 }
type HTTPRequests struct{}

//name:http-response
//struct-name:Responses
//dir:http
//is-multiple:true
//parser-type:HTTPAction
//is-interface:true
//no-init:true
//no-parse:true
//test:fail:http-response
//test:ok:http-response capture res.hdr(Server) id 0
type HTTPResponses struct{}

type TCPAction interface {
	Parse(parts []string, comment string) error
	String() string
	GetComment() string
}

//name:tcp-request
//struct-name:Requests
//dir:tcp
//is-multiple:true
//parser-type:TCPAction
//is-interface:true
//no-init:true
//no-parse:true
//test:fail:tcp-request
type TCPRequests struct{}

//name:tcp-response
//struct-name:Responses
//dir:tcp
//is-multiple:true
//parser-type:TCPAction
//is-interface:true
//no-init:true
//no-parse:true
//test:fail:tcp-response
type TCPResponses struct{}

//name:redirect
//dir:http
//is-multiple:true
//parser-type:HTTPAction
//is-interface:true
//no-init:true
//no-parse:true
//test:fail:redirect
//test:ok:redirect prefix http://www.bar.com code 301 if { hdr(host) -i foo.com }
type Redirect struct{}
