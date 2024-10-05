use anchor_lang::prelude::*;

use crate::state::{CarbonMeasurement, Consensus};
use crate::errors::DataSourceError;


#[derive(Accounts)]
#[instruction(name: String)]
pub struct Carbon<'info> {
    #[account(mut)]
    pub signer: Signer<'info>,
    #[account(
        init,
        payer=signer,
        space = CarbonMeasurement::INIT_SPACE,
        seeds = [b"carbon_measurement", name.as_str().as_bytes()],
        bump
    )]
    pub carbon_measurement: Box<Account<'info, CarbonMeasurement>>,
    #[account()]
    pub consensus:Box<Account<'info, Consensus>>,
    pub system_program: Program<'info, System>,
}

impl<'info> Carbon<'info> {

    fn is_signer_in_consensus(signers:Vec<Pubkey>,signer:Pubkey) -> bool{
        for key in signers {
            if key == signer {
                return true;
            }
        }
        return false;
    }

    pub fn create(&mut self,coordinate_x:f64,coordinate_y:f64,value:f32, bumps:&CarbonBumps)->Result<()> {
        require!(Self::is_signer_in_consensus(self.consensus.participants.clone(),self.signer.key()),DataSourceError::NotInConsensus);
        self.carbon_measurement.set_inner(CarbonMeasurement{
            coordinate_x,
            value,
            coordinate_y,
            bump:bumps.carbon_measurement
        }); 
        Ok(())
    }

}