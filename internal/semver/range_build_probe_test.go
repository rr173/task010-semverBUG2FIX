package semver

import "testing"

// Probe: range expressions with build metadata in the version spec
// should parse correctly. Build metadata after "+" must be stripped
// before prerelease parsing occurs.

func TestRangeBuildMetadataInSpec(t *testing.T) {
	v, err := Parse("1.0.0-rc.1")
	if err != nil {
		t.Fatal(err)
	}
	// Build metadata in range spec should be ignored per SemVer 2.0.0
	ok, err := SatisfiesRange(v, ">=1.0.0-rc.1+build")
	if err != nil {
		t.Fatalf("SatisfiesRange returned error: %v", err)
	}
	if !ok {
		t.Errorf("1.0.0-rc.1 should satisfy >=1.0.0-rc.1+build, got false")
	}
}

func TestRangeCaretWithBuildMeta(t *testing.T) {
	v, err := Parse("1.2.3-beta")
	if err != nil {
		t.Fatal(err)
	}
	ok, err := SatisfiesRange(v, "^1.2.3-beta+meta")
	if err != nil {
		t.Fatalf("SatisfiesRange returned error: %v", err)
	}
	if !ok {
		t.Errorf("1.2.3-beta should satisfy ^1.2.3-beta+meta, got false")
	}
}
