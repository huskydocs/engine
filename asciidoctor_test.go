package engine

import "testing"

func TestInvoke(t *testing.T) {
    opts := Opts()
    opts.files = []string{"file1", "file2"}
    opts.backend = "HTML5"
    opts.doctype = "book"
    opts.outFile = "targetFile"
    opts.noHeaderFooter = true
    opts.sectionNumbers = true
    opts.attributes = make(map[string]string)
    opts.attributes["att1"] = "value1"
    opts.attributes["att2"] = "value2"
    opts.baseDir = "baseDir"
    opts.destinationDir = "destinationDir"
    err := Invoke(opts)
    if err != nil {
        t.Errorf("Unexpected Exec error %s", err)
    }
}
