// +build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
	"text/template"

	cfschema "github.com/hashicorp/aws-cloudformation-resource-schema-sdk-go"
	"github.com/iancoleman/strcase"
	"github.com/mitchellh/cli"
)

var (
	cfTypeSchemaFile = flag.String("cfschema", "", "CloudFormation resource type schema file; required")
	packageName      = flag.String("package", "", "override package name for generated code")
	tfResourceType   = flag.String("resource", "", "Terraform resource type; required")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "\tmain.go [flags] -resource <TF-resource-type> -cfschema <CF-type-schema-file> <generated-file>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 || *tfResourceType == "" || *cfTypeSchemaFile == "" {
		flag.Usage()
		os.Exit(2)
	}

	destinationPackage := os.Getenv("GOPACKAGE")
	if *packageName != "" {
		destinationPackage = *packageName
	}

	filename := args[0]

	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	generator := NewGenerator(ui, *tfResourceType, *cfTypeSchemaFile)

	if err := generator.Generate(destinationPackage, filename); err != nil {
		ui.Error(fmt.Sprintf("error generating Terraform resource: %s", err))
		os.Exit(1)
	}
}

type Generator struct {
	cfTypeSchemaFile string
	tfResourceType   string
	ui               cli.Ui
}

func NewGenerator(ui cli.Ui, tfResourceType, cfTypeSchemaFile string) *Generator {
	return &Generator{
		cfTypeSchemaFile: cfTypeSchemaFile,
		tfResourceType:   tfResourceType,
		ui: &cli.BasicUi{
			Reader:      os.Stdin,
			Writer:      os.Stdout,
			ErrorWriter: os.Stderr,
		},
	}
}

func (g *Generator) Infof(format string, a ...interface{}) {
	g.ui.Info(fmt.Sprintf(format, a...))
}

// Generate generates the resource's definition in the specified file.
func (g *Generator) Generate(packageName, filename string) error {
	g.Infof("generating Terraform schema for %q from %q into %q", g.tfResourceType, g.cfTypeSchemaFile, filename)

	resource, err := NewResourcePath(g.tfResourceType, g.cfTypeSchemaFile)

	if err != nil {
		return fmt.Errorf("error reading CloudFormation resource schema for %s: %w", g.tfResourceType, err)
	}

	templateData := TemplateData{
		FunctionName:     resource.SourceCodeNamePrefix,
		NamePrefix:       resource.SourceCodeNamePrefix,
		PackageName:      packageName,
		ResourceTypeName: g.tfResourceType,
	}

	tmpl, err := template.New("function").Parse(templateBody)

	if err != nil {
		return fmt.Errorf("error parsing function template: %w", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, templateData)

	if err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	generatedFileContents, err := format.Source(buffer.Bytes())

	if err != nil {
		return fmt.Errorf("error formatting generated file: %w", err)
	}

	f, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("error creating file (%s): %w", filename, err)
	}

	defer f.Close()

	_, err = f.Write(generatedFileContents)

	if err != nil {
		return fmt.Errorf("error writing to file (%s): %w", filename, err)
	}

	return nil
}

var templateBody = `
// Code generated by generators/resource/main.go; DO NOT EDIT.

package {{ .PackageName }}

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResource("{{ .ResourceTypeName }}", {{ .FunctionName }}())
}

// {{ .FunctionName }} returns the Terraform {{ .ResourceTypeName }} resource.
func {{ .FunctionName }}() *schema.Resource {
	return &schema.Resource{
		CreateContext: {{.NamePrefix}}Create,
		ReadContext:   {{.NamePrefix}}Read,
		UpdateContext: {{.NamePrefix}}Update,
		DeleteContext: {{.NamePrefix}}Delete,

		// Schema: {{.NamePrefix}}Schema,
	}
}

func {{.NamePrefix}}Create(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func {{.NamePrefix}}Read(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func {{.NamePrefix}}Update(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func {{.NamePrefix}}Delete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
`

type TemplateData struct {
	FunctionName     string
	NamePrefix       string
	PackageName      string
	ResourceTypeName string
}

type Resource struct {
	CfResource           *cfschema.Resource
	SourceCodeNamePrefix string
	TfType               string
}

func NewResourcePath(resourceType, cfTypeSchemaFile string) (*Resource, error) {
	resourceSchema, err := cfschema.NewResourceJsonSchemaPath(cfTypeSchemaFile)

	if err != nil {
		return nil, fmt.Errorf("error reading CloudFormation Resource Type Schema: %w", err)
	}

	resource, err := resourceSchema.Resource()

	if err != nil {
		return nil, fmt.Errorf("error parsing CloudFormation Resource Type Schema: %w", err)
	}

	if err := resource.Expand(); err != nil {
		return nil, fmt.Errorf("error expanding JSON Pointer references: %w", err)
	}

	return &Resource{
		CfResource:           resource,
		SourceCodeNamePrefix: "resource" + strcase.ToCamel(resourceType),
		TfType:               resourceType,
	}, nil
}
