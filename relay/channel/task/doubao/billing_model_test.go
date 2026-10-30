package doubao

import "testing"

import relaycommon "github.com/QuantumNous/new-api/relay/common"

func TestBillingModelNameKeepsOfficialID(t *testing.T) {
	info := &relaycommon.RelayInfo{
		OriginModelName: "doubao-seedance-2-0-fast-260128",
		ChannelMeta: &relaycommon.ChannelMeta{
			IsModelMapped:     true,
			UpstreamModelName: "doubao-seedance-2-0-260128",
		},
	}
	if got := billingModelName(info); got != "doubao-seedance-2-0-fast-260128" {
		t.Fatalf("official id changed to %s", got)
	}
}

func TestBillingModelNameFallsBackToMappedUpstream(t *testing.T) {
	info := &relaycommon.RelayInfo{
		OriginModelName: "doubao-seedance-2-0-fast-zj",
		ChannelMeta: &relaycommon.ChannelMeta{
			IsModelMapped:     true,
			UpstreamModelName: "doubao-seedance-2-0-fast-260128",
		},
	}
	got := billingModelName(info)
	if got != "doubao-seedance-2-0-fast-260128" {
		t.Fatalf("billing model = %s", got)
	}
	if _, ok := GetVideoInputRatio(got); !ok {
		t.Fatal("upstream id is not in the video input ratio table")
	}
}

func TestBillingModelNameIgnoresUnmappedRename(t *testing.T) {
	info := &relaycommon.RelayInfo{
		OriginModelName: "doubao-seedance-2-0-fast-zj",
		ChannelMeta: &relaycommon.ChannelMeta{
			UpstreamModelName: "doubao-seedance-2-0-fast-260128",
		},
	}
	if got := billingModelName(info); got != "doubao-seedance-2-0-fast-zj" {
		t.Fatalf("unmapped rename changed to %s", got)
	}
}
