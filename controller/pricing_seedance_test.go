package controller

import "testing"

func TestSeedanceDisplayTiersUsesOfficialName(t *testing.T) {
	tiers := seedanceDisplayTiers("doubao-seedance-2-0-fast-260128", nil)
	if len(tiers) != 2 {
		t.Fatalf("tiers = %#v", tiers)
	}
	if tiers[0].Label != "不含视频 · 480P/720P" || tiers[1].Ratio <= 0 || tiers[1].Ratio >= 1 {
		t.Fatalf("unexpected official tiers: %#v", tiers)
	}
}

func TestSeedanceDisplayTiersFollowsChannelMapping(t *testing.T) {
	mappings := []map[string]string{{
		"doubao-seedance-2-0-fast-zj": "doubao-seedance-2-0-fast-260128",
	}}
	alias := seedanceDisplayTiers("doubao-seedance-2-0-fast-zj", mappings)
	official := seedanceDisplayTiers("doubao-seedance-2-0-fast-260128", nil)
	if len(alias) != len(official) {
		t.Fatalf("alias tiers = %#v, official = %#v", alias, official)
	}
	for i := range official {
		if alias[i].Label != official[i].Label || alias[i].Ratio != official[i].Ratio {
			t.Fatalf("tier %d alias=%#v official=%#v", i, alias[i], official[i])
		}
	}
}

func TestSeedanceDisplayTiersWithoutMappingStayEmpty(t *testing.T) {
	if tiers := seedanceDisplayTiers("doubao-seedance-2-0-fast-zj", nil); len(tiers) != 0 {
		t.Fatalf("unmapped alias should not invent tiers: %#v", tiers)
	}
}
