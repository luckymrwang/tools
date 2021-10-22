package iptables_exec

import (
	"testing"

	utilexec "k8s.io/utils/exec"
)

func TestEnsureRuleAlreadyExists(t *testing.T) {
	// Create a iptables utils.
	execer := utilexec.New()
	runner := New(execer, ProtocolIpv4)
	exists, err := runner.EnsureRule(Append, TableNAT, ChainOutput, "abc", "123")
	if err != nil {
		t.Errorf("expected success, got %v", err)
	}
	if !exists {
		t.Errorf("expected exists = true")
	}
}
