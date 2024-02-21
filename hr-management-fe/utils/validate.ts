export function isValidGoogleCoordinate(s: string): boolean {
    // Regular expression for Google coordinates (latitude, longitude)
    const regex = /^[-+]?\d+(?:\.\d+)?,\s*[-+]?\d+(?:\.\d+)?$/

    // Match the string against the regular expression
    return regex.test(s)
}
