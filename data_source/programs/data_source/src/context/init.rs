use anchor_lang::prelude::*;

use crate::state::Consensus;

#[derive(Accounts)]
pub struct Init<'info> {
    #[account(mut)]
    pub signer: Signer<'info>,
    #[account(
        init,
        payer=signer,
        space = Consensus::INIT_SPACE,
        seeds = [b"consensus"],
        bump
    )]
    pub consensus: Box<Account<'info, Consensus>>,
    pub system_program: Program<'info, System>,
}

impl<'info> Init<'info> {
    pub fn init_consensus(&mut self,bumps:&InitBumps) -> Result<()> {
        self.consensus.set_inner(Consensus {
            admin: self.signer.key(),
            participants: Vec::new(),
            bump: bumps.consensus
        });
        Ok(())
    }
}
