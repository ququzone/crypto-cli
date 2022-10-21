use clap::{Parser, Subcommand};

use cast_ext::utils;

#[derive(Parser, Debug)]
#[clap(name = "cast", version = crate::utils::VERSION_MESSAGE)]
pub struct Opts {
    #[clap(subcommand)]
    pub sub: Subcommands,
}

#[derive(Debug, Subcommand)]
#[clap(
    about = "Extension for cast.",
    after_help = "Find more information: https://github.com/ququzone/crypto-cli",
    next_display_order = None
)]
pub enum Subcommands {
    #[clap(name = "wallet")]
    #[clap(visible_aliases = &["w"])]
    #[clap(about = "Wallet management utilities.")]
    Wallet,
}

fn main() {
    let opts = Opts::parse();
    match opts.sub {
        // Constants
        Subcommands::Wallet => {
            println!("{}", "TODO");
        }
    }
}
