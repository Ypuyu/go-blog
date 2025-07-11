package crypto

import (
	"testing"
)

func TestGenAkas(t *testing.T) {
	tests := []struct {
		name   string
		wantAk string
		wantAs string
	}{
		{
			name:   "test",
			wantAk: "12345678901234567890123456789012",
			wantAs: "1234567890123456789012345678901234567890123456789012345678901234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAk, gotAs := GenAkas()
			if gotAk != tt.wantAk {
				t.Errorf("GenAkas() gotAk = %v, want %v", gotAk, tt.wantAk)
			}
			if gotAs != tt.wantAs {
				t.Errorf("GenAkas() gotAs = %v, want %v", gotAs, tt.wantAs)
			}
		})
	}
}
