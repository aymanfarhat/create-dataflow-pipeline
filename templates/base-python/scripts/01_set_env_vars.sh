# Base variables for the pipeline
export PROJECT_ID="{{.Pipeline.Project}}" # Change to your project ID
export REGION="{{.Pipeline.Region}}" # Change to your region e.g. us-central1
export STAGING_LOCATION="gs://[Your-bucket-name]/staging" # Change to your bucket for storing staging files
export TEMP_LOCATION="gs://[Your-bucket-name]/temp" # Change to your bucket for storing temp files
{{ if .Pipeline.Subnetwork }}
# Subnetwork config
export SERVICE_ACCOUNT="[Your-service-account]" # Service account for running the pipeline (optional)
export SUBNETWORK="regions/$REGION/subnetworks/[Your-subnetwork]" # Change to your subnetwork e.g. "regions/us-central1/subnetworks/default (optional)
{{ end }}
{{ if .Pipeline.FlexTemplate }}
# Flex template variables
export FLEX_TEMPLATE_LOCATION="gs://[Your flex template bucket]/templates/pipeline.json"
export FLEX_BASE_IMAGE="gcr.io/dataflow-templates-base/python310-template-launcher-base:flex_templates_base_image_release_20230126_RC00"
{{ end }}
{{ if .Pipeline.ContainerImage }}
# SDK container image variables
export REPO_NAME="[Your-gcr-repo-name]"
export SDK_CONTAINER_IMAGE="$REGION-docker.pkg.dev/$PROJECT_ID/$REPO_NAME/kambr-pipeline:latest"
{{ end }}
# Pipeline specific variables
{{ range $arg := .UseCaseTemplate.Parameters }}
export {{$arg.Name}}="" # {{$arg.HelpText}}
{{ end }}

echo "Please make sure to run this script as follows: source 01_set_env_vars.sh"
