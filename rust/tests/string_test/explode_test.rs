use crate::common::{assert_error, assert_result, vec_of_strings};
use pehape_rs::{explode, ExplodeError};

#[test]
pub fn can_explode_space_separated_string_by_space() {
    let str = String::from("Hello pehape world");

    assert_result(
        explode(" ", &str, None),
        vec_of_strings!["Hello", "pehape", "world"],
    );
    assert_result(explode(" ", &str, 0), vec_of_strings!["Hello"]);
    assert_result(explode(" ", &str, 1), vec_of_strings!["Hello"]);
    assert_result(
        explode(" ", &str, 5),
        vec_of_strings!["Hello", "pehape", "world"],
    );
    assert_result(explode(" ", &str, -1), vec_of_strings!["Hello", "pehape"]);
    assert_result(explode(" ", &str, -3), vec_of_strings![]);
}

#[test]
pub fn can_explode_double_space_separated_string_by_space() {
    let str = String::from("Hello  pehape  world");

    assert_result(
        explode(" ", &str, None),
        vec_of_strings!["Hello", "", "pehape", "", "world"],
    );
}

#[test]
pub fn can_explode_space_separated_string_by_double_space() {
    let str = String::from("Hello pehape world");

    assert_result(
        explode("  ", &str, None),
        vec_of_strings!["Hello pehape world"],
    );
}

#[test]
pub fn explode_space_separated_string_by_empty_string_returns_error() {
    let str = String::from("Hello pehape world");

    assert_error(explode("", &str, None), ExplodeError::EmptySeparator);
}
