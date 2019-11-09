package main

import (
	"testing"

	"github.com/blang/semver"
)

func Test_isPrereleaseSemver(t *testing.T) {
	tests := []struct {
		name      string
		tag       string
		expected  bool
		shouldErr bool
	}{
		{
			name:     "simple semver",
			tag:      "  v1.0.0",
			expected: false,
		},
		{
			name:     "simple semver with pre",
			tag:      "v1.0.0-alpha.1",
			expected: true,
		},
		{
			name:     "simple semver with pre and build",
			tag:      "v1.0.0-alpha.1+12342",
			expected: true,
		},
		{
			name:     "simple semver with build",
			tag:      "v1.0.0+12342",
			expected: false,
		},
		{
			name:      "not a semver",
			tag:       "v1.x.0+12342",
			shouldErr: true,
		},
		{
			name:      "not a semver at all",
			tag:       "v a b d1.x.0+12342",
			shouldErr: true,
		},
		{
			name:      "empty input",
			tag:       "",
			shouldErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := parseSemver(test.tag)
			if err == nil && test.shouldErr {
				t.Errorf("Expected %q to fail", test.tag)
			} else if err != nil && !test.shouldErr {
				t.Errorf("Expected %q not to fail", test.tag)
			}
		})
	}
}

func Test_toVersionMatcher(t *testing.T) {
	const custom = "2.0.1"
	tests := []struct {
		version  string
		isPatch  bool
		isMinor  bool
		isMajor  bool
		isCustom bool
	}{
		{
			version: "2.3.4-alpha.1+1234",
		},
		{
			version: "2.3.4-alpha.1",
		},
		{
			version: "2.3.4",
			isPatch: true,
		},
		{
			version: "2.3.0-alpha.2",
		},
		{
			version: "2.3.0",
			isPatch: true,
			isMinor: true,
		},
		{
			version:  "2.0.1-alpha3",
			isCustom: true,
		},
		{
			version:  "2.0.1",
			isPatch:  true,
			isCustom: true,
		},
		{
			version:  "2.0.0-alpha4",
			isCustom: true,
		},
		{
			version:  "2.0.0",
			isPatch:  true,
			isMinor:  true,
			isMajor:  true,
			isCustom: true,
		},
	}

	for _, tt := range tests {
		t.Run("any "+tt.version, func(t *testing.T) {
			matcher, err := toVersionMatcher(sinceAny)
			if err != nil {
				t.Fatal(err)
			}

			if matcher(semver.MustParse(tt.version)) != true {
				t.Errorf("'any' matcher should return true for %q", tt.version)
			}
		})
	}

	for _, tt := range tests {
		t.Run("patch "+tt.version, func(t *testing.T) {
			matcher, err := toVersionMatcher(sincePatch)
			if err != nil {
				t.Fatal(err)
			}

			if matcher(semver.MustParse(tt.version)) != tt.isPatch {
				t.Errorf("'patch' matcher should return true for %q", tt.version)
			}
		})
	}

	for _, tt := range tests {
		t.Run("minor "+tt.version, func(t *testing.T) {
			matcher, err := toVersionMatcher(sinceMinor)
			if err != nil {
				t.Fatal(err)
			}

			if matcher(semver.MustParse(tt.version)) != tt.isMinor {
				t.Errorf("'minor' matcher should return true for %q", tt.version)
			}
		})
	}

	for _, tt := range tests {
		t.Run("major "+tt.version, func(t *testing.T) {
			matcher, err := toVersionMatcher(sinceMajor)
			if err != nil {
				t.Fatal(err)
			}

			if matcher(semver.MustParse(tt.version)) != tt.isMajor {
				t.Errorf("'major' matcher should return true for %q", tt.version)
			}
		})
	}

	for _, tt := range tests {
		t.Run("since= "+custom+" "+tt.version, func(t *testing.T) {
			matcher, err := toVersionMatcher(custom)
			if err != nil {
				t.Fatal(err)
			}

			if matcher(semver.MustParse(tt.version)) != tt.isCustom {
				t.Errorf("since matcher should return true for %q", tt.version)
			}
		})
	}
}
