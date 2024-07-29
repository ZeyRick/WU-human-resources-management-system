export function nameToNumber(name: string) {
    let sum = 0
    for (let i = 0; i < name.length; i++) {
        sum += name.charCodeAt(i)
    }
    return sum
}

export function numberToColor(num: number) {
    const r = Math.floor((num * 3) % 256)
    const g = Math.floor((num * 5) % 256)
    const b = Math.floor((num * 7) % 256)
    return `rgb(${r}, ${g}, ${b})`
}

export function getTextColor(backgroundColor: any) {
    const colorPalette = ['#333', '#FF69B4', '#008080', '#FFA500', '#800080', '#FF00FF', '#00FF00', '#FFFF00']

    let bestColor = null
    let maxContrast = 0

    for (const color of colorPalette) {
        const contrast = calculateContrastRatio(backgroundColor, color)
        if (contrast > maxContrast && contrast >= 3.5) {
            // Lower contrast threshold
            bestColor = color
            maxContrast = contrast
        }
    }

    return bestColor || '#000' // Default to black if no suitable color found
}

export function generateColorAndText(name: string) {
    const number = nameToNumber(name)
    const backgroundColor = numberToColor(number)
    const textColor = getTextColor(backgroundColor)

    return {
        textColor,
        backgroundColor,
    }
}

function calculateContrastRatio(color1: any, color2: any) {
    // Convert hex to RGB
    const rgb1 = color1.match(/\d+/g).map(Number)
    const rgb2 = color2.match(/\d+/g).map(Number)

    // Calculate luminance
    const L1 = (0.2126 * rgb1[0]) / 255 + (0.7152 * rgb1[1]) / 255 + (0.0722 * rgb1[2]) / 255
    const L2 = (0.2126 * rgb2[0]) / 255 + (0.7152 * rgb2[1]) / 255 + (0.0722 * rgb2[2]) / 255

    // Calculate contrast ratio
    const lighter = Math.max(L1, L2)
    const darker = Math.min(L1, L2)
    const contrastRatio = (lighter + 0.05) / (darker + 0.05)
    return contrastRatio
}
