package envoy

import (
	"os"
	"os/exec"
	"testing"
)

func TestEnvoy(t *testing.T) {
	var testcases = []string{
		"case-badauthz",
		"case-basic",
		"case-centralconf",
		"case-cfg-resolver-dc-failover-gateways-none",
		"case-cfg-resolver-dc-failover-gateways-remote",
		"case-cfg-resolver-defaultsubset",
		"case-cfg-resolver-subset-onlypassing",
		"case-cfg-resolver-subset-redirect",
		"case-cfg-resolver-svc-failover",
		"case-cfg-resolver-svc-redirect-http",
		"case-cfg-resolver-svc-redirect-tcp",
		"case-consul-exec",
		"case-dogstatsd-udp",
		"case-gateways-local",
		"case-gateways-remote",
		"case-gateway-without-services",
		"case-grpc",
		"case-http",
		"case-http2",
		"case-http-badauthz",
		"case-ingress-gateway-http",
		"case-ingress-gateway-multiple-services",
		"case-ingress-gateway-simple",
		"case-ingress-mesh-gateways-resolver",
		"case-multidc-rsa-ca",
		"case-prometheus",
		"case-statsd-udp",
		"case-stats-proxy",
		"case-terminating-gateway-simple",
		"case-terminating-gateway-subsets",
		"case-terminating-gateway-without-services",
		"case-upstream-config",
		"case-wanfed-gw",
		"case-zipkin",
	}

	suiteSetup(t)
	defer suiteTearDown(t)

	for _, tc := range testcases {
		t.Run(tc, func(t *testing.T) {
			args := []string{"./run-tests.sh"}
			env := append(os.Environ(), "CASE_DIR="+tc)
			runCmd(t, args, env)
		})
	}
}

func runCmd(t *testing.T, args []string, env []string) {
	t.Helper()

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		t.Fatalf("command failed: %v", err)
	}
}

func suiteSetup(t *testing.T) {
	// TODO: set LOG_DIR and ENVOY_VERSION to prevent warnings from compose
	suiteTearDown(t)
	args := []string{"docker-compose", "up", "-d", "workdir"}
	runCmd(t, args, nil)
}

func suiteTearDown(t *testing.T) {
	// TODO: set LOG_DIR and ENVOY_VERSION to prevent warnings from compose
	args := []string{"docker-compose", "down", "--volumes", "--timeout=0", "--remove-orphans"}
	runCmd(t, args, nil)
}
