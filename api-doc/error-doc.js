

/**
* @apiDefine UserNotFoundError
* @apiError NotFound Recurso não encontrado.
* @apiErrorExample NotFound:
*     HTTP/1.1 404 Not Found
*     {
*	"error_code": "NOT_FOUND",
*	"message": "the requested resource was not found."
*     }
*/

/**
* @apiDefine InvalidCredentials
* @apiError InvalidCredentials Login ou senha inválido.
* @apiErrorExample InvalidCredentials:
*     HTTP/1.1 401 Unauthorized
*     {
*	"error_code": "UNAUTHORIZED",
*	"message": "Authentication failed.",
*	"developer_message": "Authentication failed: Authentication failed."
*     }
*/