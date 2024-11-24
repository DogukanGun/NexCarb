use borsh::{BorshDeserialize, BorshSerialize};
use solana_program::{
    account_info::{next_account_info, AccountInfo},
    entrypoint,
    entrypoint::ProgramResult,
    msg,
    program_error::ProgramError,
    pubkey::Pubkey,
};

#[derive(BorshSerialize, BorshDeserialize, Debug)]
pub struct SeachAnswer {
    pub writter: Pubkey,
    pub inquirer: Pubkey,
    pub answer: String
}

entrypoint!(process_instruction);

pub fn process_instruction(
    program_id: &Pubkey,
    accounts: &[AccountInfo],
    _instruction_data: &[u8],
) -> ProgramResult {
    let accounts_iter = &mut accounts.iter();

    let account = next_account_info(accounts_iter)?;

    if account.owner != program_id {
        msg!("SeachAnswer account does not have the correct program id");
        return Err(ProgramError::IncorrectProgramId);
    }

    let result = String::from_utf8(_instruction_data.to_vec());
    let my_string = match result {
        Ok(string) => string, 
        Err(e) => {
            eprintln!("Error converting to String: {}", e);
            return;
        }
    };

    let mut search_account = SeachAnswer::try_from_slice(&account.data.borrow())?;
    search_account.answer = result;

    greeting_account.serialize(&mut &mut account.data.borrow_mut()[..])?;
    Ok(())
}