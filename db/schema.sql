CREATE TABLE IF NOT EXISTS loans (
    id TEXT PRIMARY KEY,
    borrower_id TEXT NOT NULL,
    principal_amount REAL NOT NULL,
    rate REAL NOT NULL,
    roi REAL NOT NULL,
    agreement_letter_url TEXT,
    state TEXT DEFAULT 'proposed',
    total_invested REAL DEFAULT 0,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS loan_details (
    loan_id TEXT PRIMARY KEY,
    field_validator_id TEXT,
    visit_proof_url TEXT,
    approved_at DATETIME,
    field_officer_id TEXT,
    signed_agreement_url TEXT,
    disbursed_at DATETIME,
    FOREIGN KEY (loan_id) REFERENCES loans(id)
);

CREATE TABLE IF NOT EXISTS investments (
    id TEXT PRIMARY KEY,
    loan_id TEXT NOT NULL,
    investor_id TEXT NOT NULL,
    amount REAL NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (loan_id) REFERENCES loans(id)
);

CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    username TEXT NOT NULL,
    role TEXT NOT NULL, 
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS flags (
    key TEXT PRIMARY KEY,
    value TEXT NOT NULL
);
