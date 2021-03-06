// Copyright 2011 Cloud Instruments Co. Ltd. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	log "github.com/cihub/seelog"
	"fmt"
)

func main() {
	defer log.Flush()
	defaultFormat()
	stdFormat()
	dateTimeFormat()
	dateTimeCustomFormat()
	logLevelTypesFormat()
	fileTypesFormat()
	funcFormat()
	xmlFormat()
}

func defaultFormat() {
	fmt.Println("Default format")
	
	testConfig := `
<seelog type="sync" />`

	logger, err := log.LoggerFromConfigAsBytes([]byte(testConfig))
	if err != nil {
		fmt.Println(err)
	}
	log.UseLogger(logger)
	
	log.Trace("Test message!")
}

func stdFormat() {
	fmt.Println("Standard fast format")
	
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%Ns [%Level] %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.UseLogger(logger)
	
	log.Trace("Test message!")
}

func dateTimeFormat() {
	fmt.Println("Date time format")
	
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%Date/%Time [%Level] %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.UseLogger(logger)
	
	log.Trace("Test message!")
}

func dateTimeCustomFormat() {
	fmt.Println("Date time custom format")
	
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%Date(2006 Jan 02/3:04:05.000000000 PM MST) [%Level] %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.UseLogger(logger)
	
	log.Trace("Test message!")
}

func logLevelTypesFormat() {
	fmt.Println("Log level types format")
	
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%Level %Lev %LEVEL %LEV %l %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.UseLogger(logger)
	
	log.Trace("Test message!")
}

func fileTypesFormat() {
	fmt.Println("File types format")
	
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%File %FullPath %RelFile %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.UseLogger(logger)
	
	log.Trace("Test message!")
}

func funcFormat() {
	fmt.Println("Func format")
	
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%Func %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.UseLogger(logger)
	
	log.Trace("Test message!")
}

func xmlFormat() {
	fmt.Println("Xml format")
	
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="` +
		`&lt;log&gt;` +
			 `&lt;time&gt;%Ns&lt;/time&gt;` +
			 `&lt;lev&gt;%l&lt;/lev&gt;` +
			 `&lt;msg&gt;%Msg&lt;/msg&gt;` +
		 `&lt;/log&gt;"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))

	log.UseLogger(logger)
	
	log.Trace("Test message!")
}