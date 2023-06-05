package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func New() *openapi3.T {
	swagger := &openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "Global Hitss",
			Description: "REST APIs used for interacting with the ToDo Service.",
			Version:     "1.0.0",
			License: &openapi3.License{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
			Contact: &openapi3.Contact{
				Name:  "Jorge Luis Alonso",
				Email: "alonso.jorgeluis.01@gmail.com",
				URL:   "https://github.com/Alonso9696",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				URL:         "{Scheme}://{Host}:{Port}/api/v1",
				Description: "URL endpoints",
				Variables: map[string]*openapi3.ServerVariable{
					"Scheme": {
						Enum:    []string{"http", "https"},
						Default: "http",
					},
					"Host": {
						Enum:    []string{"localhost", "192.168.1.5"},
						Default: "localhost",
					},
					"Port": {
						Enum:    []string{"4000", "5000"},
						Default: "4000",
					},
				},
			},
		},
		Tags: openapi3.Tags{
			&openapi3.Tag{
				Name:        "Users",
				Description: "Endpoints to manage users.",
				ExternalDocs: &openapi3.ExternalDocs{
					Description: "external docs description",
					URL:         "https://swagger.io/specification/",
				},
			},
			&openapi3.Tag{
				Name:        "Login",
				Description: "Endpoint to login users.",
				ExternalDocs: &openapi3.ExternalDocs{
					Description: "external docs description",
					URL:         "https://swagger.io/specification/",
				},
			},
		},
		ExternalDocs: &openapi3.ExternalDocs{
			Description: "Find out more about OpenAPI Specification.",
			URL:         "https://swagger.io/specification/",
		},
		Components: &openapi3.Components{
			SecuritySchemes: openapi3.SecuritySchemes{
				"bearerAuth": &openapi3.SecuritySchemeRef{
					Value: openapi3.NewJWTSecurityScheme(),
				},
			},
		},
		Security: openapi3.SecurityRequirements{
			openapi3.SecurityRequirement{
				"bearerAuth": []string{},
			},
		},
	}

	swagger.Components.Schemas = openapi3.Schemas{
		"Login": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("email", openapi3.NewStringSchema().WithDefault("alonso12@gmail.com")).
				WithProperty("password", openapi3.NewStringSchema().WithDefault("123456"))),
		"User": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("id", openapi3.NewInt64Schema()).
				WithProperty("name", openapi3.NewStringSchema()).
				WithProperty("email", openapi3.NewStringSchema()).
				WithProperty("password", openapi3.NewStringSchema()).
				WithProperty("tag", openapi3.NewStringSchema()).
				WithProperty("active", openapi3.NewInt32Schema())),
	}

	swagger.Components.RequestBodies = openapi3.RequestBodies{
		"LoginUsersRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for logging a user.").
				WithRequired(true).
				WithJSONSchemaRef(&openapi3.SchemaRef{
					Ref: "#/components/schemas/Login",
				}),
		},
		"CreateUsersRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for creating a user.").
				WithRequired(true).
				WithJSONSchema(openapi3.NewObjectSchema().
					WithProperty("name", openapi3.NewStringSchema()).
					WithProperty("email", openapi3.NewStringSchema()).
					WithProperty("password", openapi3.NewStringSchema()).
					WithProperty("tag", openapi3.NewStringSchema())),
		},
		"UpdateUsersRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for updating a user.").
				WithRequired(true).
				WithJSONSchema(openapi3.NewObjectSchema().
					WithProperty("name", openapi3.NewStringSchema()).
					WithProperty("email", openapi3.NewStringSchema()).
					WithProperty("password", openapi3.NewStringSchema()).
					WithProperty("tag", openapi3.NewStringSchema()).
					WithProperty("active", openapi3.NewInt32Schema())),
		},
	}

	swagger.Components.Responses = openapi3.Responses{
		"ClientErrorResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response when client errors happen.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("code", openapi3.NewInt32Schema().WithDefault(400)).
					WithProperty("error", openapi3.NewStringSchema()))),
		},
		"ServerErrorResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response when server errors happen.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("code", openapi3.NewInt32Schema().WithDefault(500)).
					WithProperty("error", openapi3.NewStringSchema()))),
		},
		"LoginUsersResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after logging users.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("code", openapi3.NewInt32Schema().WithDefault(200)).
					WithProperty("token", openapi3.NewStringSchema()))),
		},
		"GetAllUsersResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after getting all users.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("code", openapi3.NewInt32Schema().WithDefault(200)).
					WithProperty("users", openapi3.NewArraySchema().WithItems(&openapi3.Schema{
						Items: &openapi3.SchemaRef{
							Ref: "#/components/schemas/User",
						},
					})))),
		},
		"CreateUsersResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after creating users.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("code", openapi3.NewInt32Schema().WithDefault(200)).
					WithPropertyRef("user", &openapi3.SchemaRef{Ref: "#/components/schemas/User"}))),
		},
		"UpdateUsersResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after updating users.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("code", openapi3.NewInt32Schema().WithDefault(200)).
					WithPropertyRef("user", &openapi3.SchemaRef{Ref: "#/components/schemas/User"}))),
		},
		"DeleteUsersResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after deleting users.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("code", openapi3.NewInt32Schema().WithDefault(200)).
					WithProperty("user", openapi3.NewStringSchema()))),
		},
	}

	swagger.Paths = openapi3.Paths{
		"/signin": &openapi3.PathItem{
			Post: &openapi3.Operation{
				Summary:     "Sing in user.",
				Description: "Login user specifying email and password to get token.",
				Tags:        []string{"Login"},
				OperationID: "SignInUser",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/LoginUsersRequest",
				},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Ref: "#/components/responses/LoginUsersResponse",
					},
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ClientErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ServerErrorResponse",
					},
				},
			},
		},
		"/users": &openapi3.PathItem{
			Get: &openapi3.Operation{
				Summary:     "Get all users.",
				Description: "Gets all user rows.",
				Tags:        []string{"Users"},
				OperationID: "GetAllUser",
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Ref: "#/components/responses/GetAllUsersResponse",
					},
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ClientErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ServerErrorResponse",
					},
				},
			},
			Post: &openapi3.Operation{
				Summary:     "Create user.",
				Description: "Register a user with the requested fields.",
				Tags:        []string{"Users"},
				OperationID: "CreateUser",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/CreateUsersRequest",
				},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Ref: "#/components/responses/CreateUsersResponse",
					},
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ClientErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ServerErrorResponse",
					},
				},
			},
		},
		"/users/{id}": &openapi3.PathItem{
			Put: &openapi3.Operation{
				Summary:     "Update user.",
				Description: "Updates a user with the requested fields.",
				Tags:        []string{"Users"},
				OperationID: "UpdateUser",
				Parameters: []*openapi3.ParameterRef{
					{
						Value: openapi3.NewPathParameter("id").WithSchema(openapi3.NewInt64Schema()),
					},
				},
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/UpdateUsersRequest",
				},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Ref: "#/components/responses/UpdateUsersResponse",
					},
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ClientErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ServerErrorResponse",
					},
				},
			},
			Delete: &openapi3.Operation{
				Summary:     "Delete user.",
				Description: "Delete a user by id.",
				Tags:        []string{"Users"},
				OperationID: "DeleteUser",
				Parameters: []*openapi3.ParameterRef{
					{
						Value: openapi3.NewPathParameter("id").WithSchema(openapi3.NewInt64Schema()),
					},
				},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Ref: "#/components/responses/DeleteUsersResponse",
					},
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ClientErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ServerErrorResponse",
					},
				},
			},
		},
	}
	return swagger
}
