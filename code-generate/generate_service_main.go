package main

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
)

// Entry point to generate services
func generateServices(objects []RamlType, resources []ResourceService) {
	typeMapping := map[string]*RamlType{}

	for i := range objects {
		item := objects[i]
		typeMapping[item.Name] = &item
	}

	for _, resourceService := range resources {
		f := jen.NewFile("commercetools")
		f.HeaderComment("Automatically generated, do not edit")

		resource := typeMapping[resourceService.ResourceType]

		if resource == nil {
			continue
		}
		// Skip this because it has a weird RAML definition
		if resource.CodeName == "CustomObject" {
			continue
		}

		// TODO: ugly hack
		resourceService.ResourceType = CreateCodeName(resourceService.ResourceType)
		resourceService.ResourceDraftType = CreateCodeName(resourceService.ResourceDraftType)
		resourceService.ResourceQueryType = CreateCodeName(resourceService.ResourceQueryType)

		urlPathName := fmt.Sprintf("%sURLPath", resource.CodeName)
		c := jen.Comment(fmt.Sprintf("%s is the commercetools API path.", urlPathName)).Line()
		c.Const().Id(urlPathName).Op("=").Lit(resourceService.BasePath).Line()
		f.Add(c)

		if resourceService.HasCreate {
			funcName := fmt.Sprintf("%sCreate", resource.CodeName)
			createCode := generateServiceCreate(funcName, resourceService, urlPathName)
			f.Add(createCode)
		}

		if resourceService.HasList {
			funcName := fmt.Sprintf("%sQuery", resource.CodeName)
			queryCode := generateServiceQuery(funcName, resourceService, urlPathName)
			f.Add(queryCode)
		}

		httpMethodsCode := generateResourceHTTPMethodsCode(resource, resourceService, urlPathName)
		for _, httpMethodCode := range httpMethodsCode {
			f.Add(httpMethodCode)
		}

		filename := generateFilename(resource.Package)
		err := f.Save("commercetools/service_" + filename)
		if err != nil {
			panic(err)
		}
	}
}

func generateResourceHTTPMethodsCode(resource *RamlType, resourceService ResourceService, urlPathName string) (httpMethodsCode []*jen.Statement) {
	httpMethodsCode = make([]*jen.Statement, 0)
	for _, resourceMethod := range resourceService.ResourceMethods {
		for _, httpMethod := range resourceMethod.HTTPMethods {
			switch httpMethod.HTTPMethod {
			case "get":
				funcName := fmt.Sprintf("%sGet%s", resource.CodeName, strings.Title(resourceMethod.MethodName))
				c := generateServiceGet(funcName, resourceService, resourceMethod, httpMethod)
				httpMethodsCode = append(httpMethodsCode, c)
			case "post":
				funcName := fmt.Sprintf("%sUpdate%s", resource.CodeName, strings.Title(resourceMethod.MethodName))
				c := generateServiceUpdate(funcName, resource.CodeName, resourceService, resourceMethod, httpMethod)
				httpMethodsCode = append(httpMethodsCode, c)
			case "delete":
				funcName := fmt.Sprintf("%sDelete%s", resource.CodeName, strings.Title(resourceMethod.MethodName))
				c := generateServiceDelete(funcName, resourceService, resourceMethod, httpMethod)
				httpMethodsCode = append(httpMethodsCode, c)
			default:
				continue

			}

		}
	}
	return httpMethodsCode
}