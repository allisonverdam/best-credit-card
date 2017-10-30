/**
* @api {post} /cards/best-card GetBestCards
* @apiVersion 1.0.0
* @apiName GetBestCards
* @apiGroup Card
* @apiDescription Retorna o melhor cartão para a compra.
*
* @apiUse AuthRequired
*
* @apiParamExample {json} Request-Example:
*     {
*       "price": 100,
		"wallet_id": 1
*     }
*
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*     [
*      {
*	 "id": 3,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "avaliable_limit": 450,
*	 "wallet_id": 1
*      }
*     ]
*
* @apiUse UserNotFoundError
* @apiUse ValidatePrice
*
* @apiError Forbidden O parametro 'wallet_id' informado não pertence ao usuário autenticado.
* @apiError ValidatePrice O parametro 'price' não pode ser menor que 0.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This wallet does not belong to the authenticated user."
*     }
*/

/**
* @api {post} /cards/pay PayCreditCard
* @apiVersion 1.0.0
* @apiName PayCreditCard
* @apiDescription Pagar um cartão para liberar crédito.
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse UserNotFoundError
* @apiUse ValidatePrice
*
* @apiError Forbidden O parametro 'wallet_id' informado não pertence ao usuário autenticado.
* @apiError ValidatePrice O parametro 'price' não pode ser menor que 0.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This card does not belong to the authenticated user."
*     }
* @apiParamExample {json} Request-Example:
*     {
*       "price": 100,
		"card_id": 1
*     }
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 1,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "avaliable_limit": 450,
*	 "wallet_id": 1
*      }
**/

/**
* @api {get} /cards/:card_id GetCard
* @apiVersion 1.0.0
* @apiName GetCard
* @apiDescription Retorna o cartao com o id passado por parametro.
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse UserNotFoundError
*
* @apiError Forbidden O cartão não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This card does not belong to the authenticated user."
*     }
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 3,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "avaliable_limit": 450,
*	 "wallet_id": 1
*      }
**/

/**
* @api {get} /cards/wallets/:wallet_id GetWalletCards
* @apiVersion 1.0.0
* @apiName GetWalletCards
* @apiDescription Retorna a lista de cartões de uma determinada carteira.
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse UserNotFoundError
*
* @apiError Forbidden O cartão não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This wallet does not belong to the authenticated user."
*     }
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*     [
*      {
*	 "id": 3,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "avaliable_limit": 450,
*	 "wallet_id": 1
*      }
*     ]
**/

/**
* @api {post} /cards CreateCard
* @apiVersion 1.0.0
* @apiName CreateCard
* @apiDescription Cria um novo cartão.
* @apiGroup Card
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
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "avaliable_limit": 450,
*	 "wallet_id": 1
*      }
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 3,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "avaliable_limit": 450,
*	 "wallet_id": 1
*      }
**/

/**
* @api {delete} /cards/:card_id DeleteCard
* @apiVersion 1.0.0
* @apiName DeleteCard
* @apiDescription Apaga o cartao com o id passado por parametro.
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse UserNotFoundError
*
* @apiError Forbidden O cartão não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This card does not belong to the authenticated user."
*     }
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 3,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "cvv": 123,
*	 "real_limit": 500,
*	 "avaliable_limit": 450,
*	 "wallet_id": 1
*      }
**/

/**
* @api {put} /cards/:card_id UpdateCard
* @apiVersion 1.0.0
* @apiName UpdateCard
* @apiDescription Atualizar um cartão.
* @apiGroup Card
* @apiUse AuthRequired
* @apiUse UserNotFoundError
*
* @apiError Forbidden Essa carteira não pertence ao usuário autenticado.
* @apiErrorExample Forbidden:
*     HTTP/1.1 403 Forbidden
*     {
*	"error_code": "FORBIDDEN",
*	"message": "You're not allowed to do this.",
*	"developer_message": "This card does not belong to the authenticated user."
*     }
* @apiParamExample {json} Request-Example:
*      {
*	 "real_limit": 700,
*	 "current_limit": 550,
*      }
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 3,
*	 "number": "1234123412341232",
*	 "due_date": 11,
*	 "expiration_month": 8,
*	 "expiration_year": 16,
*	 "cvv": 123,
*	 "real_limit": 700,
*	 "avaliable_limit": 550,
*	 "wallet_id": 1
*      }
**/