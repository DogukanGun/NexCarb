use anchor_lang::prelude::*;

#[account]
pub struct Consensus {
    pub admin: Pubkey,
    pub participants: Vec<Pubkey>,
    pub bump: u8,
}
const MAX_PARTICIPANT_KEYS: usize = 32;

impl Space for Consensus {
    const INIT_SPACE: usize = 8 +  // Discriminator
    32 + // admin
    4 + (32 * MAX_PARTICIPANT_KEYS); // participants (length + maximum number of Pubkeys)
}
