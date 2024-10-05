use anchor_lang::prelude::*;

#[account]
pub struct CarbonMeasurement {
    pub coordinate_x: f64,
    pub coordinate_y: f64,
    pub value: f32,
    pub bump: u8,
}


impl Space for CarbonMeasurement {
    const INIT_SPACE: usize = 8 +  // Discriminator
                            64 + // coordinate_x
                            64 + // coordinate_y (length + maximum length + bump)
                            32; // value (length + maximum number of Pubkeys)
}