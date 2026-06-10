import {confirmBooking, createBooking, createidempotencyKey,finalizeIdempotencyKey, getIdempotencykey} from "../repositories/booking.repositorie";
import { BadRequestError, NotFoundError } from "../utils/errors/app.error";
import { generateIdempotencyKey } from "../utils/generateidempotencykey";
import { CreateBookingDTO } from "../dto/booking.dto";
export async function createBookingService(
    createBookingDTO:CreateBookingDTO
) {
const booking=await createBooking({
    userId:createBookingDTO.userId,
    hotelId:createBookingDTO.hotelId,
    totalGuests:createBookingDTO.totalGuests,
    bookingAmount:createBookingDTO.bookingAmount,
})
const idempotencyKey=generateIdempotencyKey()
await createidempotencyKey(idempotencyKey,booking.id)
return{
    bookingId:booking.id,
    idempotencyKey:idempotencyKey,
}
}
export async function confirmBookingService(idempotencyKey:string){
    const idempotencyKeyData=await getIdempotencykey(idempotencyKey)
    if(!idempotencyKeyData){
        throw new NotFoundError("Idempotency key not found")
    }
    if(idempotencyKeyData.finalized){
        throw new BadRequestError("Idempotency key already finalized")
    }
    const booking =await confirmBooking(idempotencyKeyData.bookingId)
    await finalizeIdempotencyKey(idempotencyKey)
    return booking;
}