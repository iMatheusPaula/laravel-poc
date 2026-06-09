import {sql} from "drizzle-orm";
import {integer, sqliteTable, text} from "drizzle-orm/sqlite-core";

export const appointments = sqliteTable("appointments", {
    id: integer("id").primaryKey({autoIncrement: true}),
    contactName: text("contact_name").notNull(),
    contactEmail: text("contact_email").notNull(),
    scheduledAt: text("scheduled_at").notNull(),
    status: text("status").notNull(),
    canceledAt: text("canceled_at"),
    createdAt: text("created_at").notNull().default(sql`CURRENT_TIMESTAMP`),
    updatedAt: text("updated_at").notNull().default(sql`CURRENT_TIMESTAMP`),
});
