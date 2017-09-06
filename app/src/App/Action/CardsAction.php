<?php

namespace App\Action;

use Interop\Http\ServerMiddleware\DelegateInterface;
use Interop\Http\ServerMiddleware\MiddlewareInterface as ServerMiddlewareInterface;
use Zend\Diactoros\Response\JsonResponse;
use Psr\Http\Message\ServerRequestInterface;

class CardsAction implements ServerMiddlewareInterface
{
    public function process(ServerRequestInterface $request, DelegateInterface $delegate)
    {
    	$manager = new \MongoDB\Driver\Manager('mongodb://192.168.200.30:27017');
        

        $q = $request->getQueryParams()['q'] ?? null;


        $filter = $q  ? ['group' => $q] : [];


        $options = ['sort' => ['group' => 1, 'order' => 1]];
        $query = new \MongoDB\Driver\Query($filter, $options);

    	
		$cursor = $manager->executeQuery('fiap_que_isso.slide', $query);
        return new JsonResponse($cursor->toArray());
    }
}
