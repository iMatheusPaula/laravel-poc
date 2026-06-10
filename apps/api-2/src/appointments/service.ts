import {desc, eq} from "drizzle-orm";
import {db} from "../db/client";
import {appointments} from "../db/schema";
import {HttpError} from "../errors";
import {PubSubService} from "../pubsub";
import {AppointmentStatus} from "./appointment-status";

type CreateAppointmentInput = {
    contactName: string;
    contactEmail: string;
    scheduledAt: string;
};

export abstract class AppointmentService {
    static list() {
        return db
            .select()
            .from(appointments)
            .orderBy(desc(appointments.createdAt));
    }

    static async create(input: CreateAppointmentInput) {
        const [appointment] = await db
            .insert(appointments)
            .values({
                contactName: input.contactName,
                contactEmail: input.contactEmail,
                scheduledAt: input.scheduledAt,
                status: AppointmentStatus.PENDING,
            })
            .returning();

        await PubSubService.publish(
            {
                appointment_id: appointment.id,
                contact_name: appointment.contactName,
                contact_email: appointment.contactEmail,
                scheduled_at: appointment.scheduledAt,
            },
            "appointments.created"
        );

        return appointment;
    }

    static async cancel(id: number) {
        const [found] = await db
            .select()
            .from(appointments)
            .where(eq(appointments.id, id));

        if (!found) {
            throw new HttpError(404, "Appointment not found.");
        }

        if (found.status !== AppointmentStatus.PENDING) {
            throw new HttpError(409, "Only pending appointments can be canceled.");
        }

        const [updated] = await db
            .update(appointments)
            .set({
                status: AppointmentStatus.CANCELED,
                canceledAt: new Date().toISOString(),
            })
            .where(eq(appointments.id, id))
            .returning();

        return updated;
    }
}
