package main

import (
	"fmt"

	"github.com/cxpsemea/Cx1ClientGo"
)

type ApplicationCRUD struct {
	Name        string            `yaml:"Name"`
	Test        string            `yaml:"Test"`
	Groups      []string          `yaml:"Groups"`
	Criticality uint              `yaml:"Criticality"`
	Rules       []ApplicationRule `yaml:"Rules"`
	Tags        []Tag             `yaml:"Tags"`
	FailTest    bool              `yaml:"FailTest"`
	TestResult  bool
	Application *Cx1ClientGo.Application
}

func (o *ApplicationCRUD) String() string {
	return o.Name
}

type ApplicationRule struct {
	Type  string `yaml:"Type"`
	Value string `yaml:"Value"`
}

func (o *ApplicationRule) String() string {
	return fmt.Sprintf("%v: %v", o.Type, o.Value)
}

type CxQLCRUD struct {
	//QueryID       uint64 `yaml:"ID"`
	QueryLanguage string    `yaml:"Language"`
	QueryGroup    string    `yaml:"Group"`
	QueryName     string    `yaml:"Name"`
	Test          string    `yaml:"Test"`
	Source        string    `yaml:"Source"`
	Scope         CxQLScope `yaml:"Scope"`
	Severity      string    `yaml:"Severity"`
	FailTest      bool      `yaml:"FailTest"`
	Compile       bool      `yaml:"Compile"`
	TestResult    bool
	Query         *Cx1ClientGo.AuditQuery
}

func (o *CxQLCRUD) String() string {
	//if o.QueryName != "" {
	return fmt.Sprintf("%v: %v -> %v -> %v", o.Scope, o.QueryLanguage, o.QueryGroup, o.QueryName)
	/*} else {
		return fmt.Sprintf("QueryID#%d", o.QueryID)
	} // */
}

type CxQLScope struct {
	Corp        bool   `yaml:"Tenant"`
	Project     string `yaml:"Project"`
	Application string `yaml:"Application"`
}

type GroupCRUD struct {
	Name        string `yaml:"Name"`
	Test        string `yaml:"Test"`
	Parent      string `yaml:"Parent"`
	ClientRoles []struct {
		Client string   `yaml:"Client"`
		Roles  []string `yaml:"Roles"`
	} `yaml:"ClientRoles"`
	FailTest   bool `yaml:"FailTest"`
	TestResult bool
	Group      *Cx1ClientGo.Group
}

func (o *GroupCRUD) String() string {
	return o.Name
}

type ProjectCRUD struct {
	Name        string   `yaml:"Name"`
	Test        string   `yaml:"Test"`
	Groups      []string `yaml:"Groups"`
	Application string   `yaml:"Application"`
	Tags        []Tag    `yaml:"Tags"`
	FailTest    bool     `yaml:"FailTest"`
	Project     *Cx1ClientGo.Project
}

func (o *ProjectCRUD) String() string {
	return o.Name
}

type PresetCRUD struct {
	Name        string `yaml:"Name"`
	Description string `yaml:"Description"`
	Test        string `yaml:"Test"`
	Queries     []struct {
		QueryID       uint64 `yaml:"ID"`
		QueryLanguage string `yaml:"Language"`
		QueryGroup    string `yaml:"Group"`
		QueryName     string `yaml:"Name"`
	} `yaml:"Queries"`
	FailTest bool `yaml:"FailTest"`
	Preset   *Cx1ClientGo.Preset
}

func (o *PresetCRUD) String() string {
	return o.Name
}

type QueryCRUD struct {
	QueryID       uint64 `yaml:"ID"`
	QueryLanguage string `yaml:"Language"`
	QueryGroup    string `yaml:"Group"`
	QueryName     string `yaml:"Name"`
	Test          string `yaml:"Test"`
	FailTest      bool   `yaml:"FailTest"`
	Query         *Cx1ClientGo.Query
}

func (o *QueryCRUD) String() string {
	if o.QueryName != "" {
		return fmt.Sprintf("%v -> %v -> %v", o.QueryLanguage, o.QueryGroup, o.QueryName)
	} else {
		return fmt.Sprintf("QueryID#%d", o.QueryID)
	}
}

