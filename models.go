package main

/** Represents the user input object */
type PipelineInput struct {
	Name string
	Description string
	Project string
	Region string
	Language string
	Path string
	ServiceAccount bool
	Subnetwork bool
	FlexTemplate bool
	ContainerImage bool
	UseCaseTemplate string
	TerraformInfra bool
}


/** Represents a use case template definition object */
type UseCaseTemplate struct {
	Name string `json:"name"`
    Description string `json:"description"`
    Streaming bool `json:"streaming"`
    Language string `json:"language"`
	CodePath string `json:"code_path"`
    Parameters []Parameter `json:"parameters"`
}

/** Represents a use case parameter in the template definition object */
type Parameter struct {
    Name string `json:"name"`
    Label string `json:"label"`
    HelpText string `json:"help_text"`
}

/** Represents the object that will be passed to the template renderer */
type TemplateDataInput struct {
	Pipeline PipelineInput
	UseCaseTemplate UseCaseTemplate
}