import { createClient, defineScript } from 'redis'

(async () => {
    let cli = createClient({
        url : 'redis://192.168.0.150:6379', 
        scripts: {
            mincr: defineScript({
                NUMBER_OF_KEYS: 2,
                SCRIPT: `
                return {
                    redis.pcall("INCRBY", KEYS[1], ARGV[1]),
                    redis.pcall("INCRBY", KEYS[2], ARGV[1])
                }`,
                transformArguments(key1, key2, arg) {
                    return [key1, key2, arg.toString()]
                }
            })
        }
    })
    cli.on('error', (err) => {
        console.error(`redis client Error ${err}`)
    })
    await cli.connect()
    
    let info = await cli.info()
    console.log(info)
    await cli.set('a1', 5)
    let res = await cli.mincr('a1', 'a2', 10)
    console.log(res)

    // close connection
    await cli.quit()
})()