import { createClient, defineScript } from 'redis'

(async () => {
    let cli = createClient({
        url : 'redis://192.168.0.150:6379', 
        scripts: {
            mincr: defineScript({
                NUMBER_OF_KEYS: 2,
                SCRIPT: `return {
                    redis.pcall("INCRBY", KEYS[1], ARGV[1]),
                    redis.pcall("INCRBY", KEYS[2], ARGV[1])
                }`,
                transformArguments(key1, key2, arg) {
                    return [key1, key2, arg.toString()]
                }
            })
        }
    })

    await cli.connect()
    await cli.ping()
    // await cli.quit()

})()