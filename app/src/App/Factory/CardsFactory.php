<?php

namespace App\Factory;

use App\Action\CardsAction;
use Interop\Container\ContainerInterface;

class CardsFactory
{
    public function __invoke(ContainerInterface $container): CardsAction
    {
        return new CardsAction($container->get(MongoFactory::class));
    }
}
