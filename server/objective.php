<?php
include_once('./objects/objective.php');
class ObjectiveAPI {
	public static function CreateEndpoint() {
		$objective = ObjectiveObject::buildObjectiveFromObject(json_decode(file_get_contents('php://input')));
		ObjectiveAPI::create($objective);
	}

	public static function create($objective) {
		$db = new Database();
		if(!$db->connect() || !$db->createObjective($objective)) {
			internalServerError($db->getError());
		}

		resourceCreated('{"Result":{"Id":'.$db->connection->insert_id.'}}');
	}
}
