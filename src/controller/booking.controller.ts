// controller/booking.controller.ts

import { Request, Response, NextFunction } from 'express'
import { createBookingService, confirmBookingService } from '../services/booking.service'

export const createBookingHandler = async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    try {
        const booking = await createBookingService(req.body);
        res.status(201).json({
            message: 'Booking created successfully',
            success: true,
            bookingId: booking.bookingId,
            idempotencyKey: booking.idempotencyKey
        });
    } catch (err) {
        next(err);
    }
}

export const confirmBookingHandler = async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    try {
        const booking = await confirmBookingService(req.params.idempotencyKey as string);
        res.status(200).json({
            message: 'Booking confirmed successfully',
            success: true,
            bookingId: booking.id,
            status: booking.status
        });
    } catch (err) {
        next(err);
    }
}