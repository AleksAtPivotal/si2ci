package concourse

import (
	"testing"

	"github.com/concourse/atc"
	"github.com/concourse/fly/commands"
)

// Tests to make sure go vendor is not broken
func TestSetPipeline(t *testing.T) {
	t.Logf("If this test fails, upstream fork of vendor/github.com/concourse may require changes")
	fly := &commands.Fly
	pipelinefile := atc.PathFlag("pipeline.yml")
	credsfile := atc.PathFlag("creds.yml")
	var p commands.PipelineFlag = "test"

	fly.SetPipeline = commands.SetPipelineCommand{
		SkipInteractive:  true,
		DisableAnsiColor: false,
		Pipeline:         p,
		Config:           pipelinefile,
		VarsFrom:         []atc.PathFlag{credsfile},
	}
}
