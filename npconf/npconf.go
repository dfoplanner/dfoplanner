package npconf

import "strings"

// Neople uses a ini accent for their configuration
// This package implements a simple parser for those format

type Npconf struct {
	sectionList []string
	sections map[string]*Section
}

func (n *Npconf)Load(data interface{}) (error) {
	lineSlice := strings.Split(data, "\n")
	var sectionName string
	var rawBody string
	for idx, line := range lineSlice {
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line,"#") || strings.HasPrefix(line,";") || strings.HasPrefix(line,"//") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			if strings.HasPrefix(line, "[/") {

			} else {
				sectionName = line[1 : len(line)-1]
				rawBody = ""
			}
		}

	}
	return nil
}