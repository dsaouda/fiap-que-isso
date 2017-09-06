<?php

namespace App\Action;

use Interop\Http\ServerMiddleware\DelegateInterface;
use Interop\Http\ServerMiddleware\MiddlewareInterface as ServerMiddlewareInterface;
use MongoDB\Driver\Manager;
use Zend\Diactoros\Response\JsonResponse;
use Psr\Http\Message\ServerRequestInterface;

class GroupsAction implements ServerMiddlewareInterface
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
    	$collection = 'slide';
    	$field = 'group';

		$cmd = new \MongoDB\Driver\Command([
		    'distinct' => $collection,
		    'key' => $field,
		]);

		$cursor = $this->mongo->executeCommand('fiap_que_isso', $cmd);
        return new JsonResponse($cursor->toArray()[0]->values);
    }
}
