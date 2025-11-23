package exec

import (
	"flag"
	"testing"

	"github.com/urfave/cli"
)

func TestMapOldToExecCommand(t *testing.T) {
	testCases := []struct {
		name  string
		flags []string
		debug bool
		trace bool
	}{
		{
			name:  "no flags",
			flags: []string{},
			debug: false,
			trace: false,
		},
		{
			name:  "debug flag",
			flags: []string{"--debug"},
			debug: true,
			trace: false,
		},
		{
			name:  "trace flag",
			flags: []string{"--trace"},
			debug: false,
			trace: true,
		},
		{
			name:  "debug and trace flags",
			flags: []string{"--debug", "--trace"},
			debug: true,
			trace: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			set := flag.NewFlagSet("test", 0)
			set.Bool("debug", false, "")
			set.Bool("trace", false, "")
			set.Parse(tc.flags)

			ctx := cli.NewContext(nil, set, nil)
			cmd := mapOldToExecCommand(ctx)

			if cmd.Debug != tc.debug {
				t.Errorf("expected Debug to be %v, got %v", tc.debug, cmd.Debug)
			}
			if cmd.Trace != tc.trace {
				t.Errorf("expected Trace to be %v, got %v", tc.trace, cmd.Trace)
			}
		})
	}
}
