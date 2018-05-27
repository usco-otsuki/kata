package friend_cheating_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFriendCheating(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FriendCheating Suite")
}
