import {Elysia, t} from "elysia";
import {AppointmentService} from "./service";

export const appointmentsRoutes = new Elysia({prefix: "/appointments"})
    .get("/", () => AppointmentService.list())
    .post(
        "/",
        async ({body, status}) => {
            const scheduledAt = new Date(body.scheduled_at.replace(" ", "T"));
            if (scheduledAt <= new Date()) {
                return status(422, {message: "scheduled_at must be in the future"});
            }

            const appointment = await AppointmentService.create({
                contactName: body.contact_name,
                contactEmail: body.contact_email,
                scheduledAt: body.scheduled_at,
            });

            return status(201, appointment);
        },
        {
            body: t.Object({
                contact_name: t.String({minLength: 1}),
                contact_email: t.String({format: "email"}),
                scheduled_at: t.String({
                    pattern: "^\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}$",
                }),
            }),
        }
    )
    .patch(
        "/:id/cancel",
        ({params}) => AppointmentService.cancel(params.id),
        {
            params: t.Object({
                id: t.Numeric(),
            }),
        }
    );
