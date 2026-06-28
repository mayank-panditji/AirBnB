export interface AppError extends Error{
    statusCode: number;
}
export class InternalServerError implements AppError{
    statusCode: number;
    message: string
    name: string;
    constructor(message: string){
        this.message = message;
        this.name = 'InternalServerError';
        this.statusCode = 500;
    }
}
export class BadRequestError implements AppError{
    statusCode: number;
    message: string
    name: string;
    constructor(message: string){
        this.message = message;
        this.name = 'BadRequestError';
        this.statusCode = 400;
    }
}
export class NotFoundError implements AppError{
    statusCode: number;
    message: string
    name: string;
    constructor(message: string){
        this.message = message;
        this.name = 'NotFoundError';
        this.statusCode = 404;
    }
}
