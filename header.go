package gomitmproxy

import (
	"github.com/AdguardTeam/golibs/log"
	"github.com/minimAluminiumalism/gomitmproxy/proxyutil"
)

func headerConfig(session *Session) {
	configPath := "./headers.yaml"
	_, err := proxyutil.ParseConfig(configPath)
	if err != nil {
		log.Error("Header configuration error, ", err)
	}
	// userAgent := 
	// cookie := 
	session.req.Header = map[string][]string{"cookie": {"csrftoken=YisZfC5ccRSFnSWlUfd2fp8l7S25gKtu26EggpuyqUtC3v5pSy54baeZNzvstcUX; sessionid=lmp90aiknolao6yaluugtp7stdlg4cgm"}}
}
