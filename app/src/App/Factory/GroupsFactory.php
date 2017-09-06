<?php

namespace App\Factory;

use App\Action\GroupsAction;
use Interop\Container\ContainerInterface;

class GroupsFactory
{
    public function __invoke(ContainerInterface $container): GroupsAction
    {
        return new GroupsAction($container->get(MongoFactory::class));
    }
}