type ResultCRUD struct {
	QueryID       uint64 `yaml:"ID"`
	QueryLanguage string `yaml:"Language"`
	QueryGroup    string `yaml:"Group"`
	QueryName     string `yaml:"Name"`
	ProjectName   string `yaml:"Project"`
	Number        uint64 `yaml:"FindingNumber"`
	SimilarityID  int64  `yaml:"SimilarityID"`
	ResultHash    string `yaml:"ResultHash"`
	State         string `yaml:"State"`
	Severity      string `yaml:"Severity"`
	Comment       string `yaml:"Comment"`
	Test          string `yaml:"Test"`
	FailTest      bool   `yaml:"FailTest"`
	Result        *Cx1ClientGo.ScanResult
	Project       *Cx1ClientGo.Project
}

func (o *ResultCRUD) String() string {
	if o.QueryName != "" {
		return fmt.Sprintf("%v: %v -> %v -> %v finding #%d", o.ProjectName, o.QueryLanguage, o.QueryGroup, o.QueryName, o.Number)
	}
	if o.SimilarityID != 0 {
		return fmt.Sprintf("%v: finding with simID %d", o.ProjectName, o.SimilarityID)
	}
	if o.ResultHash != "" {
		return fmt.Sprintf("%v: finding with hash %v", o.ProjectName, o.ResultHash)
	}
	if o.QueryID != 0 {
		return fmt.Sprintf("%v: QueryID #%d finding #%d", o.ProjectName, o.QueryID, o.Number)
	}
	return fmt.Sprintf("%v: finding #%d", o.ProjectName, o.Number)
}

type RoleCRUD struct {
	Name        string   `yaml:"Name"`
	Test        string   `yaml:"Test"`
	Permissions []string `yaml:"Permissions"`
	FailTest    bool     `yaml:"FailTest"`
	Role        *Cx1ClientGo.Role
}

func (o *RoleCRUD) String() string {
	return o.Name
}

type ScanCRUD struct {
	Test        string `yaml:"Test"`
	Project     string `yaml:"Project"`
	Branch      string `yaml:"Branch"`
	Repository  string `yaml:"Repository"`
	Engine      string `yaml:"Engine"`
	Incremental bool   `yaml:"Incremental"`
	WaitForEnd  bool   `yaml:"WaitForEnd"`
	ZipFile     string `yaml:"ZipFile"`
	Preset      string `yaml:"Preset"`
	FailTest    bool   `yaml:"FailTest"`
	Scan        *Cx1ClientGo.Scan
}

func (o *ScanCRUD) String() string {
	return fmt.Sprintf("%v: %v branch %v", o.Project, o.Repository, o.Branch)
}

type UserCRUD struct {
	Name      string   `yaml:"Name"`
	Email     string   `yaml:"Email"`
	Test      string   `yaml:"Test"`
	FirstName string   `yaml:"FirstName"`
	LastName  string   `yaml:"LastName"`
	Groups    []string `yaml:"Groups"`
	Roles     []string `yaml:"Roles"`
	FailTest  bool     `yaml:"FailTest"`
	User      *Cx1ClientGo.User
}

func (o *UserCRUD) String() string {
	return fmt.Sprintf("%v (%v)", o.Name, o.Email)
}

type Tag struct {
	Key   string `yaml:"Key"`
	Value string `yaml:"Value"`
}

type TestSet struct {
	Name         string            `yaml:"Name"`
	File         string            `yaml:"File"`
	Groups       []GroupCRUD       `yaml:"Groups"`
	Users        []UserCRUD        `yaml:"Users"`
	Applications []ApplicationCRUD `yaml:"Applications"`
	Projects     []ProjectCRUD     `yaml:"Projects"`
	Queries      []CxQLCRUD        `yaml:"Queries"`
	Presets      []PresetCRUD      `yaml:"Presets"`
	Roles        []RoleCRUD        `yaml:"Roles"`
	Scans        []ScanCRUD        `yaml:"Scans"`
	Results      []ResultCRUD      `yaml:"Results"`
}

type TestConfig struct {
	Cx1URL   string    `yaml:"Cx1URL"`
	IAMURL   string    `yaml:"IAMURL"`
	Tenant   string    `yaml:"Tenant"`
	ProxyURL string    `yaml:"ProxyURL"`
	Tests    []TestSet `yaml:"Tests"`
}

type TestResult struct {
	FailTest   bool
	Result     bool
	CRUD       string
	Duration   float64
	Name       string
	Id         int
	TestObject string
	Reason     string
}
