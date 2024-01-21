import { z } from "zod";

export const ApiLoginSchema = z.string();
export type ApiLogin = z.infer<typeof ApiLoginSchema>;

export const ApiUserSchema = z.object({
  id: z.number(),
  name: z.string(),
  email: z.string().email(),
});
export type ApiUser = z.infer<typeof ApiUserSchema>;

export const ApiMemorySchema = z.object({
  id: z.number(),
  users_id: z.number(),
  title: z.string(),
  date: z.string().regex(/^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/),
  location: z.string(),
  detail: z.string(),
});
export type ApiMemory = z.infer<typeof ApiMemorySchema>;

export const ApiMemoriesSchema = z.array(ApiMemorySchema);
export type ApiMemories = z.infer<typeof ApiMemoriesSchema>;
