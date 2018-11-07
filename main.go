package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alekssaul/si2ci/pkg/concourse"
)

const (
	appVersion = "0.2"
)

func main() {
	http.HandleFunc("/", handleRoot)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("Recieved Payload: %v", r)
	log.Printf("Referrer: %s", r.Header.Get("Referer"))

	c := concourse.Concourse{
		URL:      os.Getenv("CONCOURSE_TARGET"),
		Username: os.Getenv("CONCOURSE_USERNAME"),
		Password: os.Getenv("CONCOURSE_PASSWORD"),
	}

	creds := os.Getenv("CREDENTIALS_YML")
	pipeline := os.Getenv("PIPELINE_YML")

	// add Spring Initializer content into credential file based on request parameters
	creds = addToCreds(creds, "SI2CI_NAME", r.URL.Query().Get("name"))
	creds = addToCreds(creds, "SI2CI_LANGUAGE", r.URL.Query().Get("language"))
	creds = addToCreds(creds, "SI2CI_BOOTVERSION", r.URL.Query().Get("bootVersion"))
	creds = addToCreds(creds, "SI2CI_BASEDIR", r.URL.Query().Get("baseDir"))
	creds = addToCreds(creds, "SI2CI_GROUPID", r.URL.Query().Get("groupId"))
	creds = addToCreds(creds, "SI2CI_ARTIFACTID", r.URL.Query().Get("artifactId"))
	creds = addToCreds(creds, "SI2CI_DESCRIPTION", r.URL.Query().Get("description"))
	creds = addToCreds(creds, "SI2CI_PACKAGENAME", r.URL.Query().Get("packageName"))
	creds = addToCreds(creds, "SI2CI_PACKAGING", r.URL.Query().Get("packaging"))
	creds = addToCreds(creds, "SI2CI_JAVAVERSION", r.URL.Query().Get("javaVersion"))

	err := c.CreatePipeline(creds, pipeline, r.URL.Query().Get("name"))
	if err != nil {
		log.Printf("ERROR: %s", err)
	}

	// Redirect user back to Spring Initializr
	http.Redirect(w, r, r.Header.Get("Referer"), 301)
}

// addToCreds adds key value pairs for the credentials YAML.
// it is intended to be use to merge in Spring Initialzr vars into creds.yml
func addToCreds(in string, key string, value string) (out string) {
	out = in + "\n" + key + ": " + value
	return out
}
