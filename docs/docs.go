// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/attendance": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Submits the current user's attendance for the active period",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "attendance"
                ],
                "summary": "Submit attendance",
                "responses": {
                    "200": {
                        "description": "message: Attendance submitted",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Submission failed",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/attendance-period": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Creates a new attendance period for the logged-in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "attendance"
                ],
                "summary": "Add attendance period",
                "parameters": [
                    {
                        "description": "Attendance Period Dates",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AttendancePeriodRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "message: Attendance period created",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request or validation error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "Authenticates a user and returns a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "token",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/overtime": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Submit overtime hours for the logged-in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "overtime"
                ],
                "summary": "Submit overtime",
                "parameters": [
                    {
                        "description": "Overtime hours",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.OvertimeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message: Overtime submitted successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request or submission failed",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/payroll/run": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Processes payroll for employees for a given period (admin only)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payroll"
                ],
                "summary": "Run payroll",
                "parameters": [
                    {
                        "description": "Payroll run details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RunPayrollRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message: Payroll processed successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request or processing error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/payslip-summary/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieves a summary of a payslip by payroll ID (for admin or report access)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payslip"
                ],
                "summary": "Get payslip summary",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Payroll ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Payslip summary data",
                        "schema": {
                            "$ref": "#/definitions/dto.PayslipSummary"
                        }
                    },
                    "400": {
                        "description": "Invalid payroll ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/payslip/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Generates a detailed payslip for the logged-in user based on a payroll ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "payslip"
                ],
                "summary": "Generate payslip",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Payroll ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Generated payslip data",
                        "schema": {
                            "$ref": "#/definitions/models.Payslip"
                        }
                    },
                    "400": {
                        "description": "Invalid request or payroll ID",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/reimbursement": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Submit a reimbursement request for the logged-in user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reimbursement"
                ],
                "summary": "Submit reimbursement",
                "parameters": [
                    {
                        "description": "Reimbursement details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ReimbursementRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "message: Reimbursement submitted successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request or submission error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AttendancePeriodRequest": {
            "type": "object",
            "required": [
                "end_date",
                "start_date"
            ],
            "properties": {
                "end_date": {
                    "description": "format: YYYY-MM-DD",
                    "type": "string"
                },
                "start_date": {
                    "description": "format: YYYY-MM-DD",
                    "type": "string"
                }
            }
        },
        "dto.OvertimeRequest": {
            "type": "object",
            "properties": {
                "hours": {
                    "type": "number"
                }
            }
        },
        "dto.PayslipSummary": {
            "type": "object",
            "properties": {
                "employees": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "fullName": {
                                "type": "string"
                            },
                            "totalPay": {
                                "type": "number"
                            },
                            "userID": {
                                "type": "integer"
                            }
                        }
                    }
                },
                "grandTotal": {
                    "type": "number"
                }
            }
        },
        "dto.ReimbursementRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                }
            }
        },
        "dto.RunPayrollRequest": {
            "type": "object",
            "required": [
                "attendance_period_id"
            ],
            "properties": {
                "attendance_period_id": {
                    "type": "integer"
                }
            }
        },
        "handlers.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Payslip": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "createdBy": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "overtimeHours": {
                    "type": "number"
                },
                "overtimePay": {
                    "type": "number"
                },
                "payrollID": {
                    "type": "integer"
                },
                "presentDays": {
                    "type": "integer"
                },
                "rateSalaryPerDay": {
                    "type": "number"
                },
                "rateSalaryPerHour": {
                    "type": "number"
                },
                "reimbursementTotal": {
                    "type": "number"
                },
                "totalTakeHome": {
                    "type": "number"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isAdmin": {
                    "type": "boolean"
                },
                "passwordHash": {
                    "type": "string"
                },
                "salary": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
