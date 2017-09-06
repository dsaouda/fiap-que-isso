<?php

namespace App;

use App\Factory\MongoFactory;

class ConfigProvider
{

    public function __invoke()
    {
        return [
            'dependencies' => $this->getDependencies(),
            'templates'    => $this->getTemplates(),
        ];
    }

    public function getDependencies()
    {
        return [
            'invokables' => [
            ],
            'factories'  => [
                MongoFactory::class => MongoFactory::class,
                Action\CardsAction::class => Factory\CardsFactory::class,
                Action\GroupsAction::class => Factory\GroupsFactory::class,
            ],
        ];
    }

    public function getTemplates()
    {
        return [
            'paths' => [
                'app'    => ['templates/app'],
                'error'  => ['templates/error'],
                'layout' => ['templates/layout'],
            ],
        ];
    }
}
