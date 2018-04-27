// Flagger - arbitrary CLI flags parser.
//
// Copyright (c) 2017-2018, Stanislav N. aka pztrn.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject
// to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package flagger

import (
	// stdlib
	"log"
	"os"
	"testing"
)

var (
	f *Flagger
)

func TestFlaggerInitialization(t *testing.T) {
	f = New(LoggerInterface(log.New(os.Stdout, "testing logger: ", log.Lshortfile)))
	if f == nil {
		t.Fatal("Logger initialization failed!")
		t.FailNow()
	}
	f.Initialize()
}

func TestFlaggerInitializationWithNilLogger(t *testing.T) {
	fl := New(nil)
	if f == nil {
		t.Fatal("Logger initialization failed!")
		t.FailNow()
	}
	fl.Initialize()
}

func TestFlaggerAddBoolFlag(t *testing.T) {
	flagTestBool := Flag{
		Name:         "testboolflag",
		Description:  "Testing boolean flag",
		Type:         "bool",
		DefaultValue: true,
	}
	err := f.AddFlag(&flagTestBool)
	if err != nil {
		t.Fatal("Failed to add boolean flag!")
		t.FailNow()
	}
}

func TestFlaggerAddIntFlag(t *testing.T) {
	flagTestInt := Flag{
		Name:         "testintflag",
		Description:  "Testing integer flag",
		Type:         "int",
		DefaultValue: 1,
	}
	err := f.AddFlag(&flagTestInt)
	if err != nil {
		t.Fatal("Failed to add integer flag!")
		t.FailNow()
	}
}

func TestFlaggerAddStringFlag(t *testing.T) {
	flagTestString := Flag{
		Name:         "teststringflag",
		Description:  "Testing string flag",
		Type:         "string",
		DefaultValue: "superstring",
	}
	err := f.AddFlag(&flagTestString)
	if err != nil {
		t.Fatal("Failed to add string flag!")
		t.FailNow()
	}
}

// This test doing nothing more but launching flags parsing.
func TestFlaggerParse(t *testing.T) {
	f.Parse()
}

func TestFlaggerGetBoolFlag(t *testing.T) {
	val, err := f.GetBoolValue("testboolflag")
	if err != nil {
		t.Fatal("Failed to get boolean flag: " + err.Error())
		t.FailNow()
	}

	if !val {
		t.Fatal("Failed to get boolean flag - should be true, but false received")
		t.FailNow()
	}
}

func TestFlaggerGetIntFlag(t *testing.T) {
	val, err := f.GetIntValue("testintflag")
	if err != nil {
		t.Fatal("Failed to get integer flag: " + err.Error())
		t.FailNow()
	}

	if val == 0 {
		t.Fatal("Failed to get integer flag - should be 1, but 0 received")
		t.FailNow()
	}
}

func TestFlaggerGetStringFlag(t *testing.T) {
	val, err := f.GetStringValue("teststringflag")
	if err != nil {
		t.Fatal("Failed to get string flag: " + err.Error())
		t.FailNow()
	}

	if val == "" {
		t.Fatal("Failed to get string flag - should be 'superstring', but nothing received")
		t.FailNow()
	}
}
