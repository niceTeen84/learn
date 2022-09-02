import express from 'express'
import http from 'http'
import path from 'path'
import { fileURLToPath } from 'url';

const app = express()
const svr = http.createServer(app)

const bind = '0.0.0.0'
const port = 4000

// in ES module can not use `path.join(__dirname, 'static')` directly
const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)

// response header needs to set first
// then to find static html file
app.use((_, resp, next) => {
    resp.header('Cross-Origin-Opener-Policy', 'same-origin')
    resp.header('Cross-Origin-Embedder-Policy', 'require-corp')
    next()
})

app.use(express.static(path.join(__dirname, 'static')))

svr.listen(port, bind, () => {
    console.log(`Server running at http://localhost:${port}/`)
})