package validation_test

import (
	"testing"

	"code.vegaprotocol.io/vega/oracles/validation"
	"github.com/stretchr/testify/require"
)

func TestCheckForInternalOracle(t *testing.T) {
	type args struct {
		data map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Should return an error if there is any data that contains reserved prefix",
			args: args{
				data: map[string]string{
					"aaaa":                      "aaaa",
					"bbbb":                      "bbbb",
					"cccc":                      "cccc",
					"vegaprotocol.builtin.dddd": "dddd",
				},
			},
			wantErr: true,
		},
		{
			name: "Should pass validation if none of the data contains a reserved prefix",
			args: args{
				data: map[string]string{
					"aaaa": "aaaa",
					"bbbb": "bbbb",
					"cccc": "cccc",
					"dddd": "dddd",
				},
			},
			wantErr: false,
		},
		{
			name: "Should pass validation if reserved prefix is contained in key, but key doesn't start with the prefix",
			args: args{
				data: map[string]string{
					"aaaa":                      "aaaa",
					"bbbb":                      "bbbb",
					"cccc":                      "cccc",
					"dddd.vegaprotocol.builtin": "dddd",
				},
			},
			wantErr: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(tt *testing.T) {
			if err := validation.CheckForInternalOracle(tc.args.data); tc.wantErr {
				require.Error(tt, err)
			} else {
				require.NoError(tt, err)
			}
		})
	}
}