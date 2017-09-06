<?php

namespace App\Factory;

use Interop\Container\ContainerInterface;
use MongoDB\Driver\Manager;

class MongoFactory
{
    public function __invoke(ContainerInterface $container): Manager
    {
        $env = getenv('MONGO');

        if ($env) {
            $url = $env;
        } else if (isset($container->get('config')['mongo'])) {
            $url = $container->get('config')['mongo'];
        } else {
            throw new \RuntimeException('Não foi possível configurar o mongodb');
        }

        return new \MongoDB\Driver\Manager($url);
    }
}