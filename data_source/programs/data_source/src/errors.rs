use anchor_lang::error_code;

#[error_code]
pub enum DataSourceError {
    #[msg("Signer is not in consensus")]
    NotInConsensus
}