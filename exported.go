// Flagger - arbitrary CLI flags parser.
//
// Copyright (c) 2017-2019, Stanislav N. aka pztrn.
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
	"log"
	"os"
)

var (
	logger          LoggerInterface
	applicationName string
)

// New creates new Flagger instance.
// If no logger will be passed - we will use default "log" module and will
// print logs to stdout.
func New(appName string, l LoggerInterface) *Flagger {
	applicationName = appName

	if l == nil {
		lg := log.New(os.Stdout, "Flagger: ", log.LstdFlags)
		logger = LoggerInterface(lg)
	} else {
		logger = l
	}

	// nolint:exhaustruct
	f := Flagger{}

	return &f
}
