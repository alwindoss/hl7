package hl7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	// FieldSeparator separates each segment into multiple fields and this symbol
	// is what is used to separate them
	FieldSeparator = "|"
)

// Segment holds each of the segment in the HL7 message
type Segment struct {
	Header string
}

// Field is a part of segment and one of more fields make up the segment
type Field struct {
}

// Message holds the HL7 message
type Message struct {
	Header Segment
}

// Parser parses the HL7 message and returns the object
func Parser() {
	file, err := os.Open("data/1.txt")
	if err != nil {
		log.Fatalf("unable to open file: %v", err)
	}
	// var buff make([]byte, 10)
	// buffer := bytes.NewReader(buff)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parseLine(line)
	}
}

func parseLine(line string) {
	segHead := fmt.Sprintf("%c%c%c", []rune(line)[0], []rune(line)[1], []rune(line)[2])
	parseSegment(segHead, line)
	// log.Println(segHead)
}

func parseSegment(segHead, segment string) {
	switch segHead {
	case "DG1":
		fmt.Println("Diagnosis Segment")
		fmt.Println("Content: ", segment)
		parseDiagnosisSegment(segment)
	case "PID":
		fmt.Println("Patient ID segment")
		fmt.Println("Content: ", segment)
		parsePIDSegment(segment)
	}
}

func parseDiagnosisSegment(segment string) {
	log.Println("parsing the Diagnosis Segment")
}

func parsePIDSegment(segment string) {
	fields := strings.Split(segment, FieldSeparator)
	for _, field := range fields {
		log.Println(field)
	}
}
