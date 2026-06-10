import {Elysia} from "elysia";
import {appointmentsRoutes} from "./appointments/routes";
import {HttpError} from "./errors";

const app = new Elysia()
    .onError(({error, set}) => {
        if (error instanceof HttpError) {
            set.status = error.status;
            return {message: error.message};
        }
    })
    .get("/", () => "Hello Elysia")
    .use(appointmentsRoutes)
    .listen(3000);

console.log(`🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`);
