import { RoomCategoryRepository } from "../repositary/roomCategory.repository";
import { createRoomCategoryDto } from "../dtos/roomCategory.dto";
import { HotelRepository } from "../repositary/hotel.repositry";
import { NotFoundError } from "../utils/errors/app.error";
const roomCategoryRepository=new RoomCategoryRepository()
const hotelRepository=new HotelRepository()
export async function createRoomCategoryService(createRoomCategoryDto:createRoomCategoryDto) {
    const roomCategory=await roomCategoryRepository.create(createRoomCategoryDto)
    return roomCategory
}
export async function getRoomCategoryByIdService(id:number) {
    const roomCategory=await roomCategoryRepository.findById(id);
    return roomCategory
}
export async function getAllRoomCategoryByHOtelIdService(hotelId:number) {
    const hotel=await hotelRepository.findById(hotelId)
    if(!hotel){
        throw new NotFoundError(`hotel with ${hotelId} not found`)
    }
    const roomCategories=await roomCategoryRepository.findAllByHotelId(hotelId)
    return roomCategories
}

export async function deletRoomCategoryService(id:number) {
    const roomCategory=await roomCategoryRepository.findById(id);
    if(!roomCategory){
        throw new NotFoundError(`room category with ifd ${id} not found `)

    }
    await roomCategoryRepository.delete({id})
    return true
}