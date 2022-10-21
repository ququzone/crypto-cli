/// The version message for the current program, like
/// `cast-ext 0.1.0 (f01b232bc 2022-01-22T23:28:39.493201+00:00)`
pub const VERSION_MESSAGE: &str = concat!(
    env!("CARGO_PKG_VERSION"),
    " (",
    env!("VERGEN_GIT_SHA_SHORT"),
    " ",
    env!("VERGEN_BUILD_TIMESTAMP"),
    ")"
);
