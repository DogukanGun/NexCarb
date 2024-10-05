use anchor_lang::prelude::*;
pub mod state;

pub mod context;
use context::*;

pub mod errors;
pub use errors::*;

declare_id!("5BK5A5qDxBXULtqb1cpHZN8k819PPw3yc6bHwziNxooj");

#[program]
pub mod data_source {
    use super::*;

    pub fn init_consensus(ctx: Context<Init>) -> Result<()>{
        ctx.accounts.init_consensus(&ctx.bumps)
    }

    pub fn measure_carbon(ctx: Context<Carbon>,coordinate_x:f64,coordinate_y:f64,value:f32) -> Result<()> {
        ctx.accounts.create(coordinate_x, coordinate_y, value,&ctx.bumps)
    }

}