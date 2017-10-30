/**
* @api {get} /me GetAuthenticatedUser
* @apiVersion 1.0.0
* @apiName GetAuthenticatedUser
* @apiDescription Retorna o usuário autenticado.
* @apiGroup Person
* @apiUse AuthRequired
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 1,
*	 "name": "Allison V.",
*	 "username": "allisonverdam",
*	 "email": "allison@g.com"
*      }
**/

/**
* @api {get} /me/wallets GetAuthenticatedUserWallets
* @apiVersion 1.0.0
* @apiName GetAuthenticatedUserWallets
* @apiDescription Retorna as carteiras do usuário autenticado.
* @apiGroup Person
* @apiUse AuthRequired
*
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*     [
*      {
*	 "id": 3,
*	 "id": 1,
*	 "real_limit": 0,
*	 "maximum_limit": 0,
*	 "person_id": 1
*      },
*      {
*	 "id": 4,
*	 "real_limit": 0,
*	 "maximum_limit": 0,
*	 "person_id": 1
*      }
**/

/**
* @api {put} /me UpdateAuthenticatedUser
* @apiVersion 1.0.0
* @apiName UpdateGetAuthenticatedUserWallets
* @apiDescription Atualiza o usuário autenticado.
* @apiGroup Person
* @apiUse AuthRequired
*
* @apiParamExample {json} Request-Example:
*     {
*       "email": "allison2222@g.com",
*       "name": "allison",
*       "username": "allisonverdam"
*     }
* @apiSuccessExample Success-Response:
*     HTTP/1.1 200 OK
*      {
*	 "id": 1,
*	 "name": "Allison V.",
*	 "username": "allisonverdam",
*	 "email": "allison2222@g.com"
*      }
**/