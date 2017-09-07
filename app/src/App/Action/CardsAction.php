<?php

namespace App\Action;

use Interop\Http\ServerMiddleware\DelegateInterface;
use Interop\Http\ServerMiddleware\MiddlewareInterface as ServerMiddlewareInterface;
use MongoDB\Driver\Manager;
use Zend\Diactoros\Response\JsonResponse;
use Psr\Http\Message\ServerRequestInterface;

class CardsAction implements ServerMiddlewareInterface
{
    /**
     * @var Manager
     */
    private $mongo;

    public function __construct(Manager $mongo)
    {
        $this->mongo = $mongo;
    }

    public function process(ServerRequestInterface $request, DelegateInterface $delegate): JsonResponse
    {
        $q = $request->getQueryParams()['q'] ?? null;
        $m = ($request->getQueryParams()['m'] ?? null) === 'true' ? true : false;
        
        $filter = $q  ? ['group' => $q] : [];

        $options = ['sort' => ['group' => 1, 'order' => 1]];
        $query = new \MongoDB\Driver\Query($filter, $options);
    	
		$cursor = $this->mongo->executeQuery('fiap_que_isso.slide', $query);

        $result = $cursor->toArray();
        if ($m) {
            shuffle($result);
        }


        return new JsonResponse($result);
    }
}
