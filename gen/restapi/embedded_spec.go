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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Task Management System",
    "title": "Task Scheduler",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/",
  "paths": {
    "/task": {
      "get": {
        "description": "return tasks",
        "operationId": "listTasks",
        "responses": {
          "200": {
            "description": "tasks response",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Task"
              }
            }
          },
          "404": {
            "description": "tasks not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "put": {
        "operationId": "updateTask",
        "parameters": [
          {
            "description": "task details",
            "name": "task",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Task"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "task updated",
            "schema": {
              "$ref": "#/definitions/Task"
            }
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "post": {
        "operationId": "addTask",
        "parameters": [
          {
            "description": "task's details",
            "name": "task",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Task"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "task added",
            "schema": {
              "$ref": "#/definitions/Task"
            }
          },
          "409": {
            "description": "task already exists"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/task/{ID}": {
      "delete": {
        "operationId": "deleteTask",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the task",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "task deleted"
          },
          "404": {
            "description": "task not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "Task": {
      "type": "object",
      "properties": {
        "Assignee": {
          "type": "string"
        },
        "CreatedAt": {
          "type": "string",
          "format": "date-time"
        },
        "Deadline": {
          "type": "string",
          "format": "date-time"
        },
        "Description": {
          "type": "string"
        },
        "ID": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Priority": {
          "type": "integer"
        },
        "Status": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Task Management System",
    "title": "Task Scheduler",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/",
  "paths": {
    "/task": {
      "get": {
        "description": "return tasks",
        "operationId": "listTasks",
        "responses": {
          "200": {
            "description": "tasks response",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Task"
              }
            }
          },
          "404": {
            "description": "tasks not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "put": {
        "operationId": "updateTask",
        "parameters": [
          {
            "description": "task details",
            "name": "task",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Task"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "task updated",
            "schema": {
              "$ref": "#/definitions/Task"
            }
          },
          "500": {
            "description": "internal server error"
          }
        }
      },
      "post": {
        "operationId": "addTask",
        "parameters": [
          {
            "description": "task's details",
            "name": "task",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Task"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "task added",
            "schema": {
              "$ref": "#/definitions/Task"
            }
          },
          "409": {
            "description": "task already exists"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    },
    "/task/{ID}": {
      "delete": {
        "operationId": "deleteTask",
        "parameters": [
          {
            "type": "string",
            "description": "UUID of the task",
            "name": "ID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "task deleted"
          },
          "404": {
            "description": "task not found"
          },
          "500": {
            "description": "internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "Task": {
      "type": "object",
      "properties": {
        "Assignee": {
          "type": "string"
        },
        "CreatedAt": {
          "type": "string",
          "format": "date-time"
        },
        "Deadline": {
          "type": "string",
          "format": "date-time"
        },
        "Description": {
          "type": "string"
        },
        "ID": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        },
        "Priority": {
          "type": "integer"
        },
        "Status": {
          "type": "string"
        }
      }
    }
  }
}`))
}