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
    command := make([]string, 30)
    
    if opts.backend != "" {
        command = append(command, "-b", opts.backend)
    }
    
    if opts.doctype != "" {
        command = append(command, "-d", opts.doctype)
    }
    
    if opts.outFile != "" {
        command = append(command, "-o", opts.outFile)
    }
    
    if opts.noHeaderFooter  {
        command = append(command, "-s")
    }
    if opts.sectionNumbers {
        command = append(command, "-n")
    }
    for key, value := range opts.attributes {
        command = append(command, "-a", key + "=" + value)
    }
    if opts.baseDir != "" {
        command = append(command, "-B", opts.baseDir)
    }
    if opts.destinationDir != "" {
        command = append(command, "-D", opts.destinationDir)
    }
    fmt.Println("Executing command, %v", command)
    out, err := exec.Command("asciidoctor", "--help").Output()
    if err != nil {
        return err
    }
    fmt.Printf("%s", out)
    return nil
}
