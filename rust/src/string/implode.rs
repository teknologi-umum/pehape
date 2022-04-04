/// The Implode function joins array elements with a string.
///
/// **Arguments:**
///
/// * `separator` - Specifies what to put between the array elements.
/// * `array` - The array to join to a string.
///
/// **Returns:**
///
/// A string from elements of an array
pub fn implode<T>(separator: &str, array: &[T]) -> String
where
    T: ToString,
{
    array
        .iter()
        .map(|v| v.to_string())
        .collect::<Vec<String>>()
        .join(separator)
}
