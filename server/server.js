import express from 'express'

const app = express()

app.use(express.json())

const PORT = 8080

app.get('/api/sanitycheck', (_,res)=>{
    res.status(200).send({
        message:"hello world!"
    })
})

app.listen(PORT, ()=>{
    console.log(`Server listening on http://localhost:${PORT}`)
})