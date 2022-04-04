use std::cmp::Ordering;

/// The Explode function breaks a string into an array.
///
/// **Arguments:**
///
/// * `separator` - Specifies where to break the string.
/// * `str` - The string to split.
/// * `limit` - Specifies the number of array elements to return.
/// If limit is greater than 0, then an array with a maximum of limit
/// elements is returned. If limit is less than 0 then an array excluding
/// the last -limit elements is returned. If limit is 0 then an array with
/// one element is returned.
///
/// **Returns:**
///
/// An array of strings.
pub fn explode<Limit>(separator: &str, str: &str, limit: Limit) -> Result<Vec<String>, ExplodeError>
where
    Limit: Into<Option<i32>> + PartialEq + std::fmt::Debug,
{
    if separator.is_empty() {
        return Err(ExplodeError::EmptySeparator);
    }

    match limit.into() {
        Some(limit) => match limit.cmp(&0) {
            Ordering::Equal => Ok(str
                .split(separator)
                .take(1)
                .map(|s| s.to_string())
                .collect()),
            Ordering::Greater => Ok(str
                .split(separator)
                .take(limit as usize)
                .map(|s| s.to_string())
                .collect()),
            Ordering::Less => {
                // This is a sad non one-line solution since Rust doesn't have
                // LINQ's SkipLast or Clojure's drop-last.
                let mut result_vec: Vec<String> =
                    str.split(separator).map(|s| s.to_string()).collect();
                let len = result_vec.len().saturating_sub(limit.abs() as usize);
                result_vec.truncate(len);
                Ok(result_vec)
            }
        },
        None => Ok(str.split(separator).map(|s| s.to_string()).collect()),
    }
}

#[derive(Debug, Clone, PartialEq)]
pub enum ExplodeError {
    EmptySeparator,
}
