import z from "zod";
export const pingschema=z.object({
    message:z.string().min(1)
})