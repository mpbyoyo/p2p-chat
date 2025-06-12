import z from 'zod'

export const configSchema = z.object({
    rememberUser: z.boolean()
})

export type ConfigType = z.infer<typeof configSchema>