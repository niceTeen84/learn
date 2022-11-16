(async ()=>{
    console.log('test string to hex')
    // use prototype func to do
    String.prototype.hexEncode = function () {
        let ret = ''
        for(let i = 0; i < this.length; i++) {
            let hex = this.charCodeAt(i).toString(16)
            ret += ('000' + hex).slice(-4)
        }
        return ret
    }

    String.prototype.hexDecode = function () {
        let ret = ''
        let hexs = this.match(/.{1,4}/g) || []
        for (let i = 0; i < hexs.length; i++) {
            ret += String.fromCharCode(parseInt(hexs[i], 16))
        }
        return ret
    }
    console.log('γ'.hexEncode())
    console.log('03b3'.hexDecode())
})()