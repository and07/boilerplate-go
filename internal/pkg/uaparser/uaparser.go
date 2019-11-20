package uaparser

import (
	"fmt"

	"github.com/ua-parser/uap-go/uaparser"

	parser "github.com/mssola/user_agent"
)

// Agent ...
type Agent struct {
	OS       string
	Browser  string
	Platform string
}

func (agent *Agent) GetOS() string {
	return agent.OS
}

func (agent *Agent) GetBrowser() string {
	return agent.Browser
}

func (agent *Agent) GetPlatform() string {
	return agent.Platform
}

// Parse ...
func (agent *Agent) Parse(ua string) {
	uaParsed := parser.New(ua)
	name, version := uaParsed.Browser()
	agent.Browser = fmt.Sprintf("%s %s", name, version)
	agent.OS = uaParsed.OS()
	agent.Platform = uaParsed.Platform()
}

func (agent *Agent) ParseUA(uaObject *uaparser.Parser, ua string) {
	uaParsed := uaObject.Parse(ua)
	agent.Browser = uaParsed.UserAgent.Family
	agent.OS = uaParsed.Device.Family
	agent.Platform = uaParsed.Os.Family
}
