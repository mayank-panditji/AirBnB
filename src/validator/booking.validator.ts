import {z} from 'zod'
export const createBookingschema=z.object({
    userId:z.number({message:"user id is required"}),
    hotelId:z.number({message:"hotel id must be present"}), 
    totalGuests:z.number({message:"total guest must be present"}).min(1,{message:"total guest must be atleast 1"}),
    bookingAmount:z.number({message:"booking must be present"}).min(1,{message:"booking must be atleast 1"})
})