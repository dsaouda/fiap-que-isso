<?php

namespace App\Action;

use Interop\Http\ServerMiddleware\DelegateInterface;
use Interop\Http\ServerMiddleware\MiddlewareInterface as ServerMiddlewareInterface;
use Zend\Diactoros\Response\JsonResponse;
use Psr\Http\Message\ServerRequestInterface;

class GroupsAction implements ServerMiddlewareInterface
{
    public function process(ServerRequestInterface $request, DelegateInterface $delegate)
    {
    	$manager = new \MongoDB\Driver\Manager('mongodb://192.168.200.30:27017');
    	
    	$collection = 'slide';
    	$field = 'group';

		$cmd = new \MongoDB\Driver\Command([
		    'distinct' => $collection,
		    'key' => $field,
		]);

		$cursor = $manager->executeCommand('fiap_que_isso', $cmd);
        return new JsonResponse($cursor->toArray()[0]->values);
    }
}
