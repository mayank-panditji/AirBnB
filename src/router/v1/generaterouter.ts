import { Router } from "express";
import { generateRoomsController } from "../../controller/roomgeneration";

const router = Router();

router.post("/room-categories/:roomCategoryId/generate-rooms", generateRoomsController);
// ya body se le rahe ho to bina param ke bhi chalega

export default router;