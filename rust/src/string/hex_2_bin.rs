#[derive(Debug, PartialEq)]
pub enum Hex2BinError {
    EmptyInput,
    InvalidLength,
    InvalidHexCharacter(std::num::ParseIntError),
}

const HEX_CHAR_LEN: usize = 2;
const HEX_BASE: u32 = 16;

// The hex_2_bin function converts a string of hexadecimal to the ascii characters.
// Parameter:
// - s => hexadecimal that is to be converted
// 
// Return:
// - a string of ASCII characters corresponding to the input hexadecimal
pub fn hex_2_bin(hex: &str) -> Result<String, Hex2BinError> {
    match hex.len() {
        0 => Err(Hex2BinError::EmptyInput),
        _ if hex.len() % 2 != 0 => Err(Hex2BinError::InvalidLength),
        _ => {
            let mut result = String::new();
            for i in (0..hex.len()).step_by(HEX_CHAR_LEN) {
                let byte_str = &hex[i..i + HEX_CHAR_LEN];
                let byte = u8::from_str_radix(byte_str, HEX_BASE).map_err(Hex2BinError::InvalidHexCharacter)?;
                result.push(byte as char);
            }
            Ok(result)
        }
    }
}