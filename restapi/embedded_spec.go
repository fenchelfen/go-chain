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
    "description": "GoChain API",
    "title": "GoChain",
    "contact": {
      "email": "m.lyamets@innopolis.ru"
    },
    "license": {
      "name": "WTFPL"
    },
    "version": "1.0.0"
  },
  "host": "localhost:3000",
  "basePath": "/v1",
  "paths": {
    "/chain": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Retrieve the entire chain of blocks",
        "operationId": "getChain",
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/Chain"
            }
          }
        }
      }
    },
    "/mine": {
      "get": {
        "summary": "Mine transactions and announce to peers",
        "operationId": "Mine",
        "responses": {
          "200": {
            "description": "success"
          }
        }
      }
    },
    "/new_transaction": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Commit new transaction",
        "operationId": "addTransaction",
        "parameters": [
          {
            "name": "Transaction",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Transaction"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "transaction commited"
          },
          "400": {
            "description": "invalid transaction"
          }
        }
      }
    },
    "/register_node": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Register with this node",
        "operationId": "registerNode",
        "parameters": [
          {
            "name": "Peer",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Peer"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/Chain"
            }
          }
        }
      }
    },
    "/register_with": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Register with a specific node",
        "operationId": "registerWithNode",
        "parameters": [
          {
            "name": "Peer",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Peer"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/Chain"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Block": {
      "type": "object",
      "required": [
        "index",
        "transactions",
        "proofOfWork"
      ],
      "properties": {
        "index": {
          "type": "integer"
        },
        "prevBlockHash": {
          "description": "sha256 + nonce digest",
          "type": "string"
        },
        "proofOfWork": {
          "description": "sha256 digest",
          "type": "string"
        },
        "transactions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Transaction"
          }
        }
      }
    },
    "Chain": {
      "required": [
        "length",
        "chain",
        "peers"
      ],
      "properties": {
        "chain": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Block"
          }
        },
        "length": {
          "type": "integer"
        },
        "peers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Peer"
          }
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string",
          "example": "failure"
        }
      }
    },
    "Peer": {
      "type": "object",
      "properties": {
        "node_address": {
          "type": "string",
          "format": "url"
        }
      }
    },
    "Transaction": {
      "type": "object",
      "required": [
        "author",
        "content"
      ],
      "properties": {
        "author": {
          "type": "string",
          "format": "uuid",
          "example": "d290f1ee-6c54-4b01-90e6-d701748f0851"
        },
        "content": {
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
    "description": "GoChain API",
    "title": "GoChain",
    "contact": {
      "email": "m.lyamets@innopolis.ru"
    },
    "license": {
      "name": "WTFPL"
    },
    "version": "1.0.0"
  },
  "host": "localhost:3000",
  "basePath": "/v1",
  "paths": {
    "/chain": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Retrieve the entire chain of blocks",
        "operationId": "getChain",
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/Chain"
            }
          }
        }
      }
    },
    "/mine": {
      "get": {
        "summary": "Mine transactions and announce to peers",
        "operationId": "Mine",
        "responses": {
          "200": {
            "description": "success"
          }
        }
      }
    },
    "/new_transaction": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Commit new transaction",
        "operationId": "addTransaction",
        "parameters": [
          {
            "name": "Transaction",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Transaction"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "transaction commited"
          },
          "400": {
            "description": "invalid transaction"
          }
        }
      }
    },
    "/register_node": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Register with this node",
        "operationId": "registerNode",
        "parameters": [
          {
            "name": "Peer",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Peer"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/Chain"
            }
          }
        }
      }
    },
    "/register_with": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Register with a specific node",
        "operationId": "registerWithNode",
        "parameters": [
          {
            "name": "Peer",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Peer"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "success",
            "schema": {
              "$ref": "#/definitions/Chain"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Block": {
      "type": "object",
      "required": [
        "index",
        "transactions",
        "proofOfWork"
      ],
      "properties": {
        "index": {
          "type": "integer"
        },
        "prevBlockHash": {
          "description": "sha256 + nonce digest",
          "type": "string"
        },
        "proofOfWork": {
          "description": "sha256 digest",
          "type": "string"
        },
        "transactions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Transaction"
          }
        }
      }
    },
    "Chain": {
      "required": [
        "length",
        "chain",
        "peers"
      ],
      "properties": {
        "chain": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Block"
          }
        },
        "length": {
          "type": "integer"
        },
        "peers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Peer"
          }
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string",
          "example": "failure"
        }
      }
    },
    "Peer": {
      "type": "object",
      "properties": {
        "node_address": {
          "type": "string",
          "format": "url"
        }
      }
    },
    "Transaction": {
      "type": "object",
      "required": [
        "author",
        "content"
      ],
      "properties": {
        "author": {
          "type": "string",
          "format": "uuid",
          "example": "d290f1ee-6c54-4b01-90e6-d701748f0851"
        },
        "content": {
          "type": "string"
        }
      }
    }
  }
}`))
}
