

CREATE INDEX idx_block_reward_release_detail_01 ON block_reward_release_detail (asset_pack_id, height);
CREATE INDEX idx_block_reward_release_detail_02 ON block_reward_release_detail (miner_id, height);

CREATE INDEX idx_customer_asset_pack_01 ON customer_asset_pack (address, asset_pack_id);
CREATE INDEX idx_event_sync_01 ON event_sync (asset_pack_id, height);

CREATE INDEX idx_fund_details_02 ON fund_details (asset_pack_id, fund_type,tx_type);
CREATE INDEX idx_fund_details_03 ON fund_details (from_address, fund_type,tx_type);
CREATE INDEX idx_sector_release_record_01 ON sector_release_record (miner_id,asset_pack_id);

