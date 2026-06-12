import fs from 'fs/promises'
import path from 'path'
import Handlebars from 'handlebars'
import { InternalServerError } from '../utils/errors/app.error'

export async function renderMailTemplate(templateId:string,params:Record<string,any>):Promise<string>{
const templatepath=path.join(__dirname,'mailertemplate',`${templateId}.hbs`)
try{
    const content =await fs.readFile(templatepath,'utf-8')
    const finaltemplate=Handlebars.compile(content)
    return finaltemplate(params)
}catch(error){
    throw new InternalServerError( `template not found:${templateId}`)
}
}