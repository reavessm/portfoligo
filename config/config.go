/**
 * File: config.go
 * Written by: Stephen M. Reaves
 * Created on: Thu, 03 Mar 2022
 */

package config

type Social struct {
	Type        string `yaml:"media"`
	Username    string `yaml:"username"`
	Link        string `yaml:"link"`
	Description string `yaml:"description"`
}

type Contact struct {
	Method        string `yaml:"method"`
	Ref           string `yaml:"address"`
	IncludeMailTo bool   `yaml:"includeMailTo"`
}

type Project struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Link        string `yaml:"link"`
}

type Resume struct {
	Name     string `yaml:"name"`
	FileName string `yaml:"link"`
}

type Certification struct {
	Name string `yaml:"name"`
	Link string `yaml:"link"`
}

type Event struct {
	Description string `yaml:"description"`
	Year        int    `yaml:"year"`
}

type Additional struct {
	Enabled     bool   `yaml:"enabled"`
	Description string `yaml:"description"`
}

type Config struct {
	Name           string          `yaml:"name"`
	NickName       string          `yaml:"nickname"`
	Title          string          `yaml:"title"`
	Description    string          `yaml:"description"`
	Socials        []Social        `yaml:"social"`
	Contacts       []Contact       `yaml:"contact"`
	Projects       []Project       `yaml:"projects"`
	Resumes        []Resume        `yaml:"resume"`
	Certifications []Certification `yaml:"certs"`
	History        []Event         `yaml:"history"`
	Additional     Additional      `yaml:"additional"`
}
