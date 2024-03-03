package main

import (
	"github.com/dchest/captcha"
	"testing"
)

func TestChallengeUserSuccess(t *testing.T) {
	got := captcha.ChallengeUser(stubChallenger("42"), stubPrompter("42"))

	if got != true {
		t.Fatal("Expected ChallengeUser to return true")
	}
}
