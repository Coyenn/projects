package main

import (
	"io/ioutil"
	"log"
	"os/exec"
)

type project struct {
	title string
	description string
}

func (i project) Title() string       { return i.title }
func (i project) Description() string { return i.description }
func (i project) FilterValue() string { return i.title }

type projectsFinder struct {
	projects  []project
}

func (r *projectsFinder) findProjects() {
	files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
		if _, err := ioutil.ReadDir(f.Name() + "/.git"); err == nil {
			description := "No description"

			if _, err := ioutil.ReadFile(f.Name() + "/README.md"); err == nil {
				// read first line of README.md
				stdout, stderr := exec.Command("head", "-n", "1", f.Name() + "/README.md").CombinedOutput()

				if stderr != nil {
					log.Fatal(stderr)
				}

				// Replace # with space
				description = string(stdout)
				description = description[1:len(description)-1]
			}

			r.projects = append(r.projects, project{
				title: f.Name(),
				description: description,
			})
		}
    }
}