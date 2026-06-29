package joyagent

const joyAgentHost = "agentrs.jd.com"

func buildAuthHeader(apiKey string) string {
	return "Bearer " + apiKey
}
