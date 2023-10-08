#[derive(Debug, PartialEq)]
pub enum Bin2HexError {
    EmptyStr,
    NonASCII
}

// The bin_2_hex function converts a string of ASCII characters to their hexadecimal representation.
// Parameter:
// - s => string that is to be converted
// 
// Return:
// - A string of their hexadecimal representation
pub fn bin_2_hex(s: &str) -> Result<String, Bin2HexError> {
    if s.is_empty() {
        return  Err(Bin2HexError::EmptyStr);
    }

    let mut result: String = String::new();
    for c in s.as_bytes() {
        if !c.is_ascii() {
            return Err(Bin2HexError::NonASCII);
        }

        result.push_str(&format!("{:x}", c))
    }

    return Ok(result);
}