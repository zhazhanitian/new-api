package portrait_setting

import (
	"os"
	"testing"
)

func TestResolveVolcUsesLiveSettings(t *testing.T) {
	origAK, origSK, origProject, origRegion := VolcPortraitAccessKeyId, VolcPortraitSecretAccessKey, VolcPortraitProjectName, VolcPortraitRegion
	t.Cleanup(func() {
		VolcPortraitAccessKeyId = origAK
		VolcPortraitSecretAccessKey = origSK
		VolcPortraitProjectName = origProject
		VolcPortraitRegion = origRegion
	})

	VolcPortraitAccessKeyId = ""
	VolcPortraitSecretAccessKey = ""
	if _, err := Resolve(""); err == nil {
		t.Fatal("empty volc credentials should fail")
	}
	if _, err := Resolve("volc"); err == nil {
		t.Fatal("volc without credentials should fail")
	}

	VolcPortraitAccessKeyId = "ak"
	VolcPortraitSecretAccessKey = "sk"
	VolcPortraitProjectName = "proj"
	VolcPortraitRegion = ""
	cred, err := Resolve(" VOLC ")
	if err != nil {
		t.Fatal(err)
	}
	if cred.Alias != DefaultAlias || cred.AK != "ak" || cred.SK != "sk" || cred.Project != "proj" || cred.Region != "cn-beijing" {
		t.Fatalf("unexpected volc credential: %+v", cred)
	}
}

func TestResolveRegistry(t *testing.T) {
	original := append([]string{}, os.Environ()...)
	t.Cleanup(func() { reloadRegistry(original) })

	reloadRegistry([]string{
		"PORTRAIT_HC_ACCESS_KEY_ID=hc-ak",
		"PORTRAIT_HC_SECRET_ACCESS_KEY=hc-sk",
		"PORTRAIT_HC_PROJECT_NAME=hc-proj",
		"PORTRAIT_BYTE_PLUS_ACCESS_KEY_ID=bp-ak",
		"PORTRAIT_BYTE_PLUS_SECRET_ACCESS_KEY=bp-sk",
		"PORTRAIT_BYTE_PLUS_PROJECT_NAME=bp-proj",
		"PORTRAIT_BYTE_PLUS_REGION=ap-southeast-1",
		"PORTRAIT_NOSK_ACCESS_KEY_ID=only-ak",
		"VOLC_PORTRAIT_ACCESS_KEY_ID=should-not-register",
		"VOLC_PORTRAIT_SECRET_ACCESS_KEY=should-not-register",
		"PORTRAIT_VOLC_ACCESS_KEY_ID=ignored",
		"PORTRAIT_VOLC_SECRET_ACCESS_KEY=ignored",
	})

	cred, err := Resolve("HC")
	if err != nil {
		t.Fatal(err)
	}
	if cred.Alias != "hc" || cred.AK != "hc-ak" || cred.SK != "hc-sk" || cred.Project != "hc-proj" || cred.Region != "cn-beijing" {
		t.Fatalf("unexpected hc credential: %+v", cred)
	}

	cred, err = Resolve("byte-plus")
	if err != nil {
		t.Fatal(err)
	}
	if cred.Alias != "byte-plus" || cred.AK != "bp-ak" || cred.Region != "ap-southeast-1" {
		t.Fatalf("unexpected hyphen credential: %+v", cred)
	}
	cred, err = Resolve("byte_plus")
	if err != nil {
		t.Fatal(err)
	}
	if cred.AK != "bp-ak" || cred.Alias != "byte_plus" {
		t.Fatalf("unexpected underscore credential: %+v", cred)
	}

	if _, err = Resolve("nosk"); err == nil {
		t.Fatal("missing secret should be unconfigured")
	}
	if _, err = Resolve("missing"); err == nil {
		t.Fatal("unknown alias should fail")
	}
	if _, err = Resolve("bad alias"); err == nil {
		t.Fatal("illegal alias should fail")
	}
	if _, err = Resolve("../etc"); err == nil {
		t.Fatal("path-like alias should fail")
	}
}
