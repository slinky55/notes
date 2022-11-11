import { PrismaClient } from '@prisma/client'

export const prisma = new PrismaClient();

async function createNote(_title: string, _contents: string) {
    await prisma.note.create({
        data: {
            title: _title,
            content: _contents,
        }
    })
}