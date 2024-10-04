use anchor_lang::prelude::*;

use crate::state::CarbonMeasurement;


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
    pub system_program: Program<'info, System>,
}

impl<'info> Carbon<'info>   {
    pub fn create(&mut self,coordinate_x:f64,coordinate_y:f64,value:f32)->Result<()> {
        self.carbon_measurement.set_inner(CarbonMeasurement{
            coordinate_x:coordinate_x,
            value:value,
            coordinate_y:coordinate_y,
        }); 
        Ok(())
    }
}