<?php

declare(strict_types=1);

namespace App\Services;

use Carbon\Carbon;
use Google\Cloud\PubSub\MessageBuilder;
use Google\Cloud\PubSub\PubSubClient;

/**
 * @see https://docs.cloud.google.com/pubsub/docs/publish-best-practices
 */
final class PubSubService
{
    private PubSubClient $client;

    public function __construct()
    {
        $keyPath = base_path(config('services.gcp.key_file'));

        if (!file_exists($keyPath)) {
            throw new \RuntimeException('Key file not found');
        }

        $this->client = new PubSubClient([
            'projectId' => config('services.gcp.project_id'),
            'credentials' => $keyPath,
        ]);
    }

    /**
     * Publica uma mensagem no tópico do Pub/Sub usando MessageBuilder.
     *
     * @see https://docs.cloud.google.com/pubsub/docs/publisher
     * @param array $data
     * @return void
     */
    public function publish(array $data): void
    {
        $topic = $this->client->topic('appointment-created');

        $message = new MessageBuilder()
            ->setData(json_encode($data))
            ->addAttribute('sent_at', Carbon::now()->toDateTimeString())
            ->build();

        $topic->publish($message);
    }
}
