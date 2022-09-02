import express from 'express'
import http from 'http'
import path from 'path'
import { fileURLToPath } from 'url';

const app = express()
const svr = http.createServer(app)
const bind = '0.0.0.0'
const port = 4000
const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)

app.use(express.static(path.join(__dirname, 'static')))

app.use((_, resp, next) => {
    resp.header('Cross-Origin-Opener-Policy', 'same-origin')
    resp.header('Cross-Origin-Embedder-Policy', 'require-corp')
    next()
})

svr.listen(port, bind, () => {
    console.log(`Server running at http://localhost:${port}/`)
})