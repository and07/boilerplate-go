package templater

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"
)

// Config ...
type Config struct {
	ServiceName    string
	User           string
	IgnorePatterns []string
}

// Service ...
type Service struct {
	cfg Config
}

// Replace ...
func (s *Service) Replace(tmpl string) (string, error) {
	templ, err := template.New("test").Parse(tmpl)
	if err != nil {
		return "", err
	}
	res := &bytes.Buffer{}

	err = templ.Execute(res, s.cfg)
	return res.String(), err
}

// BuildService ...
func (s *Service) BuildService(srcPath, dstPath string) error {
	finfo, err := os.Stat(srcPath)
	if err != nil {
		return err
	}

	// construct new path
	srcName := path.Base(srcPath)
	renderedSrcName, err := s.Replace(srcName)
	if err != nil {
		return err
	}
	renderedDestPath := path.Join(dstPath, renderedSrcName)

	// proccess file
	switch finfo.IsDir() {

	// if a directory
	case true:
		// read all files in it
		files, err := ioutil.ReadDir(srcPath)
		if err != nil {
			return err
		}
		err = mkdirIfNoExist(renderedDestPath)
		if err != nil {
			return err
		}
		for _, f := range files {
			err = s.BuildService(path.Join(srcPath, f.Name()), renderedDestPath)
			if err != nil {
				return err
			}
		}

	// if a file
	case false:
		fdata, err := ioutil.ReadFile(srcPath)
		if err != nil {
			return err
		}

		renderedFdata := string(fdata)
		if !s.cfg.inIgnore(srcName) {
			renderedFdata, err = s.Replace(string(fdata))
			if err != nil {
				return err
			}
		} else {
			log.Printf("template ignored in %s", srcName)
		}

		err = ioutil.WriteFile(renderedDestPath, []byte(renderedFdata), basePrmission)
		if err != nil {
			return err
		}
	}
	return nil
}

// New ...
func New(cfg Config) *Service {
	return &Service{cfg: cfg}
}
