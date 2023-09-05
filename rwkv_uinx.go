// Copyright (c) seasonjs. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

//go:build darwin || linux

package rwkv

import (
	"github.com/ebitengine/purego"
)

var libraryPath = "./deps/darwin/librwkv.dylib"

func openLibrary(name string) (uintptr, error) {
	return purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}