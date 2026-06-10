import { v4 as uuid } from 'uuid';
import { Request, Response, NextFunction } from 'express';
import { asyncLocalStorage } from '../utils/helpers/request.helper';
export const attachCorrelationIdMiddleware = (req:Request, res:Response, next:NextFunction) => {
    const correlationId = uuid();
    req.headers['X-Correlation-ID'] = correlationId;
    asyncLocalStorage.run({ correlationId:correlationId }, () => {
        next();
    });
    // next();
}