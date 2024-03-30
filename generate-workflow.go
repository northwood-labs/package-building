package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/goware/prefixer"
	"github.com/lithammer/dedent"
	log "github.com/sirupsen/logrus"
)

var (
	fCronExpr  string
	fOwnerRepo string
	fPackage   string
	fTemplate  string

	dockerfile *template.Template

	doNotEditWarning = strings.TrimSpace(dedent.Dedent(`
DO NOT EDIT THIS FILE!

1. Edit the *.gotmpl.yml files instead.
2. Run 'go run generate-workflow.go -p {package} -t {template}'.
`))
)

func init() {
	flag.StringVar(&fCronExpr, "c", "", "The cron expression for a scheduled build.")
	flag.StringVar(&fOwnerRepo, "r", "", "The GitHub owner/repo to lookup.")
	flag.StringVar(&fPackage, "p", "", "The name of the package to create stubs for.")
	flag.StringVar(&fTemplate, "t", "", "The name of the template file to use.")

	flag.Parse()
}

func main() {
	dockerfile = template.Must(template.ParseFiles(fTemplate))

	dockerfileFile, err := os.Create(fmt.Sprintf(".github/workflows/build-%s.yml", fPackage))
	if err != nil {
		log.Fatal(err)
	}

	// Do some magic to convert the DO NOT EDIT warning into a valid Dockerfile comment.
	buf := bytes.NewBufferString(doNotEditWarning)
	prefixReader := prefixer.New(buf, "# ")

	commentedBuf, err := io.ReadAll(prefixReader)
	if err != nil {
		log.Fatal(err)
	}

	doNotEditStringValue := strings.TrimSpace(dedent.Dedent(`
    ################################################################################
    %s
    ################################################################################
    `))

	// Take the template, execute it using variables we pass-in from the
	// JSON config, and write the result to the new Dockerfile.
	err = dockerfile.Execute(dockerfileFile, map[string]interface{}{
		"CronExpr":  fCronExpr,
		"OwnerRepo": fOwnerRepo,
		"Package":   fPackage,
		"DoNotEdit": fmt.Sprintf(doNotEditStringValue, string(commentedBuf)),
	})
	if err != nil {
		log.Fatal(err)
	}
}