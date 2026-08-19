package relay

import (
	"testing"

	relayconstant "github.com/QuantumNous/new-api/relay/constant"
)

func TestMusicFetchByIDResponseBuilderRegistered(t *testing.T) {
	builder, ok := fetchRespBuilders[relayconstant.RelayModeMusicFetchByID]
	if !ok || builder == nil {
		t.Fatal("RelayModeMusicFetchByID response builder is not registered")
	}
}
