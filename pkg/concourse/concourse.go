package concourse

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"

	"github.com/concourse/atc"
	"github.com/concourse/fly/commands"
)

// Concourse is the concourse object
type Concourse struct {
	URL      string
	Username string
	Password string
}

// CreatePipeline creates and unpauses a pipeline
func (c *Concourse) CreatePipeline(credentials string, pipeline string, name string) (err error) {
	fly := &commands.Fly
	fly.Target = "concourse"
	fly.Login = commands.LoginCommand{
		ATCURL:      c.URL,
		Insecure:    true,
		Username:    c.Username,
		Password:    c.Password,
		OpenBrowser: false,
	}

	err = fly.Login.Execute([]string{})
	if err != nil {
		log.Printf("ERROR: %v", err)
		return err
	}

	credsfile, err := envToFile(credentials)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return err
	}

	pipelinefile, err := envToFile(pipeline)
	if err != nil {
		log.Printf("ERROR: %v", err)
		return err
	}

	pname := commands.PipelineFlag(name + "_" + RandomString(10))

	fly.SetPipeline = commands.SetPipelineCommand{
		SkipInteractive:  true,
		DisableAnsiColor: false,
		Pipeline:         pname,
		Config:           pipelinefile,
		VarsFrom:         []atc.PathFlag{credsfile},
	}

	err = fly.SetPipeline.Execute([]string{})
	if err != nil {
		log.Printf("ERROR: %v", err)
		return err
	}

	fly.UnpausePipeline = commands.UnpausePipelineCommand{
		Pipeline: pname,
	}

	err = fly.UnpausePipeline.Execute([]string{})
	if err != nil {
		log.Printf("ERROR: %v", err)
		return err
	}

	os.Remove(fmt.Sprintf("%s", credsfile))
	os.Remove(fmt.Sprintf("%s", pipelinefile))
	return nil

}

// credsToFile - writes the string contents to a random file and provides the path
func envToFile(s string) (file atc.PathFlag, err error) {
	f, err := ioutil.TempFile("./", "envvartofile_")
	if err != nil {
		return "", err
	}

	s2b := []byte(s)
	fname := f.Name()
	err = ioutil.WriteFile(fname, s2b, 0644)
	if err != nil {
		return "", err
	}

	log.Printf("Created %s", fname)
	return atc.PathFlag(fname), nil
}

func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	return string(bytes)
}
