package main

import "testing"

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
