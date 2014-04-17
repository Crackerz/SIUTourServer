<?php
include_once('./objects/action.php');
class ActionAPI {
	public static function CreateEndpoint() {
		$action = ActionObject::buildActionFromObject(json_decode(file_get_contents('php://input')));
		ActionAPI::create($action);
	}

	public static function create($action) {
		$db = new Database();
		if(!$db->connect() || !$db->createAction($action)) {
			internalServerError($db->getError());
		}

		resourceCreated('{"Result":{"Id":'.$db->connection->insert_id.'}}');
	}
}
?>
