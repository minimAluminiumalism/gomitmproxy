package gomitmproxy

import (
	"github.com/AdguardTeam/golibs/log"
	"github.com/minimAluminiumalism/gomitmproxy/proxyutil"
)

func headerConfig(session *Session) {
	configPath := "./headers.yaml"
	header, err := proxyutil.ParseConfig(configPath)
	if err != nil {
		log.Error("Header configuration error, ", err)
	}
	session.req.Header = make(map[string][]string)
	session.req.Header["user-agent"] = []string{header.UserAgent}
	session.req.Header["cookie"] = []string{header.Cookie}
}
