#[cfg(test)]
mod tests {
    use pehape_rs::hex_2_bin;

    #[test]
    fn test_hex_2_bin() {
        assert_eq!(hex_2_bin("48656c6c6f20576f726c642121"), Ok(String::from("Hello World!!")));
        assert_eq!(hex_2_bin("6861636b746f6265726665737431323334214023"), Ok(String::from("hacktoberfest1234!@#")));
    }
}