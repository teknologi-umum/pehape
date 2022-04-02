macro_rules! vec_of_strings {
    ($($x:expr),*) => (vec![$($x.to_string()),*]);
}
pub(crate) use vec_of_strings;

pub fn assert_result<T, E>(actual: Result<T, E>, expected: T)
where
    T: PartialEq + std::fmt::Debug,
{
    if let Ok(actual) = actual {
        assert_eq!(actual, expected);
    }
}

pub fn assert_error<T, E>(actual: Result<T, E>, expected: E)
where
    E: PartialEq + std::fmt::Debug,
{
    if let Err(actual) = actual {
        assert_eq!(actual, expected);
    }
}
