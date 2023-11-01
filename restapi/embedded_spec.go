// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is server about authors",
    "title": "Book authors",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/v1",
  "paths": {
    "/authors": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Authors"
        ],
        "summary": "get all authors",
        "operationId": "getAllAuthors",
        "responses": {
          "200": {
            "description": "successful response",
            "schema": {
              "$ref": "#/definitions/Authors"
            }
          },
          "default": {
            "description": "Все нестандартное",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Authors"
        ],
        "summary": "Add a new author to the store",
        "operationId": "addAuthor",
        "parameters": [
          {
            "description": "Author that needed to be in the list",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Author"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "properties": {
                "id": {
                  "type": "integer"
                }
              }
            }
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/authors/{authorId}": {
      "get": {
        "description": "Returns a single author",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Authors"
        ],
        "summary": "Find author by ID",
        "operationId": "getAuthorById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of author to return",
            "name": "authorId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Author"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Author not found"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Authors"
        ],
        "summary": "Deletes an author",
        "operationId": "deleteAuthor",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID author to delete",
            "name": "authorId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Author not found"
          }
        }
      }
    }
  },
  "definitions": {
    "Author": {
      "type": "object",
      "required": [
        "name",
        "surname"
      ],
      "properties": {
        "author_id": {
          "type": "integer",
          "example": 1
        },
        "name": {
          "type": "string",
          "example": "Dmitrii"
        },
        "surname": {
          "type": "string",
          "example": "Urievich"
        }
      }
    },
    "Authors": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Author"
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is server about authors",
    "title": "Book authors",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/v1",
  "paths": {
    "/authors": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Authors"
        ],
        "summary": "get all authors",
        "operationId": "getAllAuthors",
        "responses": {
          "200": {
            "description": "successful response",
            "schema": {
              "$ref": "#/definitions/Authors"
            }
          },
          "default": {
            "description": "Все нестандартное",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Authors"
        ],
        "summary": "Add a new author to the store",
        "operationId": "addAuthor",
        "parameters": [
          {
            "description": "Author that needed to be in the list",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Author"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "properties": {
                "id": {
                  "type": "integer"
                }
              }
            }
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/authors/{authorId}": {
      "get": {
        "description": "Returns a single author",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Authors"
        ],
        "summary": "Find author by ID",
        "operationId": "getAuthorById",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID of author to return",
            "name": "authorId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Author"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Author not found"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Authors"
        ],
        "summary": "Deletes an author",
        "operationId": "deleteAuthor",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "ID author to delete",
            "name": "authorId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Author not found"
          }
        }
      }
    }
  },
  "definitions": {
    "Author": {
      "type": "object",
      "required": [
        "name",
        "surname"
      ],
      "properties": {
        "author_id": {
          "type": "integer",
          "example": 1
        },
        "name": {
          "type": "string",
          "example": "Dmitrii"
        },
        "surname": {
          "type": "string",
          "example": "Urievich"
        }
      }
    },
    "Authors": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Author"
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
}