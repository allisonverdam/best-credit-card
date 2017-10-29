define({ "api": [
  {
    "type": "post",
    "url": "/cards/best-card",
    "title": "Retorna o melhor cartão para a compra.",
    "version": "1.0.0",
    "name": "getBestCards",
    "group": "Card",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "price",
            "description": "<p>Users unique ID.</p>"
          },
          {
            "group": "Parameter",
            "type": "Number",
            "optional": false,
            "field": "wallet_id",
            "description": "<p>Users unique ID.</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "firstname",
            "description": "<p>Firstname of the User.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "lastname",
            "description": "<p>Lastname of the User.</p>"
          },
          {
            "group": "Success 200",
            "type": "Date",
            "optional": false,
            "field": "registered",
            "description": "<p>Date of Registration.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "\tHTTP/1.1 200 OK\n\t[\n\t {\n   \t    \"id\": 3,\n   \t    \"number\": \"1234123412341232\",\n   \t    \"due_date\": 11,\n   \t    \"expiration_month\": 8,\n   \t    \"expiration_year\": 16,\n   \t    \"cvv\": 123,\n   \t    \"real_limit\": 500,\n   \t    \"current_limit\": 450,\n   \t    \"wallet_id\": 1\n\t }\n\t]",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "UserNotFound",
            "description": "<p>The id of the User was not found.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Error-Response:",
          "content": "HTTP/1.1 404 Not Found\n{\n  \"error\": \"UserNotFound\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./controllers/card.go",
    "groupTitle": "Card"
  },
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./doc/main.js",
    "group": "_home_born_work_src_github_com_allisonverdam_best_credit_card_doc_main_js",
    "groupTitle": "_home_born_work_src_github_com_allisonverdam_best_credit_card_doc_main_js",
    "name": ""
  }
] });