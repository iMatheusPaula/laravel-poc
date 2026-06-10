import {PubSub} from "@google-cloud/pubsub";

const pubsub = new PubSub({
    projectId: process.env.GCP_PROJECT_ID,
});

export abstract class PubSubService {
    static async publish(data: object, topicName: string): Promise<void> {
        const topic = pubsub.topic(topicName);

        await topic.publishMessage({
            data: Buffer.from(JSON.stringify(data)),
            attributes: {
                sent_at: new Date().toISOString(),
            },
        });
    }
}
