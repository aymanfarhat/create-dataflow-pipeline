package main

import (
	"fmt"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/common-nighthawk/go-figure"
)


func getUseCaseTemplateByTitle(title string, useCaseTemplateOptions []UseCaseTemplate) UseCaseTemplate {
	for _, useCaseTemplate := range useCaseTemplateOptions {
		if useCaseTemplate.Name == title {
			return useCaseTemplate
		}
	}

	return UseCaseTemplate{}
}

func main() {
	defaultProject := getCurrentGcpProject()
	useCaseTemplateOptions := loadUseCaseTemplateOptions("python")
	optionTitles := make([]string, len(useCaseTemplateOptions))

	for i, useCaseTemplate := range useCaseTemplateOptions {
		optionTitles[i] = useCaseTemplate.Name
	}

	var qs = []*survey.Question{
		{
			Name: "Name",
			Prompt: &survey.Input{
				Message: "Pipeline name:",
				Help: "Choose a name for your pipeline",
			},
			Validate: survey.Required,
		},
		{
			Name: "Description",
			Prompt: &survey.Input{Message: "Describe your pipeline in a few words"},
			Validate: survey.Required,
		},
		{
			Name: "Project",
			Prompt: &survey.Input{Message: "GCP Project ID:", Default: defaultProject},
			Validate: survey.Required,
		},
		{
			Name: "Region",
			Prompt: &survey.Input{Message: "Preferred GCP Region:", Default: "europe-west1"},
			Validate: survey.Required,
		},
		{
			Name: "Language",
			Prompt: &survey.Select{
				Message: "Select a language:",
				Options: []string{"python", "java", "go"},
			},
			Validate: survey.Required,
		},
		{
			Name: "Path",
			Prompt: &survey.Input{
				Message: "Path to save the pipeline:",
				Suggest: func(toComplete string) []string {
					files, _ := filepath.Glob(toComplete + "*")
					return files
				},
			},
			Validate: survey.Required,
		},
		{
			Name: "ServiceAccount",
			Prompt: &survey.Confirm{Message: "Will this pipeline run via a service account?"},
			Validate: survey.Required,
		},
		{
			Name: "Subnetwork",
			Prompt: &survey.Confirm{Message: "Will this pipeline run in a VPC?"},
			Validate: survey.Required,
		},
		{
			Name: "FlexTemplate",
			Prompt: &survey.Confirm{Message: "Would you like a Dataflow Flex Template definition?"},
			Validate: survey.Required,
		},
		{
			Name: "ContainerImage",
			Prompt: &survey.Confirm{Message: "Would you like an SDK Container Image definition?"},
			Validate: survey.Required,
		},
		{
			Name: "TerraformInfra",
			Prompt: &survey.Confirm{
				Message: "Would you like to generate foundational Terraform infrastructure?",
				Default: true,
			},
			Validate: survey.Required,
		},
	}

	myFigure := figure.NewFigure("Dataflow", "", true)
	myFigure.Print()
	fmt.Println()
	fmt.Println("Pipeline Generator")
	fmt.Println("-----------------------")

	pipeline := PipelineInput{}
	err := survey.Ask(qs, &pipeline)

	if err != nil {
		panic(err)
	}

	// After acknowledging the programming language, ask for the use case template
	// Note this code needs to be refactored, potentially find a way to directly load
	// the use case template object base on choice
	var useCaseTemplateTile string
	var useCaseTemplate UseCaseTemplate
	err = survey.AskOne(
		&survey.Select{
			Message: "Select a starter template:",
			Options: optionTitles,
			}, 
		&useCaseTemplateTile)

	if err != nil {
		panic(err)
	}

	useCaseTemplate = getUseCaseTemplateByTitle(useCaseTemplateTile, useCaseTemplateOptions)

	targetDir := pipeline.Path + pipeline.Name // e.g. /tmp/myapp	
	baseTemplateDir := "templates/base-python" // TODO: Make this dynamic based on language choice

	renderPipeline(pipeline, useCaseTemplate, baseTemplateDir, targetDir)

	//if pipeline.TerraformInfra {
		baseInfratemplateDir := "templates/base-infra"
		renderInfra(pipeline, useCaseTemplate, baseInfratemplateDir, targetDir)
	//}

	fmt.Println("Done! proceed to", targetDir, "to see your pipeline.")
}