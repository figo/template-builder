package config

// TemplateVariables define template variables
type TemplateVariables struct {
	CompName        string
	CompNameCapital string // capitalize the first char
	MirrorURL       string
	UpstreamURL     string
	BranchName      string
	CompVersion     string
	GitRepoOrg      string
}
