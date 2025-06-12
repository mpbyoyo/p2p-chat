import z from "zod"

export const userSchema = z.object({
    username: z.string().nullable(),
})

export type UserType = z.infer<typeof userSchema>
