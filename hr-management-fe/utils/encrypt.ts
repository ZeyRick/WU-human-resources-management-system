import CryptoJS from 'crypto-js'

export function encrypteData(data: string) {
    const config = useRuntimeConfig()
    const ivKey = '03f6b349a565fcdc'
    const keyValue = config.public.aesKey
    if (data) {
        const key = CryptoJS.enc.Utf8.parse(keyValue)
        const iv = CryptoJS.enc.Utf8.parse(ivKey) // Convert string to WordArray
        const encrypted = CryptoJS.AES.encrypt(CryptoJS.enc.Utf8.parse(data), key, {
            iv: iv,
            mode: CryptoJS.mode.CBC,
            padding: CryptoJS.pad.Pkcs7,
        })
        return encrypted.ciphertext.toString(CryptoJS.enc.Base64)
    }
}

export function decrypteData(data: string) {
    const config = useRuntimeConfig()
    const ivKey = '03f6b349a565fcdc'
    const keyValue = config.public.aesKey
    if (data) {
        const key = CryptoJS.enc.Utf8.parse(keyValue)
        const iv = CryptoJS.enc.Utf8.parse(ivKey)
        const decrypted = CryptoJS.AES.decrypt(data, key, {
            iv: iv,
            mode: CryptoJS.mode.CBC,
            padding: CryptoJS.pad.Pkcs7,
        })
        const decryptedStr = CryptoJS.enc.Utf8.stringify(decrypted)
        return decryptedStr
    }
    return ''
}
