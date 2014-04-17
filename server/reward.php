<?php
include_once('./objects/reward.php');
class RewardAPI {
    public static function CreateEndpoint() {
        $reward = RewardObject::buildRewardFromObject(json_decode(file_get_contents('php://input')));
	RewardAPI::create($reward);
    }

    public static function create($reward) {
	$db = new Database();
	if(!$db->connect() || !$db->createReward($reward)) {
		internalServerError($db->getError());
	}
	
	resourceCreated('{"Result":{"Id":'.$db->connection->insert_id.'}}');
    }
}
?>
