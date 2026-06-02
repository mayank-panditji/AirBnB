import { createHotel, getHotelById } from "../repositary/hotel.repositry";
import { createHotelDTO } from "../dtos/hotel.dto";
import { BadRequestError } from "../utils/errors/app.error";


export async function createHotelService(hotelData: createHotelDTO){
  
    const hotel = await createHotel(hotelData);
    return hotel;
}
export async function getHotelByIdService(id: number){
    const hotel = await getHotelById(id);
    return hotel;
}