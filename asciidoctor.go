package engine

import (
    "fmt"
    "os"
    "os/exec"
//    "strings"
)

type AsciidoctorOpts struct {
    files []*os.File
    backend string
    doctype string
    outFile *os.File
    noHeaderFooter bool
    sectionNumbers bool
    attributes map[string]string
    baseDir string
    destinationDir string
}

func Opts() *AsciidoctorOpts {
    return new(AsciidoctorOpts)
}

func Invoke(opts *AsciidoctorOpts) (out []byte, err error) {
    command := make([]string, 0)
    
    if opts.backend != "" {
        command = append(command, "-b", opts.backend)
    }
    
    if opts.doctype != "" {
        command = append(command, "-d", opts.doctype)
    }
    
    if opts.outFile != nil && opts.outFile.Name() != "" {
        command = append(command, "-o", opts.outFile.Name())
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
    
    if opts.files != nil { 
        for _,file := range opts.files {
            command = append(command, file.Name())
        }
    }
    fmt.Printf("Executing command %v \n", command);
    out, err = exec.Command("asciidoctor", command...).CombinedOutput()
    if err != nil {
        return out, err
    }
    fmt.Printf("%s", out)
    return out, nil
}
