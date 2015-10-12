package mesosutil

import (
	"os/exec"
	"strings"

	log "github.com/golang/glog"
)

//TODO(jdef) copied from kubernetes/pkg/util/node.go
func GetHostname(hostnameOverride string) string {
	hostname := hostnameOverride
	if string(hostname) == "" {
		// Note: We use exec here instead of os.Hostname() because we
		// want the FQDN, and this is the easiest way to get it.
		nodename, err := exec.Command("uname", "-n").Output()
		if err != nil {
			log.Fatalf("Couldn't determine hostname: %v", err)
		}
		chunks := strings.Split(string(nodename), ".")
		// nodename could be a fully-qualified domain name or not. Take the
		// first word of nodename as the hostname for consistency.
		hostname = chunks[0]
	}
	return strings.TrimSpace(string(hostname))
}
