/**
* @apiDefine ValidatePrice
* @apiErrorExample ValidatePrice:
*     HTTP/1.1 400 Bad Request
*     {
*          "error_code": "INVALID_DATA",
*          "message": "There is some problem with the data you submitted. See details for more information.",
*          "details": [
*               {
*                    "field": "price",
*                    "error": "must be no less than 0"
*               }
*          ]
*     }
**/