INSERT INTO tag (name, is_index, created_at, updated_at, is_deleted, category)
VALUES
('DCO', 0, NOW(), NOW(), 0, 2),
('WEB3', 0, NOW(), NOW(), 0, 2),
('NFT', 0, NOW(), NOW(), 0, 2),
('DEFI', 0, NOW(), NOW(), 0, 2),
('CRYPTO', 0, NOW(), NOW(), 0, 2),
('PASSIVE INCOME', 0, NOW(), NOW(), 0, 2),
('LEND', 0, NOW(), NOW(), 0, 2),
('DEX', 0, NOW(), NOW(), 0, 2),
('STAKER', 0, NOW(), NOW(), 0, 2),
('GAMEFI', 0, NOW(), NOW(), 0, 2),
('NODE', 0, NOW(), NOW(), 0, 2),
('COMMUNITY', 0, NOW(), NOW(), 0, 2);

INSERT INTO chain (chain_id, name, status, reated_at, updated_at, is_deleted)
VALUES
(11155111, 'Sepolia', 1, NOW(), NOW(), 0),
(5700, 'Syscoin Tanenbaum Testnet', 1, NOW(), NOW(), 0);

INSERT INTO chain_contract (chain_id, address, project, type, abi, created_at, updated_at, is_deleted)
VALUES
(5700, '0x5A4eA3a013D42Cfd1B1609d19f6eA998EeE06D30', 1, 1, '', NOW(), NOW(), 0),
(5700, '0x86B5df6FF459854ca91318274E47F4eEE245CF28', 3, 1, '', NOW(), NOW(), 0),
(5700, '0x4798388e3adE569570Df626040F07DF71135C48E', 2, 1, '', NOW(), NOW(), 0);
