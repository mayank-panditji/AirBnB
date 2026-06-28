import {z} from "zod";
export const hotelSchema=z.object({
    name:z.string().min(1),
    address:z.string().min(1),
    location:z.string().min(1),
    ratingCount:z.number().optional(),
    rating:z.number().optional(),
})