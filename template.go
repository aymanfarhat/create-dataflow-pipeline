package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"text/template"

	"golang.org/x/oauth2/google"
)

//go:embed templates
var templates embed.FS

func renderFile(template_path string, target_path string, templateDataInput TemplateDataInput) {
	t, err := template.ParseFiles(template_path)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(target_path)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, templateDataInput)
	if err != nil {
		panic(err)
	}
}

func getCurrentGcpProject() string {
	ctx := context.Background()
	credentials, err := google.FindDefaultCredentials(ctx)

	if err != nil {
		panic(err)
	}

	return credentials.ProjectID
}

func loadUseCaseTemplateOptions(language string) ([]UseCaseTemplate){
	var useCaseTemplateOptions []UseCaseTemplate
	fs.WalkDir(templates, "templates/use-cases", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.Contains(path, "config.json") {
			useCaseTemplate, err := loadUseCaseTemplate(path)

			if err != nil {
				panic(err)
			}

			if useCaseTemplate.Language == language {
				useCaseTemplateOptions = append(useCaseTemplateOptions, *useCaseTemplate)
			}
		}

		return nil
	})

	return useCaseTemplateOptions
}

func loadUseCaseTemplate(filePath string) (*UseCaseTemplate, error){
	var useCaseTemplate UseCaseTemplate

	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &useCaseTemplate)

	if err != nil {
		return nil, err
	}

	return &useCaseTemplate, nil
}

func renderInfra(pipeline PipelineInput, useCaseTemplate UseCaseTemplate, baseTemplateDir string, targetRoot string) {
	templateDataInput := TemplateDataInput{
		Pipeline: pipeline,
		UseCaseTemplate: useCaseTemplate,
	}

	targetInfraDir := fmt.Sprintf("%s/%s_infra", targetRoot, pipeline.Name)
	err := os.MkdirAll(targetInfraDir, 0777)

	if err != nil {
		panic(err)
	}

	fs.WalkDir(templates, baseTemplateDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Set the target path to the same path in the target directory
		targetPath := strings.Replace(path, baseTemplateDir, targetInfraDir, 1)
		// Remove .tmpl extension
		targetPath = strings.Replace(targetPath, ".tmpl", "", 1)

		if d.IsDir() {
			err := os.MkdirAll(targetPath, 0777)
			if err != nil {
				panic(err)
			}

			return nil
		}

		renderFile(path, targetPath, templateDataInput)

		return nil	
	})
}

func renderPipeline(pipeline PipelineInput, useCaseTemplate UseCaseTemplate, baseTemplateDir string, targetRoot string) {
	templateDataInput := TemplateDataInput{
		Pipeline: pipeline,
		UseCaseTemplate: useCaseTemplate,
	}

	targetPipelineDir := fmt.Sprintf("%s/%s_pipeline", targetRoot, pipeline.Name)
	err := os.MkdirAll(targetPipelineDir, 0777)

	if err != nil {
		panic(err)
	}

	// Apply template files and dirs from base template
	fs.WalkDir(templates, baseTemplateDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Set the target path to the same path in the target directory
		targetPath := strings.Replace(path, baseTemplateDir, targetPipelineDir, 1)
		// Remove .tmpl extension
		targetPath = strings.Replace(targetPath, ".tmpl", "", 1)

		if d.IsDir() {
			err := os.MkdirAll(targetPath, 0777)
			if err != nil {
				panic(err)
			}

			return nil
		}

		renderFile(path, targetPath, templateDataInput)

		return nil	
	})

	// Apply use case code based on use case template
	module_name := fmt.Sprintf("%s_app", pipeline.Name)
	module_dir := fmt.Sprintf("%s/%s", targetPipelineDir, module_name)
	err = os.MkdirAll(module_dir, 0777)

	if err != nil {
		panic(err)
	}

	_, err = os.Create(module_dir + "/__init__.py")

	if err != nil {
		panic(err)
	}

	filecontent, err := templates.ReadFile(fmt.Sprintf("templates/%s", useCaseTemplate.CodePath))

	if err != nil {
		panic(err)
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s", module_dir, "app.py"), filecontent, 0777)

	if err != nil {
		panic(err)
	}
}