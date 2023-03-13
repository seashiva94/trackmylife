package tasktracker

import (
    "fmt"
    "io/ioutil"
    "strings"
)

type Parser struct {

}

type ParseOutput struct {
    Projects *[]Project
}

type Project struct {
    ProjectName string
    ProjectDescription string
    Sections *[]Section
}

type Section struct {
    Items *[]Item
}

type Item struct {
    Description string
    Status string
    ETA string
    Tags []string
    Comment string
}

func GenerateNewParser() *Parser {
    return &Parser{}
}

// read the input file and parse contents of the file
func (p *Parser) ParseTaskFile(inputFileName string) (*ParseOutput, error) {
    body, err := ioutil.ReadFile(inputFileName)
    if err != nil {
        return nil, err
    }

    // use regex grouping to split projects and sections 
    parseOutput, err := p.parseProject(strings.Split(string(body), "\n# "))
    if err != nil {
        return nil, err
    }
    return parseOutput, nil
}

// 
func (p *Parser) parseProject(lines []string) (*ParseOutput, error) {
    fmt.Println("parsing")
    return nil, nil
}

//regex matches sections and tasks: ## \w+\n([- \w+\n]+)
//regex matches projects sections and tasks (#+ [\w ]+\n([- \w+\n]+)*) -> project line and section line are separate matches
//above doesn't include special charaters
// same as above but hadles some special characters  (#+ [\w ']+\n([- \w+\n,><?\[\]]+)*) -> generalize for all special characters
