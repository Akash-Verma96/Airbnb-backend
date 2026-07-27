import z from "zod";

export const pingSchema = z.object({
    name: z.string().min(3),
    age: z.number().int().positive()
})

export const querySchema = z.object({
    search: z.string().min(3),
    company: z.string().min(3)
})