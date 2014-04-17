<?php
include_once('./objects/type.php');
class TypeAPI {
    public static function CreateActionEndpoint() {
        $action = TypeObject::buildActionFromObject(json_decode(file_get_contents('php://input')));
	TypeAPI::createAction($action);
    }
    
    public static function CreateRewardEndpoint() {
        $reward = TypeObject::buildRewardFromObject(json_decode(file_get_contents('php://input')));
	TypeAPI::createReward($reward);
    }

    public static function createAction($action) {
	$db = new Database();
	if(!$db->connect() || !$db->createActionType($action)) {
		internalServerError($db->getError());
	}
	
	resourceCreated('{"Result":{"Id":'.$db->connection->insert_id.'}}');
    }

    public static function createReward($reward) {
	$db = new Database();
	if(!$db->connect() || !$db->createRewardType($reward)) {
		internalServerError($db->getError());
	}
	
	resourceCreated('{"Result":{"Id":'.$db->connection->insert_id.'}}');
    }

}
?>
