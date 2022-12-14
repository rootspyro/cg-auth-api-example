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
  "swagger": "2.0",
  "info": {
    "description": "cidesa group api for user authorization",
    "title": "cg-auth-api",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/authorize": {
      "post": {
        "description": "validate request",
        "tags": [
          "Auth"
        ],
        "operationId": "authorizeRequest",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/auth-request"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/auth-response"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/default"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "description": "get list of users",
        "tags": [
          "Users"
        ],
        "operationId": "getUsers",
        "parameters": [
          {
            "type": "integer",
            "default": 20,
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "string",
            "name": "username",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/users-response"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/default"
            }
          }
        }
      },
      "post": {
        "description": "create a new user",
        "tags": [
          "Users"
        ],
        "operationId": "createUser",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "email": {
                  "type": "string"
                },
                "firstname": {
                  "type": "string"
                },
                "lastname": {
                  "type": "string"
                },
                "passwordHash": {
                  "type": "string"
                },
                "username": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/user-response"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/default"
            }
          }
        }
      }
    },
    "/users/{username}": {
      "get": {
        "description": "get a single user by username",
        "tags": [
          "Users"
        ],
        "operationId": "getUserByUsername",
        "parameters": [
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/user-response"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/default"
            }
          }
        }
      },
      "delete": {
        "description": "delete a single user",
        "tags": [
          "Users"
        ],
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "description": "operation message",
                  "type": "string"
                },
                "status": {
                  "type": "string",
                  "default": "success"
                }
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/default"
            }
          }
        }
      },
      "patch": {
        "description": "update a single user",
        "tags": [
          "Users"
        ],
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "email": {
                  "type": "string"
                },
                "firstname": {
                  "type": "string"
                },
                "lastname": {
                  "type": "string"
                },
                "passwordHash": {
                  "type": "string"
                },
                "username": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/user-response"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/default"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "auth-request": {
      "type": "object",
      "properties": {
        "request": {
          "type": "object",
          "properties": {
            "method": {
              "description": "GET, POST, PATCH or DELETE...",
              "type": "string"
            },
            "path": {
              "type": "string"
            }
          }
        },
        "user": {
          "type": "object",
          "properties": {
            "username": {
              "type": "string"
            }
          }
        }
      }
    },
    "auth-response": {
      "type": "object",
      "properties": {
        "data": {
          "type": "object",
          "properties": {
            "result": {
              "type": "string"
            }
          }
        },
        "status": {
          "type": "string"
        }
      }
    },
    "default": {
      "type": "object",
      "properties": {
        "data": {
          "description": "message of what went wrong",
          "type": "string"
        },
        "status": {
          "description": "success or error",
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "firstname": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "lastname": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "user-response": {
      "type": "object",
      "properties": {
        "data": {
          "$ref": "#/definitions/user"
        },
        "status": {
          "type": "string"
        }
      }
    },
    "users-response": {
      "type": "object",
      "properties": {
        "data": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/user"
          }
        },
        "status": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "users endpoints",
      "name": "Users"
    },
    {
      "description": "request authorization endpoint",
      "name": "Auth"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "description": "cidesa group api for user authorization",
    "title": "cg-auth-api",
    "version": "1.0.0"
  },
  "basePath": "/api/v1",
  "paths": {
    "/authorize": {
      "post": {
        "description": "validate request",
        "tags": [
          "Auth"
        ],
        "operationId": "authorizeRequest",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/auth-request"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/auth-response"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/default"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "description": "get list of users",
        "tags": [
          "Users"
        ],
        "operationId": "getUsers",
        "parameters": [
          {
            "type": "integer",
            "default": 20,
            "name": "limit",
            "in": "query"
          },
          {
            "type": "integer",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "string",
            "name": "username",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/users-response"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/default"
            }
          }
        }
      },
      "post": {
        "description": "create a new user",
        "tags": [
          "Users"
        ],
        "operationId": "createUser",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "email": {
                  "type": "string"
                },
                "firstname": {
                  "type": "string"
                },
                "lastname": {
                  "type": "string"
                },
                "passwordHash": {
                  "type": "string"
                },
                "username": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/user-response"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/default"
            }
          }
        }
      }
    },
    "/users/{username}": {
      "get": {
        "description": "get a single user by username",
        "tags": [
          "Users"
        ],
        "operationId": "getUserByUsername",
        "parameters": [
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/user-response"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/default"
            }
          }
        }
      },
      "delete": {
        "description": "delete a single user",
        "tags": [
          "Users"
        ],
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "object",
              "properties": {
                "data": {
                  "description": "operation message",
                  "type": "string"
                },
                "status": {
                  "type": "string",
                  "default": "success"
                }
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/default"
            }
          }
        }
      },
      "patch": {
        "description": "update a single user",
        "tags": [
          "Users"
        ],
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "email": {
                  "type": "string"
                },
                "firstname": {
                  "type": "string"
                },
                "lastname": {
                  "type": "string"
                },
                "passwordHash": {
                  "type": "string"
                },
                "username": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/user-response"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/default"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "AuthRequestRequest": {
      "type": "object",
      "properties": {
        "method": {
          "description": "GET, POST, PATCH or DELETE...",
          "type": "string"
        },
        "path": {
          "type": "string"
        }
      }
    },
    "AuthRequestUser": {
      "type": "object",
      "properties": {
        "username": {
          "type": "string"
        }
      }
    },
    "AuthResponseData": {
      "type": "object",
      "properties": {
        "result": {
          "type": "string"
        }
      }
    },
    "auth-request": {
      "type": "object",
      "properties": {
        "request": {
          "type": "object",
          "properties": {
            "method": {
              "description": "GET, POST, PATCH or DELETE...",
              "type": "string"
            },
            "path": {
              "type": "string"
            }
          }
        },
        "user": {
          "type": "object",
          "properties": {
            "username": {
              "type": "string"
            }
          }
        }
      }
    },
    "auth-response": {
      "type": "object",
      "properties": {
        "data": {
          "type": "object",
          "properties": {
            "result": {
              "type": "string"
            }
          }
        },
        "status": {
          "type": "string"
        }
      }
    },
    "default": {
      "type": "object",
      "properties": {
        "data": {
          "description": "message of what went wrong",
          "type": "string"
        },
        "status": {
          "description": "success or error",
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "firstname": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "lastname": {
          "type": "string"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "user-response": {
      "type": "object",
      "properties": {
        "data": {
          "$ref": "#/definitions/user"
        },
        "status": {
          "type": "string"
        }
      }
    },
    "users-response": {
      "type": "object",
      "properties": {
        "data": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/user"
          }
        },
        "status": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "users endpoints",
      "name": "Users"
    },
    {
      "description": "request authorization endpoint",
      "name": "Auth"
    }
  ]
}`))
}
