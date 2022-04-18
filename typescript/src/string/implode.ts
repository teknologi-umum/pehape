export const implode = (separator: string | null, arrayOfString: string[]): string => {
    if (separator === null) separator = ''
    
    return arrayOfString.join(separator)
}