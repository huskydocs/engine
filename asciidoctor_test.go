package engine

import (
    "testing"
    "os"
    "io/ioutil"
)

func TestInvoke(t *testing.T) {
    tempDocDir := createTempDir(t)
    tempDocFile := createTempDoc(t, tempDocDir)
    tempRenderedFile := createTempRendered(t, tempDocDir)
    
    opts := Opts()
    opts.files = []*os.File{tempDocFile}
    opts.baseDir = tempDocDir
    opts.outFile = tempRenderedFile
    
    out, err := Invoke(opts)
       
    if err != nil {
        t.Errorf("Unexpected exec error from Asciidoctor %s:\n%s", err, out)
    }
    removeTempDir(t, tempDocDir)
}

func TestInvokeWithError(t *testing.T) {
    tempDocDir := createTempDir(t)
    tempDocFile := createTempDoc(t, tempDocDir)
    tempRenderedFile := createTempRendered(t, tempDocDir)
    
    opts := Opts()
    opts.files = []*os.File{tempDocFile}
    opts.backend = "HTML5"
    opts.doctype = "book"
    opts.outFile = tempRenderedFile
    opts.noHeaderFooter = true
    opts.sectionNumbers = true
    opts.attributes = make(map[string]string)
    opts.attributes["att1"] = "value1"
    opts.baseDir = tempDocDir
    opts.destinationDir = tempDocDir
    out, err := Invoke(opts)
    if err == nil {
        t.Errorf("Expected an exec error from Asciidoctor:\n%s", out)
    }
    removeTempDir(t, tempDocDir)
}

func createTempDir(t *testing.T) (tempDir string) {
    tempDocDir, err := ioutil.TempDir("", "tempdocdir")
    if err != nil {
        t.Errorf("Unable to create temp doc dir %S", err)
    }
    return tempDocDir
}

func createTempDoc(t *testing.T, tempDocDir string) (tempDoc *os.File) {
    tempDocFile, err := ioutil.TempFile(tempDocDir, "tempdoc")
    if err != nil {
        t.Errorf("Unable to create temp doc %s", err)
    }
    return tempDocFile
}

func createTempRendered(t *testing.T, tempDocDir string) (tempRendered *os.File) {
    tempRenderedFile, err := ioutil.TempFile(tempDocDir, "temprendered.html")
    if err != nil {
        t.Errorf("Unable to create temp rendered file %s", err)
    }
    return tempRenderedFile
}

func removeTempDir(t *testing.T, tempDocDir string) {
    err := os.RemoveAll(tempDocDir)
    if err != nil {
        t.Errorf("Unexpected error trying to delete temp dir %s", err)
    }
}
