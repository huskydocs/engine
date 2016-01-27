package engine

import (
    "fmt"
    "os/exec"
)

type AsciidoctorOpts struct {
    files []string
    backend string
    doctype string
    outFile string
    noHeaderFooter bool
    sectionNumbers bool
    attributes map[string]string
    baseDir string
    destinationDir string
}

func Opts() *AsciidoctorOpts {
    return new(AsciidoctorOpts)
}

func Invoke(opts *AsciidoctorOpts) (err error) {
    fmt.Println("Executing help")
    out, err := exec.Command("asciidoctor", "--help").Output()
    if err != nil {
        return err
    }
    fmt.Printf("%s", out)
    return nil
}
