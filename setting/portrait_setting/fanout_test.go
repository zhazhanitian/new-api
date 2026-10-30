package portrait_setting

import (
	"reflect"
	"testing"

	"github.com/QuantumNous/new-api/constant"
)

func TestParseFanoutAliases(t *testing.T) {
	cases := []struct {
		raw  string
		want []string
	}{
		{"", []string{"volc"}},
		{"volc", []string{"volc"}},
		{"zj", []string{"volc", "zj"}},
		{"ZJ, byte_plus ,zj,,Bad Alias,volc", []string{"volc", "zj", "byte_plus"}},
		{"zj,volc", []string{"volc", "zj"}},
	}
	for _, tc := range cases {
		if got := parseFanoutAliases(tc.raw); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("parseFanoutAliases(%q) = %v, want %v", tc.raw, got, tc.want)
		}
	}
}

func TestFanoutAliasesReturnsCopy(t *testing.T) {
	t.Cleanup(func() { setFanoutAliasesForTest("") })
	setFanoutAliasesForTest("zj")
	got := FanoutAliases()
	got[0] = "changed"
	if FanoutAliases()[0] != "volc" {
		t.Fatal("FanoutAliases should return a copy")
	}
}

func TestAliasForChannelType(t *testing.T) {
	if got := AliasForChannelType(constant.ChannelTypeZhongJiVideo); got != "zj" {
		t.Errorf("zhongji = %q, want zj", got)
	}
	for _, ct := range []int{constant.ChannelTypeDoubaoVideo, constant.ChannelTypeVolcEngine, 0} {
		if got := AliasForChannelType(ct); got != DefaultAlias {
			t.Errorf("channel type %d = %q, want %q", ct, got, DefaultAlias)
		}
	}
}
