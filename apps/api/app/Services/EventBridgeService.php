<?php

namespace App\Services;

use Aws\EventBridge\EventBridgeClient;

class EventBridgeService
{
    protected EventBridgeClient $client;

    public function __construct()
    {
        $this->client = new EventBridgeClient([
            'region' => config('services.aws.region'),
            'credentials' => [
                'key' => config('services.aws.key'),
                'secret' => config('services.aws.secret'),
            ],
        ]);
    }

    public function put(array $detail, string $detailType, string $source): void
    {
        $this->client->putEvents([
            'Entries' => [[
                'EventBusName' => env('EVENTBRIDGE_BUS', 'default'),
                'Source' => $source,
                'DetailType' => $detailType,
                'Detail' => json_encode($detail),
            ]],
        ]);
    }
}
