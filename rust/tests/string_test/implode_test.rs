use pehape_rs::implode;

use crate::common::vec_of_strings;

#[test]
pub fn can_implode_array_of_string() {
    let array = vec_of_strings!["Hello", "world"];

    assert_eq!(implode(" ", &array), "Hello world".to_string());
    assert_eq!(implode("  ", &array), "Hello  world".to_string());
    assert_eq!(implode("", &array), "Helloworld".to_string());
    assert_eq!(
        implode(" wkwkwk ", &array),
        "Hello wkwkwk world".to_string()
    );
    assert_eq!(implode("\n", &array), "Hello\nworld".to_string());
}

#[test]
pub fn can_implode_array_of_int() {
    let array = vec![1, 2, 3, 4, 5];

    assert_eq!(implode(":. ", &array), "1:. 2:. 3:. 4:. 5".to_string());
}

struct AnyObject(i32);
impl AnyObject {
    fn new(num: i32) -> AnyObject {
        AnyObject(num)
    }
}

impl ToString for AnyObject {
    fn to_string(&self) -> String {
        format!("AnyObject({})", self.0)
    }
}

#[test]
pub fn can_implode_any_object() {
    let array = vec![AnyObject::new(5), AnyObject::new(5), AnyObject::new(5)];
    assert_eq!(
        implode("-", &array),
        "AnyObject(5)-AnyObject(5)-AnyObject(5)".to_string()
    );
}

#[test]
pub fn can_implode_empty_array() {
    let array: Vec<String> = Vec::new();
    assert_eq!(implode(" ", &array), "".to_string());
}
