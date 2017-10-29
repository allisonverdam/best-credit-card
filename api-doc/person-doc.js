/**
* @api {get} /me GetPerson - Retorna o usuário autenticado.
* @apiVersion 1.0.0
* @apiName GetPerson
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
* @api {get} /me/wallets GetPersonWallets - Retorna as carteiras do usuário autenticado.
* @apiVersion 1.0.0
* @apiName GetPersonWallets
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
* @api {put} /me UpdatePerson - Atualiza o usuário autenticado.
* @apiVersion 1.0.0
* @apiName UpdatePerson
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