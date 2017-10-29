/**
* @api {post} /register Register - Registra um novo usuário.
* @apiVersion 1.0.0
* @apiName Register
* @apiGroup Auth
*
* @apiUse ContentTypeJson
*
* @apiParamExample {json} Request-Example:
*     {
*       "email":"amanda@gmail.com",
*       "name":"amanda",
*       "password":as232ff,
*       "username":"amanda"
*     }
*
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 4,
*	 "name": "amanda",
*	 "username": "amanda",
*	 "email": "amanda@gmail.com"
*      }
*
*/

/**
* @api {post} /login Login - Autentica o usuário.
* @apiVersion 1.0.0
* @apiName Login
* @apiGroup Auth
*
* @apiUse ContentTypeJson
* @apiUse InvalidCredentials
*
* @apiParamExample {json} Request-Example:
*     {
*       "password":"as232ff",
*       "username":"amanda"
*     }
*
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDk0OTI0NjIsImlkIjoxLCJuYW1lIjoiQWxsaXNvbiBWLiJ9.hRpe6GDdZVqGYVNAl8OfPdoqyWfJRfwRG1i3PsM_ay0"
*      }
*
*/