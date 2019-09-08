package utils

import (
	"github.com/drakkan/sftpgo/utils"
	"testing"
)

func TestCheckPermissinoForPath(t *testing.T) {
	var perms = []string{
		"/:list",
		"/t/1:read",
		"/t/1: write",
		"/t/3/4:write",
		"/t:read",
		"/t/3:read",
	}

	if !utils.CheckPermissinoForPath("list", "/", perms) {
		t.Error("no match patch / against *")
	}
	// not defined path - inherit from /
	if !utils.CheckPermissinoForPath("list", "/test/2", perms) {
		t.Error("no match patch /test/2 against *")
	}
	// / has no read permission
	if utils.CheckPermissinoForPath("read", "/", perms) {
		t.Error("match patch / against read")
	}
	// /t/2 inherit read from /t/
	if !utils.CheckPermissinoForPath("read", "/t/2", perms) {
		t.Error("no match patch /t/2 against read")
	}
	// /t/2/ has not write permission
	if utils.CheckPermissinoForPath("write", "/t/2", perms) {
		t.Error("match patch /t/2 against write")
	}
	// /t/1 has write permission
	if !utils.CheckPermissinoForPath("write", "/t/1", perms) {
		t.Error("no match patch /t/1 against write")
	}
	// so do /t/1/3
	if !utils.CheckPermissinoForPath("write", "/t/1/3", perms) {
		t.Error("no match patch /t/1/3 against write")
	}
	// /t/3 has only read and list
	if utils.CheckPermissinoForPath("write", "/t/3", perms) {
		t.Error("match patch /t/3 against write")
	}
}
