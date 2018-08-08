<?php
return [
    'components' => [
/*        'cache' => [ //缓存配置
            'class' => 'common\components\RedisCache',
            'keyPrefix' => 'chinabrands:',
            'redis' => [
                'hostname' => '192.168.6.176',
                'port' => 26386,
                'database' => 8,
            ]
        ],*/
        /*'cache_down' => [ //客户端下载缓存配置
            'class' => 'common\components\RedisCache',
            'keyPrefix' => 'chinabrands:',
            'redis' => [
                'hostname' => '192.168.6.176',
                'port' => 26386,
                'database' => 9,
            ]
        ],*/
        /*'monitor_cache' => [ //监控系统缓存配置
            'class' => 'common\components\RedisCache',
            'keyPrefix' => 'chinabrands:',
            'redis' => [
                'hostname' => '192.168.6.176',
                'port' => 26386,
                'database' => 9,
            ]
        ],*/
        'redis'=>[
            'class' => 'mojifan\redis\Connection',
            'servers'=>[
                'tcp://192.168.6.176:26390',
                'tcp://192.168.6.176:26391',
                'tcp://192.168.6.176:26392',
            ],
            'options'=>[
                'replication' => 'sentinel',
                // master-name
                'service' => 'sentinel-192.168.6.176-26388',
                'parameters' => [
                    //'password' => $secretpassword,
                    'database' => 0,
                    'prefix'=>'chinabrands:'
                ],
            ]
        ],
    ],
];
