package main

import "testing"

func TestSDES(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		key     string
		decrypt bool
		want    string
	}{
		{
			name:    "Encryption 1",
			input:   "00000000",
			key:     "0000000000",
			decrypt: false,
			want:    "11110000",
		},
		{
			name:    "Encryption 2",
			input:   "01011011",
			key:     "1111111111",
			decrypt: false,
			want:    "01000000",
		},
		{
			name:    "Decryption 1",
			input:   "10000001",
			key:     "1000000001",
			decrypt: true,
			want:    "11101110",
		},
		{
			name:    "Decryption 2",
			input:   "10101001",
			key:     "1101110100",
			decrypt: true,
			want:    "11100010",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, key, err := formatArguments(tt.input, tt.key)
			if err != nil {
				t.Fatalf("Error formatting arguments: %v", err)
			}

			s := &sdes{input: input, key: key, decrypt: tt.decrypt}
			result, err := s.process()
			if err != nil {
				t.Fatalf("Error processing: %v", err)
			}

			got := ""
			for _, v := range result {
				got += string(v + '0')
			}

			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}
