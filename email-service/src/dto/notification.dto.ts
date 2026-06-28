export interface NotificationDto{
    to:string
    subject:string
    templateid:string
    params:Record<string,any>
}