import {AsyncLocalStorage} from 'async_hooks';
import { Request, Response, NextFunction } from 'express';
type AsynLocalStorageType={
    correlationId:string
}
 export const asyncLocalStorage = new AsyncLocalStorage<AsynLocalStorageType>();
export const getCorrelationId=()=>{
    const asyncstore = asyncLocalStorage.getStore();
    return asyncstore?.correlationId || 'unknown error while creating correlation id';
}