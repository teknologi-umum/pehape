#[cfg(test)]
mod tests {
    use pehape_rs::bin_2_hex;

    #[test]
    fn test_bin_2_hex() {
        assert_eq!(bin_2_hex("Hello World!!"), Ok(String::from("48656c6c6f20576f726c642121")));
        assert_eq!(bin_2_hex("hacktoberfest1234!@#"), Ok(String::from("6861636b746f6265726665737431323334214023")));
    }
}