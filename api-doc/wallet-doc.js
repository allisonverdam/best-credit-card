/**
* @api {get} /wallets/:wallet_id GetWallet
* @apiVersion 1.0.0
* @apiName GetWallet
* @apiDescription Retorna a carteira com o id passado por parametro.
* @apiGroup Wallet
* @apiUse AuthRequired
* @apiUse UserNotFoundError
*
* @apiError Forbidden A carteira não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This wallet does not belong to the authenticated user."
*     }
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 1,
*	 "current_limit": 750,
*	 "maximum_limit": 1000,
*	 "avaliable_limit": 250,
*	 "person_id": 1
*      }
**/

/**
* @api {post} /wallets CreateWallet
* @apiVersion 1.0.0
* @apiName CreateWallet
* @apiDescription Cria uma nova carteira.
* @apiGroup Wallet
* @apiUse AuthRequired
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 5,
*	 "current_limit": 0,
*	 "maximum_limit": 0,
*	 "avaliable_limit": 0,
*	 "person_id": 1
*      }
**/

/**
* @api {delete} /wallets/:wallet_id DeleteWallet
* @apiVersion 1.0.0
* @apiName DeleteWallet
* @apiDescription Apaga a carteira com o id passado por parametro.
* @apiGroup Wallet
* @apiUse AuthRequired
* @apiUse UserNotFoundError
*
* @apiError Forbidden A carteira não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This wallet does not belong to the authenticated user."
*     }
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 1,
*	 "current_limit": 100,
*	 "maximum_limit": 200,
*	 "avaliable_limit": 100,
*	 "person_id": 1
*      }
**/

/**
* @api {put} /wallets/:wallet_id UpdateWallet
* @apiVersion 1.0.0
* @apiName UpdateWallet
* @apiDescription Atualizar uma carteira.
* @apiGroup Wallet
* @apiUse AuthRequired
* @apiUse UserNotFoundError
*
* @apiError Forbidden Essa carteira não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This wallet does not belong to the authenticated user."
*     }
* @apiParamExample {json} Request-Example:
*      {
*	 "current_limit": 700
*      }
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 5,
*	 "current_limit": 700,
*	 "maximum_limit": 2000,
*	 "avaliable_limit": 0,
*	 "person_id": 1
*      }
**/